package armresourcemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Get.json
func ExampleMoveCollectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMoveCollectionsClient().Get(ctx, "rg1", "movecollection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MoveCollection = armresourcemover.MoveCollection{
	// 	Name: to.Ptr("movecollection1"),
	// 	Type: to.Ptr("Microsoft.Migrate/MoveCollections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1"),
	// 	Identity: &armresourcemover.Identity{
	// 		Type: to.Ptr(armresourcemover.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("7488fa50-1c8e-42c4-a117-38c8d05d8b72"),
	// 		TenantID: to.Ptr("cc7e5736-dbba-4059-a621-664eab9c1d80"),
	// 	},
	// 	Location: to.Ptr("United States"),
	// 	Properties: &armresourcemover.MoveCollectionProperties{
	// 		MoveType: to.Ptr(armresourcemover.MoveTypeRegionToRegion),
	// 		ProvisioningState: to.Ptr(armresourcemover.ProvisioningStateSucceeded),
	// 		SourceRegion: to.Ptr("eastus"),
	// 		TargetRegion: to.Ptr("westus"),
	// 		Version: to.Ptr("V1"),
	// 	},
	// 	SystemData: &armresourcemover.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-01T15:06:54.2757906Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T17:16:24.3644126Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 	},
	// }
}
