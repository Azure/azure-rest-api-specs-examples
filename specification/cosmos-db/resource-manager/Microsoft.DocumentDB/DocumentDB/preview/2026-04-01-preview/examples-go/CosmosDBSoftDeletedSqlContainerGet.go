package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-04-01-preview/CosmosDBSoftDeletedSqlContainerGet.json
func ExampleSoftDeletedSQLContainersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSoftDeletedSQLContainersClient().Get(ctx, "rg1", "West US", "softdeleted-cosmosdb-1", "MyDatabase", "MyContainer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.SoftDeletedSQLContainersClientGetResponse{
	// 	SoftDeletedSQLContainerGetResult: armcosmos.SoftDeletedSQLContainerGetResult{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/locations/westus/softDeletedDatabaseAccounts/softdeleted-cosmosdb-1/softDeletedSqlDatabases/MyDatabase/softDeletedSqlContainers/MyContainer"),
	// 		Name: to.Ptr("MyContainer"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/locations/softDeletedDatabaseAccounts/softDeletedSqlDatabases/softDeletedSqlContainers"),
	// 		Properties: &armcosmos.SoftDeletedSQLContainerProperties{
	// 			SoftDeletionMetadata: &armcosmos.SoftDeletionMetadata{
	// 				IsSoftDeleted: to.Ptr(true),
	// 				SoftDeletionStartTimestamp: to.Ptr[int64](1729000530),
	// 				SoftDeletionResourceExpirationTimestamp: to.Ptr[int64](1731592530),
	// 			},
	// 			Resource: &armcosmos.SoftDeletedSQLContainerResource{
	// 				ID: to.Ptr("MyContainer"),
	// 				Rid: to.Ptr("PD5DALigDgwBAAAAAAAAAA=="),
	// 				PartitionKey: &armcosmos.ContainerPartitionKey{
	// 					Paths: []*string{
	// 						to.Ptr("/accountNumber"),
	// 					},
	// 					Kind: to.Ptr(armcosmos.PartitionKindHash),
	// 				},
	// 				DefaultTTL: to.Ptr[int32](3600),
	// 			},
	// 		},
	// 	},
	// }
}
