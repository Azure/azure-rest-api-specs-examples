package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBSqlMaterializedViewCreateUpdate.json
func ExampleSQLResourcesClient_BeginCreateUpdateSQLContainer_cosmosDbSqlMaterializedViewCreateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLResourcesClient().BeginCreateUpdateSQLContainer(ctx, "rg1", "ddb1", "databaseName", "mvContainerName", armcosmos.SQLContainerCreateUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.SQLContainerCreateUpdateProperties{
			Resource: &armcosmos.SQLContainerResource{
				ID: to.Ptr("mvContainerName"),
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
						to.Ptr("/mvpk"),
					},
					Kind: to.Ptr(armcosmos.PartitionKindHash),
				},
				MaterializedViewDefinition: &armcosmos.MaterializedViewDefinition{
					SourceCollectionID:       to.Ptr("sourceContainerName"),
					Definition:               to.Ptr("select * from ROOT"),
					ThroughputBucketForBuild: to.Ptr[int32](1),
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
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.SQLResourcesClientCreateUpdateSQLContainerResponse{
	// 	SQLContainerGetResults: &armcosmos.SQLContainerGetResults{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName/containers/mvContainerName"),
	// 		Name: to.Ptr("mvContainerName"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/sqlContainers"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armcosmos.SQLContainerGetProperties{
	// 			Resource: &armcosmos.SQLContainerGetPropertiesResource{
	// 				ID: to.Ptr("mvContainerName"),
	// 				IndexingPolicy: &armcosmos.IndexingPolicy{
	// 					IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
	// 					Automatic: to.Ptr(true),
	// 					IncludedPaths: []*armcosmos.IncludedPath{
	// 						{
	// 							Path: to.Ptr("/*"),
	// 						},
	// 					},
	// 					ExcludedPaths: []*armcosmos.ExcludedPath{
	// 						{
	// 							Path: to.Ptr("/\"_etag\"/?"),
	// 						},
	// 					},
	// 				},
	// 				PartitionKey: &armcosmos.ContainerPartitionKey{
	// 					Paths: []*string{
	// 						to.Ptr("/mvpk"),
	// 					},
	// 					Kind: to.Ptr(armcosmos.PartitionKindHash),
	// 				},
	// 				DefaultTTL: to.Ptr[int32](-1),
	// 				UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
	// 					UniqueKeys: []*armcosmos.UniqueKey{
	// 					},
	// 				},
	// 				ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
	// 					Mode: to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
	// 					ConflictResolutionPath: to.Ptr("/_ts"),
	// 					ConflictResolutionProcedure: to.Ptr(""),
	// 				},
	// 				MaterializedViewDefinition: &armcosmos.MaterializedViewDefinition{
	// 					SourceCollectionRid: to.Ptr("vb0sn6nEu9A="),
	// 					SourceCollectionID: to.Ptr("sourceContainerName"),
	// 					Definition: to.Ptr("select * from ROOT"),
	// 					ThroughputBucketForBuild: to.Ptr[int32](1),
	// 				},
	// 				Rid: to.Ptr("vb0sn8MDxLw="),
	// 				Ts: to.Ptr[float32](1671427656),
	// 				Etag: to.Ptr("\"00000800-0000-0200-0000-639ff6480000\""),
	// 			},
	// 		},
	// 	},
	// }
}
