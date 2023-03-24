package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/ProvisionedClustersGetUpgradeProfile.json
func ExampleProvisionedClustersClient_GetUpgradeProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvisionedClustersClient().GetUpgradeProfile(ctx, "test-arcappliance-resgrp", "test-hybridakscluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProvisionedClusterUpgradeProfile = armhybridcontainerservice.ProvisionedClusterUpgradeProfile{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.HybridContainerService/provisionedClusters/upgradeprofiles"),
	// 	ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.HybridContainerService/provisionedClusters/test-hybridakscluster/upgradeprofiles"),
	// 	Properties: &armhybridcontainerservice.ProvisionedClusterUpgradeProfileProperties{
	// 		AgentPoolProfiles: []*armhybridcontainerservice.ProvisionedClusterPoolUpgradeProfile{
	// 			{
	// 				Name: to.Ptr("agent"),
	// 				KubernetesVersion: to.Ptr("1.7.7"),
	// 				OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
	// 				Upgrades: []*armhybridcontainerservice.ProvisionedClusterPoolUpgradeProfileProperties{
	// 					{
	// 						KubernetesVersion: to.Ptr("1.7.9"),
	// 					},
	// 					{
	// 						IsPreview: to.Ptr(true),
	// 						KubernetesVersion: to.Ptr("1.7.11"),
	// 				}},
	// 		}},
	// 		ControlPlaneProfile: &armhybridcontainerservice.ProvisionedClusterPoolUpgradeProfile{
	// 			Name: to.Ptr("master"),
	// 			KubernetesVersion: to.Ptr("1.7.7"),
	// 			OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
	// 			Upgrades: []*armhybridcontainerservice.ProvisionedClusterPoolUpgradeProfileProperties{
	// 				{
	// 					IsPreview: to.Ptr(true),
	// 					KubernetesVersion: to.Ptr("1.7.9"),
	// 				},
	// 				{
	// 					KubernetesVersion: to.Ptr("1.7.11"),
	// 			}},
	// 		},
	// 	},
	// }
}
