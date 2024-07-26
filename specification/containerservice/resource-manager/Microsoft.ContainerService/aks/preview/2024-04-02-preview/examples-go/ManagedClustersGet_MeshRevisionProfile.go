package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/examples/ManagedClustersGet_MeshRevisionProfile.json
func ExampleManagedClustersClient_GetMeshRevisionProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedClustersClient().GetMeshRevisionProfile(ctx, "location1", "istio", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MeshRevisionProfile = armcontainerservice.MeshRevisionProfile{
	// 	Name: to.Ptr("istio"),
	// 	Type: to.Ptr("Microsoft.ContainerService/locations/meshRevisionProfiles"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ContainerService/locations/location1/meshRevisionProfiles/istio"),
	// 	Properties: &armcontainerservice.MeshRevisionProfileProperties{
	// 		MeshRevisions: []*armcontainerservice.MeshRevision{
	// 			{
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
	// 					{
	// 						CompatibleWith: []*armcontainerservice.CompatibleVersions{
	// 							{
	// 								Name: to.Ptr("kubernetes"),
	// 								Versions: []*string{
	// 									to.Ptr("1.24"),
	// 									to.Ptr("1.25"),
	// 									to.Ptr("1.26"),
	// 									to.Ptr("1.27")},
	// 							}},
	// 							Revision: to.Ptr("asm-1-18"),
	// 							Upgrades: []*string{
	// 							},
	// 					}},
	// 				},
	// 			}
}
