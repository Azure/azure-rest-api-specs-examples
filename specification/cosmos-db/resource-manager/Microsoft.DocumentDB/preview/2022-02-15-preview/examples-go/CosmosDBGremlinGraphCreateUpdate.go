package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBGremlinGraphCreateUpdate.json
func ExampleGremlinResourcesClient_BeginCreateUpdateGremlinGraph() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewGremlinResourcesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateUpdateGremlinGraph(ctx,
		"rg1",
		"ddb1",
		"databaseName",
		"graphName",
		armcosmos.GremlinGraphCreateUpdateParameters{
			Location: to.Ptr("West US"),
			Tags:     map[string]*string{},
			Properties: &armcosmos.GremlinGraphCreateUpdateProperties{
				Options: &armcosmos.CreateUpdateOptions{},
				Resource: &armcosmos.GremlinGraphResource{
					ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
						ConflictResolutionPath: to.Ptr("/path"),
						Mode:                   to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
					},
					DefaultTTL: to.Ptr[int32](100),
					ID:         to.Ptr("graphName"),
					IndexingPolicy: &armcosmos.IndexingPolicy{
						Automatic:     to.Ptr(true),
						ExcludedPaths: []*armcosmos.ExcludedPath{},
						IncludedPaths: []*armcosmos.IncludedPath{
							{
								Path: to.Ptr("/*"),
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
