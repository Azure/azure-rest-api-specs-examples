package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBMongoDBCollectionCreateUpdate.json
func ExampleMongoDBResourcesClient_BeginCreateUpdateMongoDBCollection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewMongoDBResourcesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateUpdateMongoDBCollection(ctx,
		"rg1",
		"ddb1",
		"databaseName",
		"collectionName",
		armcosmos.MongoDBCollectionCreateUpdateParameters{
			Location: to.Ptr("West US"),
			Tags:     map[string]*string{},
			Properties: &armcosmos.MongoDBCollectionCreateUpdateProperties{
				Options: &armcosmos.CreateUpdateOptions{},
				Resource: &armcosmos.MongoDBCollectionResource{
					AnalyticalStorageTTL: to.Ptr[int32](500),
					ID:                   to.Ptr("collectionName"),
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
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
