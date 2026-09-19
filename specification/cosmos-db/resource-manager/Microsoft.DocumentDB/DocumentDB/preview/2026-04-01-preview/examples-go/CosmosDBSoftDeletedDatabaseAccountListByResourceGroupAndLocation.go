package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-04-01-preview/CosmosDBSoftDeletedDatabaseAccountListByResourceGroupAndLocation.json
func ExampleSoftDeletedDatabaseAccountsClient_ListByResourceGroupAndLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSoftDeletedDatabaseAccountsClient().ListByResourceGroupAndLocation(ctx, "rg1", "West US", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.SoftDeletedDatabaseAccountsClientListByResourceGroupAndLocationResponse{
	// 	SoftDeletedDatabaseAccountsListResult: armcosmos.SoftDeletedDatabaseAccountsListResult{
	// 		Value: []*armcosmos.SoftDeletedDatabaseAccountGetResult{
	// 			{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/locations/westus/softDeletedDatabaseAccounts/a1b2c3d4-5678-90ab-cdef-1234567890ab"),
	// 				Name: to.Ptr("a1b2c3d4-5678-90ab-cdef-1234567890ab"),
	// 				Type: to.Ptr("Microsoft.DocumentDB/locations/softDeletedDatabaseAccounts"),
	// 				Properties: &armcosmos.SoftDeletedDatabaseAccountProperties{
	// 					AccountName: to.Ptr("softdeleted-cosmosdb-1"),
	// 					SoftDeletionMetadata: &armcosmos.SoftDeletionMetadata{
	// 						IsSoftDeleted: to.Ptr(true),
	// 						SoftDeletionStartTimestamp: to.Ptr[int64](1729000530),
	// 						SoftDeletionResourceExpirationTimestamp: to.Ptr[int64](1731592530),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/locations/westus/softDeletedDatabaseAccounts/b2c3d4e5-6789-01bc-def2-234567890abc"),
	// 				Name: to.Ptr("b2c3d4e5-6789-01bc-def2-234567890abc"),
	// 				Type: to.Ptr("Microsoft.DocumentDB/locations/softDeletedDatabaseAccounts"),
	// 				Properties: &armcosmos.SoftDeletedDatabaseAccountProperties{
	// 					AccountName: to.Ptr("softdeleted-cosmosdb-2"),
	// 					SoftDeletionMetadata: &armcosmos.SoftDeletionMetadata{
	// 						IsSoftDeleted: to.Ptr(true),
	// 						SoftDeletionStartTimestamp: to.Ptr[int64](1728655215),
	// 						SoftDeletionResourceExpirationTimestamp: to.Ptr[int64](1731247215),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
