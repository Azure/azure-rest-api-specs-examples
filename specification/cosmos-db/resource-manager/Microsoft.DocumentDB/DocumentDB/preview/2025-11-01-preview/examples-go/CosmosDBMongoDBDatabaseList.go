package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBMongoDBDatabaseList.json
func ExampleMongoDBResourcesClient_NewListMongoDBDatabasesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMongoDBResourcesClient().NewListMongoDBDatabasesPager("rgName", "ddb1", nil)
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
		// page = armcosmos.MongoDBResourcesClientListMongoDBDatabasesResponse{
		// 	MongoDBDatabaseListResult: armcosmos.MongoDBDatabaseListResult{
		// 		Value: []*armcosmos.MongoDBDatabaseGetResults{
		// 			{
		// 				Name: to.Ptr("databaseName"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/mongodbDatabases/databaseName"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armcosmos.MongoDBDatabaseGetProperties{
		// 					Resource: &armcosmos.MongoDBDatabaseGetPropertiesResource{
		// 						ID: to.Ptr("databaseName"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
