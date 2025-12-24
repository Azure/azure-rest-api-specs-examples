package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/455d20a5e76d8184f7cff960501a57e1f88986b7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/ManagedClustersList_MeshUpgradeProfiles.json
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
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/meshUpgradeProfiles/istio"),
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
		// 						to.Ptr("1-18")},
		// 					},
		// 			}},
		// 		}
	}
}
