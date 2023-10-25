package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/89260be1a92c914b7b48af8e8f75938d5e76851d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-08-02-preview/examples/ManagedClustersList_MeshRevisionProfiles.json
func ExampleManagedClustersClient_NewListMeshRevisionProfilesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedClustersClient().NewListMeshRevisionProfilesPager("location1", nil)
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
		// page.MeshRevisionProfileList = armcontainerservice.MeshRevisionProfileList{
		// 	Value: []*armcontainerservice.MeshRevisionProfile{
		// 		{
		// 			Name: to.Ptr("istio"),
		// 			Type: to.Ptr("Microsoft.ContainerService/locations/meshRevisionProfiles"),
		// 			ID: to.Ptr("/subscriptions/subid1/providers/Microsoft.ContainerService/locations/location1/meshRevisionProfiles/istio"),
		// 			Properties: &armcontainerservice.MeshRevisionProfileProperties{
		// 				MeshRevisions: []*armcontainerservice.MeshRevision{
		// 					{
		// 						CompatibleWith: []*armcontainerservice.CompatibleVersions{
		// 							{
		// 								Name: to.Ptr("kubernetes"),
		// 								Versions: []*string{
		// 									to.Ptr("1.23"),
		// 									to.Ptr("1.24"),
		// 									to.Ptr("1.25"),
		// 									to.Ptr("1.26")},
		// 							}},
		// 							Revision: to.Ptr("asm-1-17"),
		// 							Upgrades: []*string{
		// 								to.Ptr("1-18")},
		// 							},
		// 							{
		// 								CompatibleWith: []*armcontainerservice.CompatibleVersions{
		// 									{
		// 										Name: to.Ptr("kubernetes"),
		// 										Versions: []*string{
		// 											to.Ptr("1.24"),
		// 											to.Ptr("1.25"),
		// 											to.Ptr("1.26"),
		// 											to.Ptr("1.27")},
		// 									}},
		// 									Revision: to.Ptr("asm-1-18"),
		// 									Upgrades: []*string{
		// 									},
		// 							}},
		// 						},
		// 				}},
		// 			}
	}
}
