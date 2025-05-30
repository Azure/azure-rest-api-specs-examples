package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/011ecc5633300a5eefe43dde748f269d39e96458/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBMongoDBCollectionList.json
func ExampleMongoDBResourcesClient_NewListMongoDBCollectionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMongoDBResourcesClient().NewListMongoDBCollectionsPager("rgName", "ddb1", "databaseName", nil)
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
		// page.MongoDBCollectionListResult = armcosmos.MongoDBCollectionListResult{
		// 	Value: []*armcosmos.MongoDBCollectionGetResults{
		// 		{
		// 			Name: to.Ptr("collectionName"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/mongodbCollections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/mongodbDatabases/databaseName/mongodbCollections/collectionName"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcosmos.MongoDBCollectionGetProperties{
		// 				Resource: &armcosmos.MongoDBCollectionGetPropertiesResource{
		// 					ID: to.Ptr("testcoll"),
		// 					Indexes: []*armcosmos.MongoIndex{
		// 						{
		// 							Key: &armcosmos.MongoIndexKeys{
		// 								Keys: []*string{
		// 									to.Ptr("testKey")},
		// 								},
		// 								Options: &armcosmos.MongoIndexOptions{
		// 									ExpireAfterSeconds: to.Ptr[int32](100),
		// 									Unique: to.Ptr(true),
		// 								},
		// 						}},
		// 						ShardKey: map[string]*string{
		// 							"testKey": to.Ptr("Hash"),
		// 						},
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
