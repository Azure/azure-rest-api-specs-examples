package armresourcemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_ListMoveCollectionsBySubscription.json
func ExampleMoveCollectionsClient_NewListMoveCollectionsBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMoveCollectionsClient().NewListMoveCollectionsBySubscriptionPager(nil)
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
		// page.MoveCollectionResultList = armresourcemover.MoveCollectionResultList{
		// 	Value: []*armresourcemover.MoveCollection{
		// 		{
		// 			Name: to.Ptr("movecollection1"),
		// 			Type: to.Ptr("Microsoft.Migrate/MoveCollections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1"),
		// 			Identity: &armresourcemover.Identity{
		// 				Type: to.Ptr(armresourcemover.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("7488fa50-1c8e-42c4-a117-38c8d05d8b72"),
		// 				TenantID: to.Ptr("cc7e5736-dbba-4059-a621-664eab9c1d80"),
		// 			},
		// 			Location: to.Ptr("United States"),
		// 			Properties: &armresourcemover.MoveCollectionProperties{
		// 				MoveType: to.Ptr(armresourcemover.MoveTypeRegionToRegion),
		// 				ProvisioningState: to.Ptr(armresourcemover.ProvisioningStateSucceeded),
		// 				SourceRegion: to.Ptr("eastus"),
		// 				TargetRegion: to.Ptr("westus"),
		// 				Version: to.Ptr("V1"),
		// 			},
		// 			SystemData: &armresourcemover.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-29T15:06:54.275Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@microsoft.com"),
		// 				CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-29T15:06:54.275Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
