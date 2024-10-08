/*
Copyright 2018 The Kubernetes Authors.

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

package app

import (
	"fmt"

	"k8s.io/klog/v2"

	"k8s.io/client-go/informers"
	cloudprovider "k8s.io/cloud-provider"
)

// createCloudProvider helps consolidate what is needed for cloud providers, we explicitly list the things
// that the cloud providers need as parameters, so we can control
func createCloudProvider(logger klog.Logger, cloudProvider string, externalCloudVolumePlugin string, cloudConfigFile string,
	allowUntaggedCloud bool, sharedInformers informers.SharedInformerFactory) (cloudprovider.Interface, ControllerLoopMode, error) {
	var cloud cloudprovider.Interface
	var err error
	loopMode := ExternalLoops

	if cloudprovider.IsExternal(cloudProvider) {
		if externalCloudVolumePlugin == "" {
			// externalCloudVolumePlugin is temporary until we split all cloud providers out.
			// So we just tell the caller that we need to run ExternalLoops without any cloud provider.
			return nil, loopMode, nil
		}
		cloud, err = cloudprovider.InitCloudProvider(externalCloudVolumePlugin, cloudConfigFile)
	} else {
		// in the case where the cloudProvider is not set, we need to inform the caller that there
		// is no cloud provider and that the default loops should be used, and there is no error.
		// this will cause the kube-controller-manager to start the default controller contexts
		// without being attached to a specific cloud.
		if len(cloudProvider) == 0 {
			loopMode = IncludeCloudLoops
		} else {
			// for all other cloudProvider values the internal cloud loops are disabled
			cloudprovider.DisableWarningForProvider(cloudProvider)
			err = cloudprovider.ErrorForDisabledProvider(cloudProvider)
		}
	}
	if err != nil {
		return nil, loopMode, fmt.Errorf("cloud provider could not be initialized: %v", err)
	}

	if cloud != nil && !cloud.HasClusterID() {
		if allowUntaggedCloud {
			logger.Info("Warning: detected a cluster without a ClusterID.  A ClusterID will be required in the future.  Please tag your cluster to avoid any future issues")
		} else {
			return nil, loopMode, fmt.Errorf("no ClusterID Found.  A ClusterID is required for the cloud provider to function properly.  This check can be bypassed by setting the allow-untagged-cloud option")
		}
	}

	if informerUserCloud, ok := cloud.(cloudprovider.InformerUser); ok {
		informerUserCloud.SetInformers(sharedInformers)
	}
	return cloud, loopMode, err
}
