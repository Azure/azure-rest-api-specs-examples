package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_ListByStreamingJob.json
func ExampleOutputsClient_NewListByStreamingJobPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOutputsClient().NewListByStreamingJobPager("sjrg2157", "sj6458", &armstreamanalytics.OutputsClientListByStreamingJobOptions{Select: nil})
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
		// page.OutputListResult = armstreamanalytics.OutputListResult{
		// 	Value: []*armstreamanalytics.Output{
		// 		{
		// 			Name: to.Ptr("output1755"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg2157/providers/Microsoft.StreamAnalytics/streamingjobs/sj6458/outputs/output1755"),
		// 			Properties: &armstreamanalytics.OutputProperties{
		// 				Datasource: &armstreamanalytics.AzureSQLDatabaseOutputDataSource{
		// 					Type: to.Ptr("Microsoft.Sql/Server/Database"),
		// 					Properties: &armstreamanalytics.AzureSQLDatabaseOutputDataSourceProperties{
		// 						Database: to.Ptr("someDatabase"),
		// 						Server: to.Ptr("someServer"),
		// 						Table: to.Ptr("differentTable"),
		// 						User: to.Ptr("someUser"),
		// 					},
		// 				},
		// 				Etag: to.Ptr("f489d6f3-fcd5-4bcb-b642-81e987ee16d6"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("output958"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg2157/providers/Microsoft.StreamAnalytics/streamingjobs/sj6458/outputs/output958"),
		// 			Properties: &armstreamanalytics.OutputProperties{
		// 				Datasource: &armstreamanalytics.AzureTableOutputDataSource{
		// 					Type: to.Ptr("Microsoft.Storage/Table"),
		// 					Properties: &armstreamanalytics.AzureTableOutputDataSourceProperties{
		// 						AccountName: to.Ptr("someAccountName"),
		// 						BatchSize: to.Ptr[int32](25),
		// 						ColumnsToRemove: []*string{
		// 							to.Ptr("column1"),
		// 							to.Ptr("column2")},
		// 							PartitionKey: to.Ptr("differentPartitionKey"),
		// 							RowKey: to.Ptr("rowKey"),
		// 							Table: to.Ptr("samples"),
		// 						},
		// 					},
		// 					Etag: to.Ptr("ea1d20bf-6cb3-40bc-bc7b-ec3a7fd5977e"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("output1623"),
		// 				Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
		// 				ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg2157/providers/Microsoft.StreamAnalytics/streamingjobs/sj6458/outputs/output1623"),
		// 				Properties: &armstreamanalytics.OutputProperties{
		// 					Datasource: &armstreamanalytics.BlobOutputDataSource{
		// 						Type: to.Ptr("Microsoft.Storage/Blob"),
		// 						Properties: &armstreamanalytics.BlobOutputDataSourceProperties{
		// 							Container: to.Ptr("differentContainer"),
		// 							DateFormat: to.Ptr("yyyy/MM/dd"),
		// 							PathPattern: to.Ptr("{date}/{time}"),
		// 							StorageAccounts: []*armstreamanalytics.StorageAccount{
		// 								{
		// 									AccountName: to.Ptr("someAccountName"),
		// 							}},
		// 							TimeFormat: to.Ptr("HH"),
		// 						},
		// 					},
		// 					Etag: to.Ptr("3a1b2023-79a9-4b33-93e8-f49fc3e573fe"),
		// 					Serialization: &armstreamanalytics.CSVSerialization{
		// 						Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
		// 						Properties: &armstreamanalytics.CSVSerializationProperties{
		// 							Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
		// 							FieldDelimiter: to.Ptr("|"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("output3022"),
		// 				Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
		// 				ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg2157/providers/Microsoft.StreamAnalytics/streamingjobs/sj6458/outputs/output3022"),
		// 				Properties: &armstreamanalytics.OutputProperties{
		// 					Datasource: &armstreamanalytics.DocumentDbOutputDataSource{
		// 						Type: to.Ptr("Microsoft.Storage/DocumentDB"),
		// 						Properties: &armstreamanalytics.DocumentDbOutputDataSourceProperties{
		// 							AccountID: to.Ptr("someAccountId"),
		// 							CollectionNamePattern: to.Ptr("collection"),
		// 							Database: to.Ptr("db01"),
		// 							DocumentID: to.Ptr("documentId"),
		// 							PartitionKey: to.Ptr("differentPartitionKey"),
		// 						},
		// 					},
		// 					Etag: to.Ptr("7849c132-e995-4631-91c3-931606eec432"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("output5195"),
		// 				Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
		// 				ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg2157/providers/Microsoft.StreamAnalytics/streamingjobs/sj6458/outputs/output5195"),
		// 				Properties: &armstreamanalytics.OutputProperties{
		// 					Datasource: &armstreamanalytics.EventHubOutputDataSource{
		// 						Type: to.Ptr("Microsoft.ServiceBus/EventHub"),
		// 						Properties: &armstreamanalytics.EventHubOutputDataSourceProperties{
		// 							ServiceBusNamespace: to.Ptr("sdktest"),
		// 							SharedAccessPolicyName: to.Ptr("RootManageSharedAccessKey"),
		// 							EventHubName: to.Ptr("sdkeventhub"),
		// 							PartitionKey: to.Ptr("differentPartitionKey"),
		// 						},
		// 					},
		// 					Etag: to.Ptr("5020de6b-5bb3-4b88-8606-f11fb3c46185"),
		// 					Serialization: &armstreamanalytics.JSONSerialization{
		// 						Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
		// 						Properties: &armstreamanalytics.JSONSerializationProperties{
		// 							Format: to.Ptr(armstreamanalytics.JSONOutputSerializationFormatLineSeparated),
		// 							Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("output3456"),
		// 				Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
		// 				ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg2157/providers/Microsoft.StreamAnalytics/streamingjobs/sj6458/outputs/output3456"),
		// 				Properties: &armstreamanalytics.OutputProperties{
		// 					Datasource: &armstreamanalytics.ServiceBusQueueOutputDataSource{
		// 						Type: to.Ptr("Microsoft.ServiceBus/Queue"),
		// 						Properties: &armstreamanalytics.ServiceBusQueueOutputDataSourceProperties{
		// 							ServiceBusNamespace: to.Ptr("sdktest"),
		// 							SharedAccessPolicyName: to.Ptr("RootManageSharedAccessKey"),
		// 							PropertyColumns: []*string{
		// 								to.Ptr("column1"),
		// 								to.Ptr("column2")},
		// 								QueueName: to.Ptr("differentQueueName"),
		// 							},
		// 						},
		// 						Etag: to.Ptr("429adaec-a777-4750-8a39-8d0c931d801c"),
		// 						Serialization: &armstreamanalytics.JSONSerialization{
		// 							Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
		// 							Properties: &armstreamanalytics.JSONSerializationProperties{
		// 								Format: to.Ptr(armstreamanalytics.JSONOutputSerializationFormatLineSeparated),
		// 								Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
		// 							},
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("output7886"),
		// 					Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
		// 					ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg2157/providers/Microsoft.StreamAnalytics/streamingjobs/sj6458/outputs/output7886"),
		// 					Properties: &armstreamanalytics.OutputProperties{
		// 						Datasource: &armstreamanalytics.ServiceBusTopicOutputDataSource{
		// 							Type: to.Ptr("Microsoft.ServiceBus/Topic"),
		// 							Properties: &armstreamanalytics.ServiceBusTopicOutputDataSourceProperties{
		// 								ServiceBusNamespace: to.Ptr("sdktest"),
		// 								SharedAccessPolicyName: to.Ptr("RootManageSharedAccessKey"),
		// 								PropertyColumns: []*string{
		// 									to.Ptr("column1"),
		// 									to.Ptr("column2")},
		// 									TopicName: to.Ptr("differentTopicName"),
		// 								},
		// 							},
		// 							Etag: to.Ptr("c1c2007f-45b2-419a-ae7d-4d2148998460"),
		// 							Serialization: &armstreamanalytics.CSVSerialization{
		// 								Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
		// 								Properties: &armstreamanalytics.CSVSerializationProperties{
		// 									Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
		// 									FieldDelimiter: to.Ptr("|"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("output3021"),
		// 						Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
		// 						ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg2157/providers/Microsoft.StreamAnalytics/streamingjobs/sj6458/outputs/output3021"),
		// 						Properties: &armstreamanalytics.OutputProperties{
		// 							Datasource: &armstreamanalytics.PowerBIOutputDataSource{
		// 								Type: to.Ptr("PowerBI"),
		// 								Properties: &armstreamanalytics.PowerBIOutputDataSourceProperties{
		// 									TokenUserDisplayName: to.Ptr("Bob Smith"),
		// 									TokenUserPrincipalName: to.Ptr("bobsmith@contoso.com"),
		// 									Dataset: to.Ptr("differentDataset"),
		// 									GroupID: to.Ptr("ac40305e-3e8d-43ac-8161-c33799f43e95"),
		// 									GroupName: to.Ptr("MyPowerBIGroup"),
		// 									Table: to.Ptr("someTable"),
		// 								},
		// 							},
		// 							Etag: to.Ptr("4a492191-9672-4178-be10-043b9dbd4b9f"),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("output5196"),
		// 						Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
		// 						ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg2157/providers/Microsoft.StreamAnalytics/streamingjobs/sj6458/outputs/output5196"),
		// 						Properties: &armstreamanalytics.OutputProperties{
		// 							Datasource: &armstreamanalytics.AzureDataLakeStoreOutputDataSource{
		// 								Type: to.Ptr("Microsoft.DataLake/Accounts"),
		// 								Properties: &armstreamanalytics.AzureDataLakeStoreOutputDataSourceProperties{
		// 									TokenUserDisplayName: to.Ptr("Bob Smith"),
		// 									TokenUserPrincipalName: to.Ptr("bobsmith@contoso.com"),
		// 									AccountName: to.Ptr("differentaccount"),
		// 									DateFormat: to.Ptr("yyyy/MM/dd"),
		// 									FilePathPrefix: to.Ptr("{date}/{time}"),
		// 									TenantID: to.Ptr("cea4e98b-c798-49e7-8c40-4a2b3beb47dd"),
		// 									TimeFormat: to.Ptr("HH"),
		// 								},
		// 							},
		// 							Etag: to.Ptr("39ab7642-8c1e-48ed-85eb-949068d68002"),
		// 							Serialization: &armstreamanalytics.JSONSerialization{
		// 								Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
		// 								Properties: &armstreamanalytics.JSONSerializationProperties{
		// 									Format: to.Ptr(armstreamanalytics.JSONOutputSerializationFormatLineSeparated),
		// 									Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
		// 								},
		// 							},
		// 						},
		// 				}},
		// 			}
	}
}
