package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/011ecc5633300a5eefe43dde748f269d39e96458/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBSqlContainerList.json
func ExampleSQLResourcesClient_NewListSQLContainersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLResourcesClient().NewListSQLContainersPager("rgName", "ddb1", "databaseName", nil)
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
		// page.SQLContainerListResult = armcosmos.SQLContainerListResult{
		// 	Value: []*armcosmos.SQLContainerGetResults{
		// 		{
		// 			Name: to.Ptr("containerName"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/sqlContainers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName/containers/containerName"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcosmos.SQLContainerGetProperties{
		// 				Resource: &armcosmos.SQLContainerGetPropertiesResource{
		// 					Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
		// 					Rid: to.Ptr("PD5DALigDgw="),
		// 					Ts: to.Ptr[float32](1459200611),
		// 					ClientEncryptionPolicy: &armcosmos.ClientEncryptionPolicy{
		// 						IncludedPaths: []*armcosmos.ClientEncryptionIncludedPath{
		// 							{
		// 								Path: to.Ptr("/path"),
		// 								ClientEncryptionKeyID: to.Ptr("keyId"),
		// 								EncryptionAlgorithm: to.Ptr("AEAD_AES_256_CBC_HMAC_SHA256"),
		// 								EncryptionType: to.Ptr("Deterministic"),
		// 						}},
		// 						PolicyFormatVersion: to.Ptr[int32](1),
		// 					},
		// 					ComputedProperties: []*armcosmos.ComputedProperty{
		// 						{
		// 							Name: to.Ptr("cp_lowerName"),
		// 							Query: to.Ptr("SELECT VALUE LOWER(c.name) FROM c"),
		// 					}},
		// 					ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
		// 						ConflictResolutionPath: to.Ptr("/path"),
		// 						Mode: to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
		// 					},
		// 					DefaultTTL: to.Ptr[int32](100),
		// 					ID: to.Ptr("testctn"),
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
		// 				},
		// 				{
		// 					Name: to.Ptr("containerName1"),
		// 					Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/sqlContainers"),
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName/sqlContainers/containerName1"),
		// 					Location: to.Ptr("West US"),
		// 					Tags: map[string]*string{
		// 					},
		// 					Properties: &armcosmos.SQLContainerGetProperties{
		// 						Resource: &armcosmos.SQLContainerGetPropertiesResource{
		// 							Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
		// 							Rid: to.Ptr("PD5DALigDiw="),
		// 							Ts: to.Ptr[float32](1459200611),
		// 							ClientEncryptionPolicy: &armcosmos.ClientEncryptionPolicy{
		// 								IncludedPaths: []*armcosmos.ClientEncryptionIncludedPath{
		// 									{
		// 										Path: to.Ptr("/path"),
		// 										ClientEncryptionKeyID: to.Ptr("keyId"),
		// 										EncryptionAlgorithm: to.Ptr("AEAD_AES_256_CBC_HMAC_SHA256"),
		// 										EncryptionType: to.Ptr("Deterministic"),
		// 								}},
		// 								PolicyFormatVersion: to.Ptr[int32](1),
		// 							},
		// 							ComputedProperties: []*armcosmos.ComputedProperty{
		// 							},
		// 							ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
		// 								ConflictResolutionPath: to.Ptr("/path"),
		// 								Mode: to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
		// 							},
		// 							DefaultTTL: to.Ptr[int32](100),
		// 							FullTextPolicy: &armcosmos.FullTextPolicy{
		// 								DefaultLanguage: to.Ptr("1033"),
		// 								FullTextPaths: []*armcosmos.FullTextPath{
		// 									{
		// 										Path: to.Ptr("/ftPath1"),
		// 										Language: to.Ptr("en-US"),
		// 									},
		// 									{
		// 										Path: to.Ptr("/ftPath2"),
		// 										Language: to.Ptr("fr-FR"),
		// 									},
		// 									{
		// 										Path: to.Ptr("/ftPath3"),
		// 										Language: to.Ptr("de-DE"),
		// 								}},
		// 							},
		// 							ID: to.Ptr("testctn1"),
		// 							IndexingPolicy: &armcosmos.IndexingPolicy{
		// 								Automatic: to.Ptr(true),
		// 								ExcludedPaths: []*armcosmos.ExcludedPath{
		// 								},
		// 								IncludedPaths: []*armcosmos.IncludedPath{
		// 									{
		// 										Path: to.Ptr("/*"),
		// 										Indexes: []*armcosmos.Indexes{
		// 											{
		// 												DataType: to.Ptr(armcosmos.DataTypeString),
		// 												Kind: to.Ptr(armcosmos.IndexKindRange),
		// 												Precision: to.Ptr[int32](-1),
		// 											},
		// 											{
		// 												DataType: to.Ptr(armcosmos.DataTypeNumber),
		// 												Kind: to.Ptr(armcosmos.IndexKindRange),
		// 												Precision: to.Ptr[int32](-1),
		// 										}},
		// 								}},
		// 								IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
		// 								VectorIndexes: []*armcosmos.VectorIndex{
		// 									{
		// 										Type: to.Ptr(armcosmos.VectorIndexTypeFlat),
		// 										Path: to.Ptr("/vectorPath1"),
		// 								}},
		// 							},
		// 							PartitionKey: &armcosmos.ContainerPartitionKey{
		// 								Kind: to.Ptr(armcosmos.PartitionKindMultiHash),
		// 								Paths: []*string{
		// 									to.Ptr("/AccountNumber"),
		// 									to.Ptr("/AccountLocation")},
		// 									Version: to.Ptr[int32](2),
		// 								},
		// 								UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
		// 									UniqueKeys: []*armcosmos.UniqueKey{
		// 										{
		// 											Paths: []*string{
		// 												to.Ptr("/testPath")},
		// 										}},
		// 									},
		// 									VectorEmbeddingPolicy: &armcosmos.VectorEmbeddingPolicy{
		// 										VectorEmbeddings: []*armcosmos.VectorEmbedding{
		// 											{
		// 												Path: to.Ptr("/vectorPath1"),
		// 												DataType: to.Ptr(armcosmos.VectorDataTypeFloat32),
		// 												Dimensions: to.Ptr[int32](400),
		// 												DistanceFunction: to.Ptr(armcosmos.DistanceFunctionEuclidean),
		// 										}},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("containerName2"),
		// 							Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/sqlContainers"),
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName/sqlContainers/containerName2"),
		// 							Location: to.Ptr("West US"),
		// 							Tags: map[string]*string{
		// 							},
		// 							Properties: &armcosmos.SQLContainerGetProperties{
		// 								Resource: &armcosmos.SQLContainerGetPropertiesResource{
		// 									Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
		// 									Rid: to.Ptr("PD5DALigDhw="),
		// 									Ts: to.Ptr[float32](1459200611),
		// 									ClientEncryptionPolicy: &armcosmos.ClientEncryptionPolicy{
		// 										IncludedPaths: []*armcosmos.ClientEncryptionIncludedPath{
		// 											{
		// 												Path: to.Ptr("/path"),
		// 												ClientEncryptionKeyID: to.Ptr("keyId"),
		// 												EncryptionAlgorithm: to.Ptr("AEAD_AES_256_CBC_HMAC_SHA256"),
		// 												EncryptionType: to.Ptr("Deterministic"),
		// 										}},
		// 										PolicyFormatVersion: to.Ptr[int32](1),
		// 									},
		// 									ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
		// 										ConflictResolutionPath: to.Ptr("/path"),
		// 										Mode: to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
		// 									},
		// 									DefaultTTL: to.Ptr[int32](100),
		// 									FullTextPolicy: &armcosmos.FullTextPolicy{
		// 										DefaultLanguage: to.Ptr("1033"),
		// 										FullTextPaths: []*armcosmos.FullTextPath{
		// 											{
		// 												Path: to.Ptr("/ftPath1"),
		// 												Language: to.Ptr("en-US"),
		// 											},
		// 											{
		// 												Path: to.Ptr("/ftPath2"),
		// 												Language: to.Ptr("fr-FR"),
		// 											},
		// 											{
		// 												Path: to.Ptr("/ftPath3"),
		// 												Language: to.Ptr("de-DE"),
		// 										}},
		// 									},
		// 									ID: to.Ptr("testctn2"),
		// 									IndexingPolicy: &armcosmos.IndexingPolicy{
		// 										Automatic: to.Ptr(true),
		// 										ExcludedPaths: []*armcosmos.ExcludedPath{
		// 										},
		// 										IncludedPaths: []*armcosmos.IncludedPath{
		// 											{
		// 												Path: to.Ptr("/*"),
		// 												Indexes: []*armcosmos.Indexes{
		// 													{
		// 														DataType: to.Ptr(armcosmos.DataTypeString),
		// 														Kind: to.Ptr(armcosmos.IndexKindRange),
		// 														Precision: to.Ptr[int32](-1),
		// 													},
		// 													{
		// 														DataType: to.Ptr(armcosmos.DataTypeNumber),
		// 														Kind: to.Ptr(armcosmos.IndexKindRange),
		// 														Precision: to.Ptr[int32](-1),
		// 												}},
		// 										}},
		// 										IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
		// 										VectorIndexes: []*armcosmos.VectorIndex{
		// 											{
		// 												Type: to.Ptr(armcosmos.VectorIndexTypeFlat),
		// 												Path: to.Ptr("/vectorPath1"),
		// 											},
		// 											{
		// 												Type: to.Ptr(armcosmos.VectorIndexTypeQuantizedFlat),
		// 												Path: to.Ptr("/vectorPath2"),
		// 											},
		// 											{
		// 												Type: to.Ptr(armcosmos.VectorIndexTypeDiskANN),
		// 												Path: to.Ptr("/vectorPath3"),
		// 										}},
		// 									},
		// 									PartitionKey: &armcosmos.ContainerPartitionKey{
		// 										Kind: to.Ptr(armcosmos.PartitionKindHash),
		// 										Paths: []*string{
		// 											to.Ptr("/_partitionKey")},
		// 											SystemKey: to.Ptr(true),
		// 											Version: to.Ptr[int32](2),
		// 										},
		// 										UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
		// 											UniqueKeys: []*armcosmos.UniqueKey{
		// 												{
		// 													Paths: []*string{
		// 														to.Ptr("/testPath")},
		// 												}},
		// 											},
		// 											VectorEmbeddingPolicy: &armcosmos.VectorEmbeddingPolicy{
		// 												VectorEmbeddings: []*armcosmos.VectorEmbedding{
		// 													{
		// 														Path: to.Ptr("/vectorPath1"),
		// 														DataType: to.Ptr(armcosmos.VectorDataTypeFloat32),
		// 														Dimensions: to.Ptr[int32](400),
		// 														DistanceFunction: to.Ptr(armcosmos.DistanceFunctionEuclidean),
		// 													},
		// 													{
		// 														Path: to.Ptr("/vectorPath2"),
		// 														DataType: to.Ptr(armcosmos.VectorDataTypeUint8),
		// 														Dimensions: to.Ptr[int32](512),
		// 														DistanceFunction: to.Ptr(armcosmos.DistanceFunctionCosine),
		// 													},
		// 													{
		// 														Path: to.Ptr("/vectorPath3"),
		// 														DataType: to.Ptr(armcosmos.VectorDataTypeInt8),
		// 														Dimensions: to.Ptr[int32](512),
		// 														DistanceFunction: to.Ptr(armcosmos.DistanceFunctionDotproduct),
		// 												}},
		// 											},
		// 										},
		// 									},
		// 							}},
		// 						}
	}
}
