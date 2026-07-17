package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBSqlContainerCreateUpdate.json
func ExampleSQLResourcesClient_BeginCreateUpdateSQLContainer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLResourcesClient().BeginCreateUpdateSQLContainer(ctx, "rg1", "ddb1", "databaseName", "containerName", armcosmos.SQLContainerCreateUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.SQLContainerCreateUpdateProperties{
			Resource: &armcosmos.SQLContainerResource{
				ID: to.Ptr("containerName"),
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
					VectorIndexes: []*armcosmos.VectorIndex{
						{
							Path: to.Ptr("/vectorPath1"),
							Type: to.Ptr(armcosmos.VectorIndexTypeFlat),
						},
						{
							Path: to.Ptr("/vectorPath2"),
							Type: to.Ptr(armcosmos.VectorIndexTypeQuantizedFlat),
						},
						{
							Path: to.Ptr("/vectorPath3"),
							Type: to.Ptr(armcosmos.VectorIndexTypeDiskANN),
						},
					},
					FullTextIndexes: []*armcosmos.FullTextIndexPath{
						{
							Path: to.Ptr("/ftPath1"),
						},
						{
							Path: to.Ptr("/ftPath2"),
						},
						{
							Path: to.Ptr("/ftPath3"),
						},
					},
				},
				VectorEmbeddingPolicy: &armcosmos.VectorEmbeddingPolicy{
					VectorEmbeddings: []*armcosmos.VectorEmbedding{
						{
							Path:             to.Ptr("/vectorPath1"),
							DataType:         to.Ptr(armcosmos.VectorDataTypeFloat32),
							Dimensions:       to.Ptr[int32](400),
							DistanceFunction: to.Ptr(armcosmos.DistanceFunctionEuclidean),
						},
						{
							Path:             to.Ptr("/vectorPath2"),
							DataType:         to.Ptr(armcosmos.VectorDataTypeUint8),
							Dimensions:       to.Ptr[int32](512),
							DistanceFunction: to.Ptr(armcosmos.DistanceFunctionCosine),
						},
						{
							Path:             to.Ptr("/vectorPath3"),
							DataType:         to.Ptr(armcosmos.VectorDataTypeInt8),
							Dimensions:       to.Ptr[int32](512),
							DistanceFunction: to.Ptr(armcosmos.DistanceFunctionDotproduct),
						},
					},
				},
				FullTextPolicy: &armcosmos.FullTextPolicy{
					DefaultLanguage: to.Ptr("1033"),
					FullTextPaths: []*armcosmos.FullTextPath{
						{
							Path:     to.Ptr("/ftPath1"),
							Language: to.Ptr("en-US"),
						},
						{
							Path:     to.Ptr("/ftPath2"),
							Language: to.Ptr("fr-FR"),
						},
						{
							Path:     to.Ptr("/ftPath3"),
							Language: to.Ptr("de-DE"),
						},
					},
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
				ClientEncryptionPolicy: &armcosmos.ClientEncryptionPolicy{
					IncludedPaths: []*armcosmos.ClientEncryptionIncludedPath{
						{
							Path:                  to.Ptr("/path"),
							ClientEncryptionKeyID: to.Ptr("keyId"),
							EncryptionAlgorithm:   to.Ptr("AEAD_AES_256_CBC_HMAC_SHA256"),
							EncryptionType:        to.Ptr("Deterministic"),
						},
					},
					PolicyFormatVersion: to.Ptr[int32](2),
				},
				ComputedProperties: []*armcosmos.ComputedProperty{
					{
						Name:  to.Ptr("cp_lowerName"),
						Query: to.Ptr("SELECT VALUE LOWER(c.name) FROM c"),
					},
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
	// res = armcosmos.SQLResourcesClientCreateUpdateSQLContainerResponse{
	// 	SQLContainerGetResults: armcosmos.SQLContainerGetResults{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName/containers/containerName"),
	// 		Name: to.Ptr("containerName"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/sqlContainers"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armcosmos.SQLContainerGetProperties{
	// 			Resource: &armcosmos.SQLContainerGetPropertiesResource{
	// 				ID: to.Ptr("containerName"),
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
	// 					VectorIndexes: []*armcosmos.VectorIndex{
	// 						{
	// 							Path: to.Ptr("/vectorPath1"),
	// 							Type: to.Ptr(armcosmos.VectorIndexTypeFlat),
	// 						},
	// 						{
	// 							Path: to.Ptr("/vectorPath2"),
	// 							Type: to.Ptr(armcosmos.VectorIndexTypeQuantizedFlat),
	// 							QuantizationByteSize: to.Ptr[int64](100),
	// 							VectorIndexShardKey: []*string{
	// 								to.Ptr("/vectorShardKey1"),
	// 							},
	// 						},
	// 						{
	// 							Path: to.Ptr("/vectorPath3"),
	// 							Type: to.Ptr(armcosmos.VectorIndexTypeDiskANN),
	// 							QuantizationByteSize: to.Ptr[int64](100),
	// 							IndexingSearchListSize: to.Ptr[int64](25),
	// 							VectorIndexShardKey: []*string{
	// 								to.Ptr("/vectorShardKey1"),
	// 								to.Ptr("/vectorShardKey2"),
	// 							},
	// 						},
	// 					},
	// 					FullTextIndexes: []*armcosmos.FullTextIndexPath{
	// 						{
	// 							Path: to.Ptr("/ftPath1"),
	// 						},
	// 						{
	// 							Path: to.Ptr("/ftPath2"),
	// 						},
	// 						{
	// 							Path: to.Ptr("/ftPath3"),
	// 						},
	// 					},
	// 				},
	// 				VectorEmbeddingPolicy: &armcosmos.VectorEmbeddingPolicy{
	// 					VectorEmbeddings: []*armcosmos.VectorEmbedding{
	// 						{
	// 							Path: to.Ptr("/vectorPath1"),
	// 							DataType: to.Ptr(armcosmos.VectorDataTypeFloat32),
	// 							Dimensions: to.Ptr[int32](400),
	// 							DistanceFunction: to.Ptr(armcosmos.DistanceFunctionEuclidean),
	// 						},
	// 						{
	// 							Path: to.Ptr("/vectorPath2"),
	// 							DataType: to.Ptr(armcosmos.VectorDataTypeUint8),
	// 							Dimensions: to.Ptr[int32](512),
	// 							DistanceFunction: to.Ptr(armcosmos.DistanceFunctionCosine),
	// 						},
	// 						{
	// 							Path: to.Ptr("/vectorPath3"),
	// 							DataType: to.Ptr(armcosmos.VectorDataTypeInt8),
	// 							Dimensions: to.Ptr[int32](512),
	// 							DistanceFunction: to.Ptr(armcosmos.DistanceFunctionDotproduct),
	// 						},
	// 					},
	// 				},
	// 				FullTextPolicy: &armcosmos.FullTextPolicy{
	// 					DefaultLanguage: to.Ptr("1033"),
	// 					FullTextPaths: []*armcosmos.FullTextPath{
	// 						{
	// 							Path: to.Ptr("/ftPath1"),
	// 							Language: to.Ptr("en-US"),
	// 						},
	// 						{
	// 							Path: to.Ptr("/ftPath2"),
	// 							Language: to.Ptr("fr-FR"),
	// 						},
	// 						{
	// 							Path: to.Ptr("/ftPath3"),
	// 							Language: to.Ptr("de-DE"),
	// 						},
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
	// 				ClientEncryptionPolicy: &armcosmos.ClientEncryptionPolicy{
	// 					IncludedPaths: []*armcosmos.ClientEncryptionIncludedPath{
	// 						{
	// 							Path: to.Ptr("/path"),
	// 							ClientEncryptionKeyID: to.Ptr("keyId"),
	// 							EncryptionAlgorithm: to.Ptr("AEAD_AES_256_CBC_HMAC_SHA256"),
	// 							EncryptionType: to.Ptr("Deterministic"),
	// 						},
	// 					},
	// 					PolicyFormatVersion: to.Ptr[int32](1),
	// 				},
	// 				ComputedProperties: []*armcosmos.ComputedProperty{
	// 					{
	// 						Name: to.Ptr("cp_lowerName"),
	// 						Query: to.Ptr("SELECT VALUE LOWER(c.name) FROM c"),
	// 					},
	// 				},
	// 				Rid: to.Ptr("PD5DALigDgw="),
	// 				Ts: to.Ptr[float32](1459200611),
	// 				Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			},
	// 		},
	// 	},
	// }
}
