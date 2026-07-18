package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBGremlinGraphList.json
func ExampleGremlinResourcesClient_NewListGremlinGraphsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
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
		// page = armcosmos.GremlinResourcesClientListGremlinGraphsResponse{
		// 	GremlinGraphListResult: armcosmos.GremlinGraphListResult{
		// 		Value: []*armcosmos.GremlinGraphGetResults{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/gremlinDatabases/databaseName/graphs/testgrf"),
		// 				Name: to.Ptr("testgrf"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/gremlinDatabases/graphs"),
		// 				Properties: &armcosmos.GremlinGraphGetProperties{
		// 					Resource: &armcosmos.GremlinGraphGetPropertiesResource{
		// 						ID: to.Ptr("testgrf"),
		// 						IndexingPolicy: &armcosmos.IndexingPolicy{
		// 							IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
		// 							Automatic: to.Ptr(true),
		// 							IncludedPaths: []*armcosmos.IncludedPath{
		// 								{
		// 									Path: to.Ptr("/*"),
		// 									Indexes: []*armcosmos.Indexes{
		// 										{
		// 											Kind: to.Ptr(armcosmos.IndexKindRange),
		// 											DataType: to.Ptr(armcosmos.DataTypeString),
		// 											Precision: to.Ptr[int32](-1),
		// 										},
		// 										{
		// 											Kind: to.Ptr(armcosmos.IndexKindRange),
		// 											DataType: to.Ptr(armcosmos.DataTypeNumber),
		// 											Precision: to.Ptr[int32](-1),
		// 										},
		// 									},
		// 								},
		// 							},
		// 							ExcludedPaths: []*armcosmos.ExcludedPath{
		// 							},
		// 						},
		// 						PartitionKey: &armcosmos.ContainerPartitionKey{
		// 							Paths: []*string{
		// 								to.Ptr("/AccountNumber"),
		// 							},
		// 							Kind: to.Ptr(armcosmos.PartitionKindHash),
		// 						},
		// 						DefaultTTL: to.Ptr[int32](100),
		// 						UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
		// 							UniqueKeys: []*armcosmos.UniqueKey{
		// 								{
		// 									Paths: []*string{
		// 										to.Ptr("/testPath"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 						ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
		// 							Mode: to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
		// 							ConflictResolutionPath: to.Ptr("/path"),
		// 						},
		// 						Rid: to.Ptr("PD5DALigDgw="),
		// 						Ts: to.Ptr[float32](1459200611),
		// 						Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
