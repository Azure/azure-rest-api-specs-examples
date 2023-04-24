package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/CosmosDBGremlinGraphGet.json
func ExampleGremlinResourcesClient_GetGremlinGraph() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGremlinResourcesClient().GetGremlinGraph(ctx, "rgName", "ddb1", "databaseName", "graphName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GremlinGraphGetResults = armcosmos.GremlinGraphGetResults{
	// 	Name: to.Ptr("graphName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/apis/databases/graphs"),
	// 	ID: to.Ptr("graphName"),
	// 	Properties: &armcosmos.GremlinGraphGetProperties{
	// 		Resource: &armcosmos.GremlinGraphGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
	// 				ConflictResolutionPath: to.Ptr("/path"),
	// 				Mode: to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
	// 			},
	// 			DefaultTTL: to.Ptr[int32](100),
	// 			ID: to.Ptr("graphName"),
	// 			IndexingPolicy: &armcosmos.IndexingPolicy{
	// 				Automatic: to.Ptr(true),
	// 				ExcludedPaths: []*armcosmos.ExcludedPath{
	// 				},
	// 				IncludedPaths: []*armcosmos.IncludedPath{
	// 					{
	// 						Path: to.Ptr("/*"),
	// 						Indexes: []*armcosmos.Indexes{
	// 							{
	// 								DataType: to.Ptr(armcosmos.DataTypeString),
	// 								Kind: to.Ptr(armcosmos.IndexKindRange),
	// 								Precision: to.Ptr[int32](-1),
	// 							},
	// 							{
	// 								DataType: to.Ptr(armcosmos.DataTypeNumber),
	// 								Kind: to.Ptr(armcosmos.IndexKindRange),
	// 								Precision: to.Ptr[int32](-1),
	// 						}},
	// 				}},
	// 				IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
	// 			},
	// 			PartitionKey: &armcosmos.ContainerPartitionKey{
	// 				Kind: to.Ptr(armcosmos.PartitionKindHash),
	// 				Paths: []*string{
	// 					to.Ptr("/AccountNumber")},
	// 				},
	// 				UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
	// 					UniqueKeys: []*armcosmos.UniqueKey{
	// 						{
	// 							Paths: []*string{
	// 								to.Ptr("/testPath")},
	// 						}},
	// 					},
	// 				},
	// 			},
	// 		}
}
