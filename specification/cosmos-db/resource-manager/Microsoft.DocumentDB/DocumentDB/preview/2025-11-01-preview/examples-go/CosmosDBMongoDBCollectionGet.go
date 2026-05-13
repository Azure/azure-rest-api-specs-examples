package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBMongoDBCollectionGet.json
func ExampleMongoDBResourcesClient_GetMongoDBCollection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMongoDBResourcesClient().GetMongoDBCollection(ctx, "rgName", "ddb1", "databaseName", "collectionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.MongoDBResourcesClientGetMongoDBCollectionResponse{
	// 	MongoDBCollectionGetResults: &armcosmos.MongoDBCollectionGetResults{
	// 		Name: to.Ptr("collectionName"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/mongodbCollections"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/mongodbDatabases/databaseName/mongodbCollections/collectionName"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armcosmos.MongoDBCollectionGetProperties{
	// 			Resource: &armcosmos.MongoDBCollectionGetPropertiesResource{
	// 				Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 				Rid: to.Ptr("PD5DALigDgw="),
	// 				Ts: to.Ptr[float32](1459200611),
	// 				AnalyticalStorageTTL: to.Ptr[int32](500),
	// 				ID: to.Ptr("testcoll"),
	// 				Indexes: []*armcosmos.MongoIndex{
	// 					{
	// 						Key: &armcosmos.MongoIndexKeys{
	// 							Keys: []*string{
	// 								to.Ptr("testKey"),
	// 							},
	// 						},
	// 						Options: &armcosmos.MongoIndexOptions{
	// 							ExpireAfterSeconds: to.Ptr[int32](100),
	// 							Unique: to.Ptr(true),
	// 						},
	// 					},
	// 				},
	// 				ShardKey: map[string]*string{
	// 					"testKey": to.Ptr("Hash"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
