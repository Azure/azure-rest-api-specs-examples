package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/ListProvisionedClusterByResourceGroup.json
func ExampleProvisionedClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProvisionedClustersClient().NewListByResourceGroupPager("test-arcappliance-resgrp", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ProvisionedClustersResponseListResult = armhybridcontainerservice.ProvisionedClustersResponseListResult{
		// 	Value: []*armhybridcontainerservice.ProvisionedClustersResponse{
		// 		{
		// 			Name: to.Ptr("test-hybridakscluster"),
		// 			Type: to.Ptr("microsoft.hybridcontainerservice/provisionedclusters"),
		// 			ID: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/provisionedClusters/test-hybridakscluster"),
		// 			Location: to.Ptr("westus"),
		// 			ExtendedLocation: &armhybridcontainerservice.ProvisionedClustersResponseExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armhybridcontainerservice.ProvisionedClustersResponseProperties{
		// 				AgentPoolProfiles: []*armhybridcontainerservice.NamedAgentPoolProfile{
		// 					{
		// 						Name: to.Ptr("default-nodepool-1"),
		// 						Count: to.Ptr[int32](1),
		// 						OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
		// 						VMSize: to.Ptr("Standard_A4_v2"),
		// 				}},
		// 				CloudProviderProfile: &armhybridcontainerservice.CloudProviderProfile{
		// 					InfraNetworkProfile: &armhybridcontainerservice.CloudProviderProfileInfraNetworkProfile{
		// 						VnetSubnetIDs: []*string{
		// 							to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/virtualNetworks/test-vnet-static")},
		// 						},
		// 						InfraStorageProfile: &armhybridcontainerservice.CloudProviderProfileInfraStorageProfile{
		// 							StorageSpaceIDs: []*string{
		// 								to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/storageSpaces/test-storage")},
		// 							},
		// 						},
		// 						ControlPlane: &armhybridcontainerservice.ControlPlaneProfile{
		// 							LinuxProfile: &armhybridcontainerservice.LinuxProfileProperties{
		// 								SSH: &armhybridcontainerservice.LinuxProfilePropertiesSSH{
		// 									PublicKeys: []*armhybridcontainerservice.LinuxProfilePropertiesSSHPublicKeysItem{
		// 										{
		// 											KeyData: to.Ptr("ssh-rsa AAAAB1NzaC1yc2EAAAADAQABAAACAQCY......"),
		// 									}},
		// 								},
		// 							},
		// 							Count: to.Ptr[int32](1),
		// 							OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
		// 							VMSize: to.Ptr("Standard_A4_v2"),
		// 						},
		// 						KubernetesVersion: to.Ptr("v1.20.5"),
		// 						LinuxProfile: &armhybridcontainerservice.LinuxProfileProperties{
		// 							SSH: &armhybridcontainerservice.LinuxProfilePropertiesSSH{
		// 								PublicKeys: []*armhybridcontainerservice.LinuxProfilePropertiesSSHPublicKeysItem{
		// 									{
		// 										KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCY......."),
		// 								}},
		// 							},
		// 						},
		// 						NetworkProfile: &armhybridcontainerservice.NetworkProfile{
		// 							LoadBalancerProfile: &armhybridcontainerservice.LoadBalancerProfile{
		// 								LinuxProfile: &armhybridcontainerservice.LinuxProfileProperties{
		// 									SSH: &armhybridcontainerservice.LinuxProfilePropertiesSSH{
		// 										PublicKeys: []*armhybridcontainerservice.LinuxProfilePropertiesSSHPublicKeysItem{
		// 											{
		// 												KeyData: to.Ptr("ssh-rsa AAAAB2NzaC1yc2EAAAADAQABAAACAQCY......"),
		// 										}},
		// 									},
		// 								},
		// 								Count: to.Ptr[int32](1),
		// 								OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
		// 								VMSize: to.Ptr("Standard_K8S3_v1"),
		// 							},
		// 							LoadBalancerSKU: to.Ptr(armhybridcontainerservice.LoadBalancerSKUUnstackedHaproxy),
		// 							NetworkPolicy: to.Ptr(armhybridcontainerservice.NetworkPolicyCalico),
		// 							PodCidr: to.Ptr("10.244.0.0/16"),
		// 						},
		// 						ProvisioningState: to.Ptr(armhybridcontainerservice.ProvisioningStateSucceeded),
		// 					},
		// 			}},
		// 		}
	}
}
