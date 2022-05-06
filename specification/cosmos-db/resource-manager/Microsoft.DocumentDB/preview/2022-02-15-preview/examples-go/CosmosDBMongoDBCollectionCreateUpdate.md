Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcosmos%2Farmcosmos%2Fv0.5.0/sdk/resourcemanager/cosmos/armcosmos/README.md) on how to add the SDK to your project and authenticate.

```go
package armcosmos_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBMongoDBCollectionCreateUpdate.json
func ExampleMongoDBResourcesClient_BeginCreateUpdateMongoDBCollection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcosmos.NewMongoDBResourcesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateUpdateMongoDBCollection(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<database-name>",
		"<collection-name>",
		armcosmos.MongoDBCollectionCreateUpdateParameters{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armcosmos.MongoDBCollectionCreateUpdateProperties{
				Options: &armcosmos.CreateUpdateOptions{},
				Resource: &armcosmos.MongoDBCollectionResource{
					AnalyticalStorageTTL: to.Ptr[int32](500),
					ID:                   to.Ptr("<id>"),
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
		&armcosmos.MongoDBResourcesClientBeginCreateUpdateMongoDBCollectionOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
