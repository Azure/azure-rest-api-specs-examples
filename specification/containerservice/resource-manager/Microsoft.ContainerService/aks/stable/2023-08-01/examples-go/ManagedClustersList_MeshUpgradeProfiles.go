package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0a25ea9680cf080b7d34e8c5f35f564425c6b1f7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-08-01/examples/ManagedClustersList_MeshUpgradeProfiles.json
func ExampleManagedClustersClient_NewListMeshUpgradeProfilesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedClustersClient().NewListMeshUpgradeProfilesPager("rg1", "clustername1", nil)
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
		// page.MeshUpgradeProfileList = armcontainerservice.MeshUpgradeProfileList{
		// 	Value: []*armcontainerservice.MeshUpgradeProfile{
		// 		{
		// 			Name: to.Ptr("istio"),
		// 			Type: to.Ptr("Microsoft.ContainerService/managedClusters/meshUpgradeProfiles"),
		// 			ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/meshUpgradeProfiles/istio"),
		// 			Properties: &armcontainerservice.MeshUpgradeProfileProperties{
		// 				CompatibleWith: []*armcontainerservice.CompatibleVersions{
		// 					{
		// 						Name: to.Ptr("kubernetes"),
		// 						Versions: []*string{
		// 							to.Ptr("1.23"),
		// 							to.Ptr("1.24"),
		// 							to.Ptr("1.25"),
		// 							to.Ptr("1.26")},
		// 					}},
		// 					Revision: to.Ptr("asm-1-17"),
		// 					Upgrades: []*string{
		// 						to.Ptr("asm-1-18")},
		// 					},
		// 			}},
		// 		}
	}
}
