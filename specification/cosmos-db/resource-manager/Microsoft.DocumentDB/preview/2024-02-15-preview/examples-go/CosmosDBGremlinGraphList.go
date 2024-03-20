package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-02-15-preview/examples/CosmosDBGremlinGraphList.json
func ExampleGremlinResourcesClient_NewListGremlinGraphsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGremlinResourcesClient().NewListGremlinGraphsPager("rgName", "ddb1", "databaseName", nil)
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
		// page.GremlinGraphListResult = armcosmos.GremlinGraphListResult{
		// 	Value: []*armcosmos.GremlinGraphGetResults{
		// 		{
		// 			Name: to.Ptr("testgrf"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/apis/databases/graphs"),
		// 			ID: to.Ptr("testgrf"),
		// 			Properties: &armcosmos.GremlinGraphGetProperties{
		// 				Resource: &armcosmos.GremlinGraphGetPropertiesResource{
		// 					Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
		// 					Rid: to.Ptr("PD5DALigDgw="),
		// 					Ts: to.Ptr[float32](1459200611),
		// 					ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
		// 						ConflictResolutionPath: to.Ptr("/path"),
		// 						Mode: to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
		// 					},
		// 					DefaultTTL: to.Ptr[int32](100),
		// 					ID: to.Ptr("testgrf"),
		// 					IndexingPolicy: &armcosmos.IndexingPolicy{
		// 						Automatic: to.Ptr(true),
		// 						ExcludedPaths: []*armcosmos.ExcludedPath{
		// 						},
		// 						IncludedPaths: []*armcosmos.IncludedPath{
		// 							{
		// 								Path: to.Ptr("/*"),
		// 								Indexes: []*armcosmos.Indexes{
		// 									{
		// 										DataType: to.Ptr(armcosmos.DataTypeString),
		// 										Kind: to.Ptr(armcosmos.IndexKindRange),
		// 										Precision: to.Ptr[int32](-1),
		// 									},
		// 									{
		// 										DataType: to.Ptr(armcosmos.DataTypeNumber),
		// 										Kind: to.Ptr(armcosmos.IndexKindRange),
		// 										Precision: to.Ptr[int32](-1),
		// 								}},
		// 						}},
		// 						IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
		// 					},
		// 					PartitionKey: &armcosmos.ContainerPartitionKey{
		// 						Kind: to.Ptr(armcosmos.PartitionKindHash),
		// 						Paths: []*string{
		// 							to.Ptr("/AccountNumber")},
		// 						},
		// 						UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
		// 							UniqueKeys: []*armcosmos.UniqueKey{
		// 								{
		// 									Paths: []*string{
		// 										to.Ptr("/testPath")},
		// 								}},
		// 							},
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
