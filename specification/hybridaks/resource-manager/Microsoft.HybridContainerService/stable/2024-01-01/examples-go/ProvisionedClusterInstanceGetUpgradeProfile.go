package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/ProvisionedClusterInstanceGetUpgradeProfile.json
func ExampleProvisionedClusterInstancesClient_GetUpgradeProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvisionedClusterInstancesClient().GetUpgradeProfile(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProvisionedClusterUpgradeProfile = armhybridcontainerservice.ProvisionedClusterUpgradeProfile{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.HybridContainerService/provisionedClusterInstances/upgradeprofiles"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/upgradeprofiles/default"),
	// 	Properties: &armhybridcontainerservice.ProvisionedClusterUpgradeProfileProperties{
	// 		ControlPlaneProfile: &armhybridcontainerservice.ProvisionedClusterPoolUpgradeProfile{
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
