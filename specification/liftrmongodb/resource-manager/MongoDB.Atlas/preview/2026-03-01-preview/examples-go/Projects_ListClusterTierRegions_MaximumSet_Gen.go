package armmongodbatlas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongodbatlas/armmongodbatlas"
)

// Generated from example definition: 2026-03-01-preview/Projects_ListClusterTierRegions_MaximumSet_Gen.json
func ExampleProjectsClient_ListClusterTierRegions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongodbatlas.NewClientFactory("1E4BD993-6890-4E69-8966-81482D7502EF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().ListClusterTierRegions(ctx, "rgopenapi", "myOrganization", "myProject", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmongodbatlas.ProjectsClientListClusterTierRegionsResponse{
	// 	RegionsByTierResponse: armmongodbatlas.RegionsByTierResponse{
	// 		OrganizationID: to.Ptr("507f1f77bcf86cd799439011"),
	// 		ProjectID: to.Ptr("507f1f77bcf86cd799439011"),
	// 		RegionsByTier: []*armmongodbatlas.TierRegions{
	// 			{
	// 				Tier: to.Ptr(armmongodbatlas.ClusterTierFREE),
	// 				Regions: []*string{
	// 					to.Ptr("eastus"),
	// 					to.Ptr("westus"),
	// 				},
	// 			},
	// 			{
	// 				Tier: to.Ptr(armmongodbatlas.ClusterTierFLEX),
	// 				Regions: []*string{
	// 					to.Ptr("eastus"),
	// 					to.Ptr("westeurope"),
	// 				},
	// 			},
	// 			{
	// 				Tier: to.Ptr(armmongodbatlas.ClusterTierM10),
	// 				Regions: []*string{
	// 					to.Ptr("eastus"),
	// 				},
	// 			},
	// 			{
	// 				Tier: to.Ptr(armmongodbatlas.ClusterTierM30),
	// 				Regions: []*string{
	// 					to.Ptr("eastus"),
	// 					to.Ptr("westus"),
	// 					to.Ptr("westeurope"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
