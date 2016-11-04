/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kutil

import (
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	"os"
	"os/exec"
	"path"
	"strings"
)

// KubeconfigBuilder builds a kubecfg file
// This logic previously lives in the bash scripts (create-kubeconfig in cluster/common.sh)
type KubeconfigBuilder struct {
	KubectlPath    string
	KubeconfigPath string

	KubeMasterIP string

	Context   string
	Namespace string

	KubeBearerToken string
	KubeUser        string
	KubePassword    string

	CACert     []byte
	ClientCert []byte
	ClientKey  []byte
}

func NewKubeconfigBuilder() *KubeconfigBuilder {
	c := &KubeconfigBuilder{}
	c.KubectlPath = "kubectl" // default to in-path

	kubeconfig := os.Getenv(clientcmd.RecommendedConfigPathEnvVar)
	if kubeconfig == "" {
		kubeconfig = clientcmd.RecommendedHomeFile
	}
	c.KubeconfigPath = kubeconfig
	return c
}

func (c *KubeconfigBuilder) BuildRestConfig() (*restclient.Config, error) {
	restConfig := &restclient.Config{
		Host: "https://" + c.KubeMasterIP,
	}
	restConfig.CAData = c.CACert
	restConfig.CertData = c.ClientCert
	restConfig.KeyData = c.ClientKey

	// username/password or bearer token may be set, but not both
	if c.KubeBearerToken != "" {
		restConfig.BearerToken = c.KubeBearerToken
	} else {
		restConfig.Username = c.KubeUser
		restConfig.Password = c.KubePassword
	}

	return restConfig, nil
}

func (c *KubeconfigBuilder) WriteKubecfg() error {
	tmpdir, err := ioutil.TempDir("", "k8s")
	if err != nil {
		return fmt.Errorf("error creating temporary directory: %v", err)
	}

	defer func() {
		err := os.RemoveAll(tmpdir)
		if err != nil {
			glog.Warningf("error deleting tempdir %q: %v", tmpdir, err)
		}
	}()

	if _, err := os.Stat(c.KubeconfigPath); os.IsNotExist(err) {
		err := os.MkdirAll(path.Dir(c.KubeconfigPath), 0700)
		if err != nil {
			return fmt.Errorf("error creating directories for %q: %v", c.KubeconfigPath, err)
		}
		f, err := os.OpenFile(c.KubeconfigPath, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			return fmt.Errorf("error creating config file %q: %v", c.KubeconfigPath, err)
		}
		f.Close()
	}

	var clusterArgs []string

	clusterArgs = append(clusterArgs, "--server=https://"+c.KubeMasterIP)

	if c.CACert == nil {
		clusterArgs = append(clusterArgs, "--insecure-skip-tls-verify=true")
	} else {
		caCert := path.Join(tmpdir, "ca.crt")
		if err := ioutil.WriteFile(caCert, c.CACert, 0600); err != nil {
			return err
		}
		clusterArgs = append(clusterArgs, "--certificate-authority="+caCert)
		clusterArgs = append(clusterArgs, "--embed-certs=true")
	}

	var userArgs []string

	if c.KubeBearerToken != "" {
		userArgs = append(userArgs, "--token="+c.KubeBearerToken)
	} else if c.KubeUser != "" && c.KubePassword != "" {
		userArgs = append(userArgs, "--username="+c.KubeUser)
		userArgs = append(userArgs, "--password="+c.KubePassword)
	}

	if c.ClientCert != nil && c.ClientKey != nil {
		clientCert := path.Join(tmpdir, "client.crt")
		if err := ioutil.WriteFile(clientCert, c.ClientCert, 0600); err != nil {
			return err
		}
		clientKey := path.Join(tmpdir, "client.key")
		if err := ioutil.WriteFile(clientKey, c.ClientKey, 0600); err != nil {
			return err
		}

		userArgs = append(userArgs, "--client-certificate="+clientCert)
		userArgs = append(userArgs, "--client-key="+clientKey)
		userArgs = append(userArgs, "--embed-certs=true")
	}

	setClusterArgs := []string{"config", "set-cluster", c.Context}
	setClusterArgs = append(setClusterArgs, clusterArgs...)
	err = c.execKubectl(setClusterArgs...)
	if err != nil {
		return err
	}

	if len(userArgs) != 0 {
		setCredentialsArgs := []string{"config", "set-credentials", c.Context}
		setCredentialsArgs = append(setCredentialsArgs, userArgs...)
		err := c.execKubectl(setCredentialsArgs...)
		if err != nil {
			return err
		}
	}

	{
		args := []string{"config", "set-context", c.Context, "--cluster=" + c.Context, "--user=" + c.Context}
		if c.Namespace != "" {
			args = append(args, "--namespace", c.Namespace)
		}
		err = c.execKubectl(args...)
		if err != nil {
			return err
		}
	}
	err = c.execKubectl("config", "use-context", c.Context, "--cluster="+c.Context, "--user="+c.Context)
	if err != nil {
		return err
	}

	// If we have a bearer token, also create a credential entry with basic auth
	// so that it is easy to discover the basic auth password for your cluster
	// to use in a web browser.
	if c.KubeUser != "" && c.KubePassword != "" {
		err := c.execKubectl("config", "set-credentials", c.Context+"-basic-auth", "--username="+c.KubeUser, "--password="+c.KubePassword)
		if err != nil {
			return err
		}
	}

	split := strings.Split(c.KubeconfigPath, ":")
	path := c.KubeconfigPath
	if len(split) > 1 {
		path = split[0]
	}
	fmt.Printf("Wrote config for %s to %q\n", c.Context, path)
	return nil
}

func (c *KubeconfigBuilder) execKubectl(args ...string) error {
	cmd := exec.Command(c.KubectlPath, args...)
	env := os.Environ()
	env = append(env, fmt.Sprintf("KUBECONFIG=%s", c.KubeconfigPath))
	cmd.Env = env

	glog.V(2).Infof("Running command: %s", strings.Join(cmd.Args, " "))
	output, err := cmd.CombinedOutput()
	if err != nil {
		if len(output) != 0 {
			glog.Info("error running kubectl.  Output follows:")
			glog.Info(string(output))
		}
		return fmt.Errorf("error running kubectl: %v", err)
	}

	glog.V(2).Info(string(output))
	return nil
}
