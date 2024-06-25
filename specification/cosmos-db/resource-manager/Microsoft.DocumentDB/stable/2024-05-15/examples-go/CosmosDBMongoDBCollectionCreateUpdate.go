package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ec7ee8842bf615c2f0354bf8b5b8725fdac9454a/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/CosmosDBMongoDBCollectionCreateUpdate.json
func ExampleMongoDBResourcesClient_BeginCreateUpdateMongoDBCollection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMongoDBResourcesClient().BeginCreateUpdateMongoDBCollection(ctx, "rg1", "ddb1", "databaseName", "collectionName", armcosmos.MongoDBCollectionCreateUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.MongoDBCollectionCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{},
			Resource: &armcosmos.MongoDBCollectionResource{
				ID: to.Ptr("collectionName"),
				Indexes: []*armcosmos.MongoIndex{
					{
						Key: &armcosmos.MongoIndexKeys{
							Keys: []*string{
								to.Ptr("_ts")},
						},
						Options: &armcosmos.MongoIndexOptions{
							ExpireAfterSeconds: to.Ptr[int32](100),
							Unique:             to.Ptr(true),
						},
					},
					{
						Key: &armcosmos.MongoIndexKeys{
							Keys: []*string{
								to.Ptr("_id")},
						},
					}},
				ShardKey: map[string]*string{
					"testKey": to.Ptr("Hash"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MongoDBCollectionGetResults = armcosmos.MongoDBCollectionGetResults{
	// 	Name: to.Ptr("collectionName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/mongodbCollections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/mongodbDatabases/databaseName/mongodbCollections/collectionName"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.MongoDBCollectionGetProperties{
	// 		Resource: &armcosmos.MongoDBCollectionGetPropertiesResource{
	// 			ID: to.Ptr("collectionName"),
	// 			Indexes: []*armcosmos.MongoIndex{
	// 				{
	// 					Key: &armcosmos.MongoIndexKeys{
	// 						Keys: []*string{
	// 							to.Ptr("_ts")},
	// 						},
	// 						Options: &armcosmos.MongoIndexOptions{
	// 							ExpireAfterSeconds: to.Ptr[int32](100),
	// 							Unique: to.Ptr(true),
	// 						},
	// 					},
	// 					{
	// 						Key: &armcosmos.MongoIndexKeys{
	// 							Keys: []*string{
	// 								to.Ptr("_id")},
	// 							},
	// 					}},
	// 					ShardKey: map[string]*string{
	// 						"testKey": to.Ptr("Hash"),
	// 					},
	// 				},
	// 			},
	// 		}
}
