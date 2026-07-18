package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBMongoDBDatabaseCreateUpdate.json
func ExampleMongoDBResourcesClient_BeginCreateUpdateMongoDBDatabase() {
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
		Tags:     map[string]*string{},
		Properties: &armcosmos.MongoDBDatabaseCreateUpdateProperties{
			Resource: &armcosmos.MongoDBDatabaseResource{
				ID: to.Ptr("databaseName"),
			},
			Options: &armcosmos.CreateUpdateOptions{},
		},
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
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/mongodbDatabases/databaseName"),
	// 		Name: to.Ptr("databaseName"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armcosmos.MongoDBDatabaseGetProperties{
	// 			Resource: &armcosmos.MongoDBDatabaseGetPropertiesResource{
	// 				ID: to.Ptr("databaseName"),
	// 			},
	// 		},
	// 	},
	// }
}
