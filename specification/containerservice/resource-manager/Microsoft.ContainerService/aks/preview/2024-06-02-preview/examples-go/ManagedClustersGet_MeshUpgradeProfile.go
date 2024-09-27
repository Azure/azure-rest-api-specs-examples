package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-06-02-preview/examples/ManagedClustersGet_MeshUpgradeProfile.json
func ExampleManagedClustersClient_GetMeshUpgradeProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedClustersClient().GetMeshUpgradeProfile(ctx, "rg1", "clustername1", "istio", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MeshUpgradeProfile = armcontainerservice.MeshUpgradeProfile{
	// 	Name: to.Ptr("istio"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/meshUpgradeProfiles"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/meshUpgradeProfiles/istio"),
	// 	Properties: &armcontainerservice.MeshUpgradeProfileProperties{
	// 		CompatibleWith: []*armcontainerservice.CompatibleVersions{
	// 			{
	// 				Name: to.Ptr("kubernetes"),
	// 				Versions: []*string{
	// 					to.Ptr("1.23"),
	// 					to.Ptr("1.24"),
	// 					to.Ptr("1.25"),
	// 					to.Ptr("1.26")},
	// 			}},
	// 			Revision: to.Ptr("asm-1-17"),
	// 			Upgrades: []*string{
	// 				to.Ptr("1-18")},
	// 			},
	// 		}
}
