package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/StreamingJob_List_BySubscription_Expand.json
func ExampleStreamingJobsClient_NewListPager_listAllStreamingJobsInASubscriptionAndUseTheExpandODataQueryParameterToExpandInputsOutputsTransformationAndFunctions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStreamingJobsClient().NewListPager(&armstreamanalytics.StreamingJobsClientListOptions{Expand: to.Ptr("inputs,outputs,transformation,functions")})
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
		// page.StreamingJobListResult = armstreamanalytics.StreamingJobListResult{
		// 	Value: []*armstreamanalytics.StreamingJob{
		// 		{
		// 			Name: to.Ptr("sj7804"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7804"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key3": to.Ptr("value3"),
		// 				"randomKey": to.Ptr("randomValue"),
		// 			},
		// 			Properties: &armstreamanalytics.StreamingJobProperties{
		// 				CompatibilityLevel: to.Ptr(armstreamanalytics.CompatibilityLevelOne0),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-11T04:37:11.540Z"); return t}()),
		// 				DataLocale: to.Ptr("en-US"),
		// 				Etag: to.Ptr("8081b2a3-dfe6-457f-8740-1a22d209bf8a"),
		// 				EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](5),
		// 				EventsOutOfOrderMaxDelayInSeconds: to.Ptr[int32](0),
		// 				EventsOutOfOrderPolicy: to.Ptr(armstreamanalytics.EventsOutOfOrderPolicyAdjust),
		// 				Functions: []*armstreamanalytics.Function{
		// 				},
		// 				Inputs: []*armstreamanalytics.Input{
		// 					{
		// 						Name: to.Ptr("inputtest"),
		// 						Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs"),
		// 						ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7804/inputs/inputtest"),
		// 						Properties: &armstreamanalytics.StreamInputProperties{
		// 							Type: to.Ptr("Stream"),
		// 							Etag: to.Ptr("ca88f8fa-605b-4c7f-8695-46f5faa60cd0"),
		// 							Serialization: &armstreamanalytics.JSONSerialization{
		// 								Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
		// 								Properties: &armstreamanalytics.JSONSerializationProperties{
		// 									Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
		// 								},
		// 							},
		// 							Datasource: &armstreamanalytics.BlobStreamInputDataSource{
		// 								Type: to.Ptr("Microsoft.Storage/Blob"),
		// 								Properties: &armstreamanalytics.BlobStreamInputDataSourceProperties{
		// 									Container: to.Ptr("containerName"),
		// 									PathPattern: to.Ptr(""),
		// 									StorageAccounts: []*armstreamanalytics.StorageAccount{
		// 										{
		// 											AccountName: to.Ptr("accountName"),
		// 									}},
		// 								},
		// 							},
		// 						},
		// 				}},
		// 				JobID: to.Ptr("732e4b1d-94a7-43ae-8297-3ad04f1540b9"),
		// 				JobState: to.Ptr("Created"),
		// 				OutputErrorPolicy: to.Ptr(armstreamanalytics.OutputErrorPolicyStop),
		// 				Outputs: []*armstreamanalytics.Output{
		// 					{
		// 						Name: to.Ptr("outputtest"),
		// 						Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
		// 						ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7804/outputs/outputtest"),
		// 						Properties: &armstreamanalytics.OutputProperties{
		// 							Datasource: &armstreamanalytics.AzureSQLDatabaseOutputDataSource{
		// 								Type: to.Ptr("Microsoft.Sql/Server/Database"),
		// 								Properties: &armstreamanalytics.AzureSQLDatabaseOutputDataSourceProperties{
		// 									Database: to.Ptr("databaseName"),
		// 									Server: to.Ptr("serverName"),
		// 									Table: to.Ptr("tableName"),
		// 									User: to.Ptr("userName"),
		// 								},
		// 							},
		// 							Etag: to.Ptr("62097c3c-b503-41ff-a56f-196a9598ab90"),
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SKU: &armstreamanalytics.SKU{
		// 					Name: to.Ptr(armstreamanalytics.SKUNameStandard),
		// 				},
		// 				Transformation: &armstreamanalytics.Transformation{
		// 					Name: to.Ptr("transformationtest"),
		// 					Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/transformations"),
		// 					ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7804/transformations/transformationtest"),
		// 					Properties: &armstreamanalytics.TransformationProperties{
		// 						Etag: to.Ptr("91d8fcbe-60b3-49c3-9f21-9942b95602b8"),
		// 						Query: to.Ptr("Select Id, Name from inputtest"),
		// 						StreamingUnits: to.Ptr[int32](1),
		// 					},
		// 				},
		// 			},
		// 			SKU: &armstreamanalytics.SKU{
		// 				Name: to.Ptr(armstreamanalytics.SKUNameStandard),
		// 				Capacity: to.Ptr[int32](36),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sj7805"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7805"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key3": to.Ptr("value3"),
		// 				"randomKey": to.Ptr("randomValue"),
		// 			},
		// 			Properties: &armstreamanalytics.StreamingJobProperties{
		// 				CompatibilityLevel: to.Ptr(armstreamanalytics.CompatibilityLevelOne0),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-11T04:38:11.540Z"); return t}()),
		// 				DataLocale: to.Ptr("en-US"),
		// 				Etag: to.Ptr("5420059f-e5d7-47d4-be44-40816a4dca7e"),
		// 				EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](5),
		// 				EventsOutOfOrderMaxDelayInSeconds: to.Ptr[int32](0),
		// 				EventsOutOfOrderPolicy: to.Ptr(armstreamanalytics.EventsOutOfOrderPolicyAdjust),
		// 				Functions: []*armstreamanalytics.Function{
		// 				},
		// 				Inputs: []*armstreamanalytics.Input{
		// 					{
		// 						Name: to.Ptr("inputtest"),
		// 						Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs"),
		// 						ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7805/inputs/inputtest"),
		// 						Properties: &armstreamanalytics.StreamInputProperties{
		// 							Type: to.Ptr("Stream"),
		// 							Etag: to.Ptr("45dcf40f-88bb-4776-b5ca-7b10a607cb59"),
		// 							Serialization: &armstreamanalytics.JSONSerialization{
		// 								Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
		// 								Properties: &armstreamanalytics.JSONSerializationProperties{
		// 									Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
		// 								},
		// 							},
		// 							Datasource: &armstreamanalytics.BlobStreamInputDataSource{
		// 								Type: to.Ptr("Microsoft.Storage/Blob"),
		// 								Properties: &armstreamanalytics.BlobStreamInputDataSourceProperties{
		// 									Container: to.Ptr("containerName"),
		// 									PathPattern: to.Ptr(""),
		// 									StorageAccounts: []*armstreamanalytics.StorageAccount{
		// 										{
		// 											AccountName: to.Ptr("accountName"),
		// 									}},
		// 								},
		// 							},
		// 						},
		// 				}},
		// 				JobID: to.Ptr("d8f4041a-0793-433e-a38d-5499d5332113"),
		// 				JobState: to.Ptr("Created"),
		// 				OutputErrorPolicy: to.Ptr(armstreamanalytics.OutputErrorPolicyStop),
		// 				Outputs: []*armstreamanalytics.Output{
		// 					{
		// 						Name: to.Ptr("outputtest"),
		// 						Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
		// 						ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7805/outputs/outputtest"),
		// 						Properties: &armstreamanalytics.OutputProperties{
		// 							Datasource: &armstreamanalytics.AzureSQLDatabaseOutputDataSource{
		// 								Type: to.Ptr("Microsoft.Sql/Server/Database"),
		// 								Properties: &armstreamanalytics.AzureSQLDatabaseOutputDataSourceProperties{
		// 									Database: to.Ptr("databaseName"),
		// 									Server: to.Ptr("serverName"),
		// 									Table: to.Ptr("tableName"),
		// 									User: to.Ptr("userName"),
		// 								},
		// 							},
		// 							Etag: to.Ptr("b09606c8-1b0d-43c9-affb-fac0e18b9481"),
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SKU: &armstreamanalytics.SKU{
		// 					Name: to.Ptr(armstreamanalytics.SKUNameStandard),
		// 				},
		// 				Transformation: &armstreamanalytics.Transformation{
		// 					Name: to.Ptr("transformationtest"),
		// 					Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/transformations"),
		// 					ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3276/providers/Microsoft.StreamAnalytics/streamingjobs/sj7805/transformations/transformationtest"),
		// 					Properties: &armstreamanalytics.TransformationProperties{
		// 						Etag: to.Ptr("288d95c0-204b-4c54-828f-f57aa6896b2a"),
		// 						Query: to.Ptr("Select Id, Name from inputtest"),
		// 						StreamingUnits: to.Ptr[int32](1),
		// 					},
		// 				},
		// 			},
		// 			SKU: &armstreamanalytics.SKU{
		// 				Name: to.Ptr(armstreamanalytics.SKUNameStandard),
		// 				Capacity: to.Ptr[int32](36),
		// 			},
		// 	}},
		// }
	}
}
