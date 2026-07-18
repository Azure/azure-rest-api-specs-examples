package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBGremlinGraphCreateUpdate.json
func ExampleGremlinResourcesClient_BeginCreateUpdateGremlinGraph() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginCreateUpdateGremlinGraph(ctx, "rg1", "ddb1", "databaseName", "graphName", armcosmos.GremlinGraphCreateUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.GremlinGraphCreateUpdateProperties{
			Resource: &armcosmos.GremlinGraphResource{
				ID: to.Ptr("graphName"),
				IndexingPolicy: &armcosmos.IndexingPolicy{
					IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
					Automatic:    to.Ptr(true),
					IncludedPaths: []*armcosmos.IncludedPath{
						{
							Path: to.Ptr("/*"),
							Indexes: []*armcosmos.Indexes{
								{
									Kind:      to.Ptr(armcosmos.IndexKindRange),
									DataType:  to.Ptr(armcosmos.DataTypeString),
									Precision: to.Ptr[int32](-1),
								},
								{
									Kind:      to.Ptr(armcosmos.IndexKindRange),
									DataType:  to.Ptr(armcosmos.DataTypeNumber),
									Precision: to.Ptr[int32](-1),
								},
							},
						},
					},
					ExcludedPaths: []*armcosmos.ExcludedPath{},
				},
				PartitionKey: &armcosmos.ContainerPartitionKey{
					Paths: []*string{
						to.Ptr("/AccountNumber"),
					},
					Kind: to.Ptr(armcosmos.PartitionKindHash),
				},
				DefaultTTL: to.Ptr[int32](100),
				UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
					UniqueKeys: []*armcosmos.UniqueKey{
						{
							Paths: []*string{
								to.Ptr("/testPath"),
							},
						},
					},
				},
				ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
					Mode:                   to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
					ConflictResolutionPath: to.Ptr("/path"),
				},
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
	// res = armcosmos.GremlinResourcesClientCreateUpdateGremlinGraphResponse{
	// 	GremlinGraphGetResults: armcosmos.GremlinGraphGetResults{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/gremlinDatabases/databaseName/gremlinGraphs/graphName"),
	// 		Name: to.Ptr("graphName"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/gremlinDatabases/gremlinGraphs"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armcosmos.GremlinGraphGetProperties{
	// 			Resource: &armcosmos.GremlinGraphGetPropertiesResource{
	// 				ID: to.Ptr("graphName"),
	// 				IndexingPolicy: &armcosmos.IndexingPolicy{
	// 					IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
	// 					Automatic: to.Ptr(true),
	// 					IncludedPaths: []*armcosmos.IncludedPath{
	// 						{
	// 							Path: to.Ptr("/*"),
	// 							Indexes: []*armcosmos.Indexes{
	// 								{
	// 									Kind: to.Ptr(armcosmos.IndexKindRange),
	// 									DataType: to.Ptr(armcosmos.DataTypeString),
	// 									Precision: to.Ptr[int32](-1),
	// 								},
	// 								{
	// 									Kind: to.Ptr(armcosmos.IndexKindRange),
	// 									DataType: to.Ptr(armcosmos.DataTypeNumber),
	// 									Precision: to.Ptr[int32](-1),
	// 								},
	// 							},
	// 						},
	// 					},
	// 					ExcludedPaths: []*armcosmos.ExcludedPath{
	// 					},
	// 				},
	// 				PartitionKey: &armcosmos.ContainerPartitionKey{
	// 					Paths: []*string{
	// 						to.Ptr("/AccountNumber"),
	// 					},
	// 					Kind: to.Ptr(armcosmos.PartitionKindHash),
	// 				},
	// 				DefaultTTL: to.Ptr[int32](100),
	// 				UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
	// 					UniqueKeys: []*armcosmos.UniqueKey{
	// 						{
	// 							Paths: []*string{
	// 								to.Ptr("/testPath"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 				ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
	// 					Mode: to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
	// 					ConflictResolutionPath: to.Ptr("/path"),
	// 				},
	// 				Rid: to.Ptr("PD5DALigDgw="),
	// 				Ts: to.Ptr[float32](1459200611),
	// 				Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			},
	// 		},
	// 	},
	// }
}
