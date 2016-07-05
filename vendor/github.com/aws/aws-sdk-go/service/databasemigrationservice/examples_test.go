// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package databasemigrationservice_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDatabaseMigrationService_AddTagsToResource() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.AddTagsToResourceInput{
		ResourceArn: aws.String("String"), // Required
		Tags: []*databasemigrationservice.Tag{ // Required
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.AddTagsToResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_CreateEndpoint() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.CreateEndpointInput{
		EndpointIdentifier:        aws.String("String"),                       // Required
		EndpointType:              aws.String("ReplicationEndpointTypeValue"), // Required
		EngineName:                aws.String("String"),                       // Required
		Password:                  aws.String("SecretString"),                 // Required
		Port:                      aws.Int64(1),                               // Required
		ServerName:                aws.String("String"),                       // Required
		Username:                  aws.String("String"),                       // Required
		DatabaseName:              aws.String("String"),
		ExtraConnectionAttributes: aws.String("String"),
		KmsKeyId:                  aws.String("String"),
		Tags: []*databasemigrationservice.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateEndpoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_CreateReplicationInstance() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.CreateReplicationInstanceInput{
		ReplicationInstanceClass:         aws.String("String"), // Required
		ReplicationInstanceIdentifier:    aws.String("String"), // Required
		AllocatedStorage:                 aws.Int64(1),
		AutoMinorVersionUpgrade:          aws.Bool(true),
		AvailabilityZone:                 aws.String("String"),
		EngineVersion:                    aws.String("String"),
		KmsKeyId:                         aws.String("String"),
		PreferredMaintenanceWindow:       aws.String("String"),
		PubliclyAccessible:               aws.Bool(true),
		ReplicationSubnetGroupIdentifier: aws.String("String"),
		Tags: []*databasemigrationservice.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
		VpcSecurityGroupIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateReplicationInstance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_CreateReplicationSubnetGroup() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.CreateReplicationSubnetGroupInput{
		ReplicationSubnetGroupDescription: aws.String("String"), // Required
		ReplicationSubnetGroupIdentifier:  aws.String("String"), // Required
		SubnetIds: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Tags: []*databasemigrationservice.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateReplicationSubnetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_CreateReplicationTask() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.CreateReplicationTaskInput{
		MigrationType:             aws.String("MigrationTypeValue"), // Required
		ReplicationInstanceArn:    aws.String("String"),             // Required
		ReplicationTaskIdentifier: aws.String("String"),             // Required
		SourceEndpointArn:         aws.String("String"),             // Required
		TableMappings:             aws.String("String"),             // Required
		TargetEndpointArn:         aws.String("String"),             // Required
		CdcStartTime:              aws.Time(time.Now()),
		ReplicationTaskSettings:   aws.String("String"),
		Tags: []*databasemigrationservice.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateReplicationTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DeleteEndpoint() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DeleteEndpointInput{
		EndpointArn: aws.String("String"), // Required
	}
	resp, err := svc.DeleteEndpoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DeleteReplicationInstance() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DeleteReplicationInstanceInput{
		ReplicationInstanceArn: aws.String("String"), // Required
	}
	resp, err := svc.DeleteReplicationInstance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DeleteReplicationSubnetGroup() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DeleteReplicationSubnetGroupInput{
		ReplicationSubnetGroupIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.DeleteReplicationSubnetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DeleteReplicationTask() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DeleteReplicationTaskInput{
		ReplicationTaskArn: aws.String("String"), // Required
	}
	resp, err := svc.DeleteReplicationTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeAccountAttributes() {
	svc := databasemigrationservice.New(session.New())

	var params *databasemigrationservice.DescribeAccountAttributesInput
	resp, err := svc.DescribeAccountAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeConnections() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DescribeConnectionsInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeConnections(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeEndpointTypes() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DescribeEndpointTypesInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeEndpointTypes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeEndpoints() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DescribeEndpointsInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeEndpoints(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeOrderableReplicationInstances() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DescribeOrderableReplicationInstancesInput{
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeOrderableReplicationInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeRefreshSchemasStatus() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DescribeRefreshSchemasStatusInput{
		EndpointArn: aws.String("String"), // Required
	}
	resp, err := svc.DescribeRefreshSchemasStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeReplicationInstances() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DescribeReplicationInstancesInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeReplicationInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeReplicationSubnetGroups() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DescribeReplicationSubnetGroupsInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeReplicationSubnetGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeReplicationTasks() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DescribeReplicationTasksInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeReplicationTasks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeSchemas() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DescribeSchemasInput{
		EndpointArn: aws.String("String"), // Required
		Marker:      aws.String("String"),
		MaxRecords:  aws.Int64(1),
	}
	resp, err := svc.DescribeSchemas(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeTableStatistics() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.DescribeTableStatisticsInput{
		ReplicationTaskArn: aws.String("String"), // Required
		Marker:             aws.String("String"),
		MaxRecords:         aws.Int64(1),
	}
	resp, err := svc.DescribeTableStatistics(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ListTagsForResource() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.ListTagsForResourceInput{
		ResourceArn: aws.String("String"), // Required
	}
	resp, err := svc.ListTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ModifyEndpoint() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.ModifyEndpointInput{
		EndpointArn:               aws.String("String"), // Required
		DatabaseName:              aws.String("String"),
		EndpointIdentifier:        aws.String("String"),
		EndpointType:              aws.String("ReplicationEndpointTypeValue"),
		EngineName:                aws.String("String"),
		ExtraConnectionAttributes: aws.String("String"),
		Password:                  aws.String("SecretString"),
		Port:                      aws.Int64(1),
		ServerName:                aws.String("String"),
		Username:                  aws.String("String"),
	}
	resp, err := svc.ModifyEndpoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ModifyReplicationInstance() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.ModifyReplicationInstanceInput{
		ReplicationInstanceArn:        aws.String("String"), // Required
		AllocatedStorage:              aws.Int64(1),
		AllowMajorVersionUpgrade:      aws.Bool(true),
		ApplyImmediately:              aws.Bool(true),
		AutoMinorVersionUpgrade:       aws.Bool(true),
		EngineVersion:                 aws.String("String"),
		PreferredMaintenanceWindow:    aws.String("String"),
		ReplicationInstanceClass:      aws.String("String"),
		ReplicationInstanceIdentifier: aws.String("String"),
		VpcSecurityGroupIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyReplicationInstance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ModifyReplicationSubnetGroup() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.ModifyReplicationSubnetGroupInput{
		ReplicationSubnetGroupIdentifier: aws.String("String"), // Required
		SubnetIds: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		ReplicationSubnetGroupDescription: aws.String("String"),
	}
	resp, err := svc.ModifyReplicationSubnetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_RefreshSchemas() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.RefreshSchemasInput{
		EndpointArn:            aws.String("String"), // Required
		ReplicationInstanceArn: aws.String("String"), // Required
	}
	resp, err := svc.RefreshSchemas(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_RemoveTagsFromResource() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.RemoveTagsFromResourceInput{
		ResourceArn: aws.String("String"), // Required
		TagKeys: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_StartReplicationTask() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.StartReplicationTaskInput{
		ReplicationTaskArn:       aws.String("String"),                        // Required
		StartReplicationTaskType: aws.String("StartReplicationTaskTypeValue"), // Required
		CdcStartTime:             aws.Time(time.Now()),
	}
	resp, err := svc.StartReplicationTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_StopReplicationTask() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.StopReplicationTaskInput{
		ReplicationTaskArn: aws.String("String"), // Required
	}
	resp, err := svc.StopReplicationTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_TestConnection() {
	svc := databasemigrationservice.New(session.New())

	params := &databasemigrationservice.TestConnectionInput{
		EndpointArn:            aws.String("String"), // Required
		ReplicationInstanceArn: aws.String("String"), // Required
	}
	resp, err := svc.TestConnection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
