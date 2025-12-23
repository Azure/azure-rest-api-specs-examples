package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/455d20a5e76d8184f7cff960501a57e1f88986b7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/ManagedClustersGetUpgradeProfile.json
func ExampleManagedClustersClient_GetUpgradeProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedClustersClient().GetUpgradeProfile(ctx, "rg1", "clustername1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedClusterUpgradeProfile = armcontainerservice.ManagedClusterUpgradeProfile{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/upgradeprofiles"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/upgradeprofiles/default"),
	// 	Properties: &armcontainerservice.ManagedClusterUpgradeProfileProperties{
	// 		AgentPoolProfiles: []*armcontainerservice.ManagedClusterPoolUpgradeProfile{
	// 			{
	// 				Name: to.Ptr("agent"),
	// 				KubernetesVersion: to.Ptr("1.7.7"),
	// 				OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 				Upgrades: []*armcontainerservice.ManagedClusterPoolUpgradeProfileUpgradesItem{
	// 					{
	// 						KubernetesVersion: to.Ptr("1.7.9"),
	// 					},
	// 					{
	// 						IsPreview: to.Ptr(true),
	// 						KubernetesVersion: to.Ptr("1.7.11"),
	// 				}},
	// 		}},
	// 		ControlPlaneProfile: &armcontainerservice.ManagedClusterPoolUpgradeProfile{
	// 			Name: to.Ptr("master"),
	// 			KubernetesVersion: to.Ptr("1.7.7"),
	// 			OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 			Upgrades: []*armcontainerservice.ManagedClusterPoolUpgradeProfileUpgradesItem{
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
