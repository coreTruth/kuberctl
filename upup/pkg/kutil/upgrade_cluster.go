package kutil

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/glog"
	"k8s.io/kops/upup/pkg/api"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup/awsup"
	"time"
)

// UpgradeCluster performs an upgrade of a k8s cluster
type UpgradeCluster struct {
	OldClusterName string
	NewClusterName string
	Cloud          fi.Cloud

	ClusterRegistry *api.ClusterRegistry

	ClusterConfig  *api.Cluster
	InstanceGroups []*api.InstanceGroup
}

func (x *UpgradeCluster) Upgrade() error {
	awsCloud := x.Cloud.(*awsup.AWSCloud)

	cluster := x.ClusterConfig

	newClusterName := x.NewClusterName
	if newClusterName == "" {
		return fmt.Errorf("NewClusterName must be specified")
	}
	oldClusterName := x.OldClusterName
	if oldClusterName == "" {
		return fmt.Errorf("OldClusterName must be specified")
	}

	newKeyStore := x.ClusterRegistry.KeyStore(newClusterName)
	oldKeyStore := x.ClusterRegistry.KeyStore(oldClusterName)

	oldTags := awsCloud.Tags()

	newTags := awsCloud.Tags()
	newTags["KubernetesCluster"] = newClusterName

	// Try to pre-query as much as possible before doing anything destructive
	instances, err := findInstances(awsCloud)
	if err != nil {
		return fmt.Errorf("error finding instances: %v", err)
	}

	volumes, err := DescribeVolumes(x.Cloud)
	if err != nil {
		return err
	}

	dhcpOptions, err := DescribeDhcpOptions(x.Cloud)
	if err != nil {
		return err
	}

	autoscalingGroups, err := findAutoscalingGroups(awsCloud, oldTags)
	if err != nil {
		return err
	}

	elbs, _, err := DescribeELBs(x.Cloud)
	if err != nil {
		return err
	}

	// Find masters
	var masters []*ec2.Instance
	for _, instance := range instances {
		role, _ := awsup.FindEC2Tag(instance.Tags, "Role")
		if role == oldClusterName+"-master" {
			masters = append(masters, instance)
		}
	}
	if len(masters) == 0 {
		return fmt.Errorf("could not find masters")
	}

	// Stop autoscalingGroups
	for _, group := range autoscalingGroups {
		name := aws.StringValue(group.AutoScalingGroupName)
		glog.Infof("Stopping instances in autoscaling group %q", name)

		request := &autoscaling.UpdateAutoScalingGroupInput{
			AutoScalingGroupName: group.AutoScalingGroupName,
			DesiredCapacity:      aws.Int64(0),
			MinSize:              aws.Int64(0),
			MaxSize:              aws.Int64(0),
		}

		_, err := awsCloud.Autoscaling.UpdateAutoScalingGroup(request)
		if err != nil {
			return fmt.Errorf("error updating autoscaling group %q: %v", name, err)
		}
	}

	// Stop masters
	for _, master := range masters {
		masterInstanceID := aws.StringValue(master.InstanceId)

		masterState := aws.StringValue(master.State.Name)
		if masterState == "terminated" {
			glog.Infof("master already terminated: %q", masterInstanceID)
			continue
		}

		glog.Infof("Stopping master: %q", masterInstanceID)

		request := &ec2.StopInstancesInput{
			InstanceIds: []*string{master.InstanceId},
		}

		_, err := awsCloud.EC2.StopInstances(request)
		if err != nil {
			return fmt.Errorf("error stopping master instance: %v", err)
		}
	}

	// Detach volumes from masters
	for _, master := range masters {
		for _, bdm := range master.BlockDeviceMappings {
			if bdm.Ebs == nil || bdm.Ebs.VolumeId == nil {
				continue
			}
			volumeID := aws.StringValue(bdm.Ebs.VolumeId)
			masterInstanceID := aws.StringValue(master.InstanceId)
			glog.Infof("Detaching volume %q from instance %q", volumeID, masterInstanceID)

			request := &ec2.DetachVolumeInput{
				VolumeId:   bdm.Ebs.VolumeId,
				InstanceId: master.InstanceId,
			}

			for {
				_, err := awsCloud.EC2.DetachVolume(request)
				if err != nil {
					if awsup.AWSErrorCode(err) == "IncorrectState" {
						glog.Infof("retrying to detach volume (master has probably not stopped yet): %q", err)
						time.Sleep(5 * time.Second)
						continue
					}
					return fmt.Errorf("error detaching volume %q from master instance %q: %v", volumeID, masterInstanceID, err)
				} else {
					break
				}
			}
		}
	}

	//subnets, err := DescribeSubnets(x.Cloud)
	//if err != nil {
	//	return fmt.Errorf("error finding subnets: %v", err)
	//}
	//for _, s := range subnets {
	//	id := aws.StringValue(s.SubnetId)
	//	err := awsCloud.AddAWSTags(id, newTags)
	//	if err != nil {
	//		return fmt.Errorf("error re-tagging subnet %q: %v", id, err)
	//	}
	//}

	// Retag VPC
	// We have to be careful because VPCs can be shared
	{
		vpcID := cluster.Spec.NetworkID
		retagGateway := false

		if vpcID != "" {
			tags, err := awsCloud.GetTags(vpcID)
			if err != nil {
				return fmt.Errorf("error getting VPC tags: %v", err)
			}

			clusterTag := tags[awsup.TagClusterName]
			if clusterTag != "" {
				if clusterTag != oldClusterName {
					return fmt.Errorf("VPC is tagged with a different cluster: %v", clusterTag)
				}
				replaceTags := make(map[string]string)
				replaceTags[awsup.TagClusterName] = newClusterName

				glog.Infof("Retagging VPC %q", vpcID)

				err := awsCloud.CreateTags(vpcID, replaceTags)
				if err != nil {
					return fmt.Errorf("error re-tagging VPC: %v", err)
				}

				// The VPC was tagged as ours, so make sure the gateway is consistently retagged
				retagGateway = true
			}
		}

		if retagGateway {
			gateways, err := DescribeInternetGatewaysIgnoreTags(x.Cloud)
			if err != nil {
				return fmt.Errorf("error listing gateways: %v", err)
			}
			for _, igw := range gateways {
				match := false
				for _, a := range igw.Attachments {
					if vpcID == aws.StringValue(a.VpcId) {
						match = true
					}
				}
				if !match {
					continue
				}

				id := aws.StringValue(igw.InternetGatewayId)

				clusterTag, _ := awsup.FindEC2Tag(igw.Tags, awsup.TagClusterName)
				if clusterTag == "" || clusterTag == oldClusterName {
					replaceTags := make(map[string]string)
					replaceTags[awsup.TagClusterName] = newClusterName

					glog.Infof("Retagging InternetGateway %q", id)

					err := awsCloud.CreateTags(id, replaceTags)
					if err != nil {
						return fmt.Errorf("error re-tagging InternetGateway: %v", err)
					}
				}
			}
		}
	}

	// Retag DHCP options
	// We have to be careful because DHCP options can be shared
	for _, dhcpOption := range dhcpOptions {
		id := aws.StringValue(dhcpOption.DhcpOptionsId)

		clusterTag, _ := awsup.FindEC2Tag(dhcpOption.Tags, awsup.TagClusterName)
		if clusterTag != "" {
			if clusterTag != oldClusterName {
				return fmt.Errorf("DHCP options are tagged with a different cluster: %v", clusterTag)
			}
			replaceTags := make(map[string]string)
			replaceTags[awsup.TagClusterName] = newClusterName

			glog.Infof("Retagging DHCPOptions %q", id)

			err := awsCloud.CreateTags(id, replaceTags)
			if err != nil {
				return fmt.Errorf("error re-tagging DHCP options: %v", err)
			}
		}

	}

	// Adopt LoadBalancers & LoadBalancer Security Groups
	for _, elb := range elbs {
		id := aws.StringValue(elb.LoadBalancerName)

		// TODO: Batch re-tag?
		replaceTags := make(map[string]string)
		replaceTags[awsup.TagClusterName] = newClusterName

		glog.Infof("Retagging ELB %q", id)
		err := awsCloud.CreateELBTags(id, replaceTags)
		if err != nil {
			return fmt.Errorf("error re-tagging ELB %q: %v", id, err)
		}

	}

	for _, elb := range elbs {
		for _, sg := range elb.SecurityGroups {
			id := aws.StringValue(sg)

			// TODO: Batch re-tag?
			replaceTags := make(map[string]string)
			replaceTags[awsup.TagClusterName] = newClusterName

			glog.Infof("Retagging ELB security group %q", id)
			err := awsCloud.CreateTags(id, replaceTags)
			if err != nil {
				return fmt.Errorf("error re-tagging ELB security group %q: %v", id, err)
			}
		}

	}

	// Adopt Volumes
	for _, volume := range volumes {
		id := aws.StringValue(volume.VolumeId)

		// TODO: Batch re-tag?
		replaceTags := make(map[string]string)
		replaceTags[awsup.TagClusterName] = newClusterName

		name, _ := awsup.FindEC2Tag(volume.Tags, "Name")
		if name == oldClusterName+"-master-pd" {
			glog.Infof("Found master volume %q: %s", id, name)

			az := aws.StringValue(volume.AvailabilityZone)
			replaceTags["Name"] = az + ".etcd-main." + newClusterName
		}
		glog.Infof("Retagging volume %q", id)
		err := awsCloud.CreateTags(id, replaceTags)
		if err != nil {
			return fmt.Errorf("error re-tagging volume %q: %v", id, err)
		}
	}

	cluster.Name = newClusterName
	err = api.CreateClusterConfig(x.ClusterRegistry, cluster, x.InstanceGroups)
	if err != nil {
		return fmt.Errorf("error writing updated configuration: %v", err)
	}

	oldCACertPool, err := oldKeyStore.CertificatePool(fi.CertificateId_CA)
	if err != nil {
		return fmt.Errorf("error reading old CA certs: %v", err)
	}
	for _, ca := range oldCACertPool.Secondary {
		err := newKeyStore.AddCert(fi.CertificateId_CA, ca)
		if err != nil {
			return fmt.Errorf("error importing old CA certs: %v", err)
		}
	}
	if oldCACertPool.Primary != nil {
		err := newKeyStore.AddCert(fi.CertificateId_CA, oldCACertPool.Primary)
		if err != nil {
			return fmt.Errorf("error importing old CA certs: %v", err)
		}
	}

	return nil
}
