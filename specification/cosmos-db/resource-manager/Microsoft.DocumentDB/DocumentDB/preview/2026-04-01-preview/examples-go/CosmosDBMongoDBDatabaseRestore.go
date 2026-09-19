package armcosmos_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-04-01-preview/CosmosDBMongoDBDatabaseRestore.json
func ExampleMongoDBResourcesClient_BeginCreateUpdateMongoDBDatabase_cosmosDbMongoDbdatabaseRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMongoDBResourcesClient().BeginCreateUpdateMongoDBDatabase(ctx, "rg1", "ddb1", "databaseName", armcosmos.MongoDBDatabaseCreateUpdateParameters{
		Location: to.Ptr("West US"),
		Properties: &armcosmos.MongoDBDatabaseCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{},
			Resource: &armcosmos.MongoDBDatabaseResource{
				CreateMode: to.Ptr(armcosmos.CreateModeRestore),
				ID:         to.Ptr("databaseName"),
				RestoreParameters: &armcosmos.ResourceRestoreParameters{
					RestoreSource:          to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.DocumentDB/locations/WestUS/restorableDatabaseAccounts/restorableDatabaseAccountId"),
					RestoreTimestampInUTC:  to.Ptr(time.Date(2022, time.July, 20, 18, 28, 0, 0, time.UTC)),
					RestoreWithTTLDisabled: to.Ptr(false),
				},
			},
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.MongoDBResourcesClientCreateUpdateMongoDBDatabaseResponse{
	// 	MongoDBDatabaseGetResults: armcosmos.MongoDBDatabaseGetResults{
	// 		Name: to.Ptr("databaseName"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/mongodbDatabases/databaseName"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armcosmos.MongoDBDatabaseGetProperties{
	// 			Resource: &armcosmos.MongoDBDatabaseGetPropertiesResource{
	// 				ID: to.Ptr("updatedDatabaseName"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
