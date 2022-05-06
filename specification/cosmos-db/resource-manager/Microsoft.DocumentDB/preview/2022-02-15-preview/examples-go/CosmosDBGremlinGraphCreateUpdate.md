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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBGremlinGraphCreateUpdate.json
func ExampleGremlinResourcesClient_BeginCreateUpdateGremlinGraph() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcosmos.NewGremlinResourcesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateUpdateGremlinGraph(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<database-name>",
		"<graph-name>",
		armcosmos.GremlinGraphCreateUpdateParameters{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armcosmos.GremlinGraphCreateUpdateProperties{
				Options: &armcosmos.CreateUpdateOptions{},
				Resource: &armcosmos.GremlinGraphResource{
					ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
						ConflictResolutionPath: to.Ptr("<conflict-resolution-path>"),
						Mode:                   to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
					},
					DefaultTTL: to.Ptr[int32](100),
					ID:         to.Ptr("<id>"),
					IndexingPolicy: &armcosmos.IndexingPolicy{
						Automatic:     to.Ptr(true),
						ExcludedPaths: []*armcosmos.ExcludedPath{},
						IncludedPaths: []*armcosmos.IncludedPath{
							{
								Path: to.Ptr("<path>"),
								Indexes: []*armcosmos.Indexes{
									{
										DataType:  to.Ptr(armcosmos.DataTypeString),
										Kind:      to.Ptr(armcosmos.IndexKindRange),
										Precision: to.Ptr[int32](-1),
									},
									{
										DataType:  to.Ptr(armcosmos.DataTypeNumber),
										Kind:      to.Ptr(armcosmos.IndexKindRange),
										Precision: to.Ptr[int32](-1),
									}},
							}},
						IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
					},
					PartitionKey: &armcosmos.ContainerPartitionKey{
						Kind: to.Ptr(armcosmos.PartitionKindHash),
						Paths: []*string{
							to.Ptr("/AccountNumber")},
					},
					UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
						UniqueKeys: []*armcosmos.UniqueKey{
							{
								Paths: []*string{
									to.Ptr("/testPath")},
							}},
					},
				},
			},
		},
		&armcosmos.GremlinResourcesClientBeginCreateUpdateGremlinGraphOptions{ResumeToken: ""})
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
