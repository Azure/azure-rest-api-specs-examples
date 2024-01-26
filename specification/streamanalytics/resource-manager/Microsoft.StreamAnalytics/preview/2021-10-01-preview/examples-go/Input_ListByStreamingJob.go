package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_ListByStreamingJob.json
func ExampleInputsClient_NewListByStreamingJobPager_listAllInputsInAStreamingJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInputsClient().NewListByStreamingJobPager("sjrg8440", "sj9597", &armstreamanalytics.InputsClientListByStreamingJobOptions{Select: nil})
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
		// page.InputListResult = armstreamanalytics.InputListResult{
		// 	Value: []*armstreamanalytics.Input{
		// 		{
		// 			Name: to.Ptr("input7225"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg8440/providers/Microsoft.StreamAnalytics/streamingjobs/sj9597/inputs/input7225"),
		// 			Properties: &armstreamanalytics.ReferenceInputProperties{
		// 				Type: to.Ptr("Reference"),
		// 				Etag: to.Ptr("a4ceb697-1c8f-40c8-b951-fb5ee4757437"),
		// 				Serialization: &armstreamanalytics.CSVSerialization{
		// 					Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
		// 					Properties: &armstreamanalytics.CSVSerializationProperties{
		// 						Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
		// 						FieldDelimiter: to.Ptr("|"),
		// 					},
		// 				},
		// 				Datasource: &armstreamanalytics.BlobReferenceInputDataSource{
		// 					Type: to.Ptr("Microsoft.Storage/Blob"),
		// 					Properties: &armstreamanalytics.BlobReferenceInputDataSourceProperties{
		// 						Container: to.Ptr("differentContainer"),
		// 						DateFormat: to.Ptr("yyyy/MM/dd"),
		// 						PathPattern: to.Ptr("{date}/{time}"),
		// 						StorageAccounts: []*armstreamanalytics.StorageAccount{
		// 							{
		// 								AccountName: to.Ptr("someAccountName"),
		// 						}},
		// 						TimeFormat: to.Ptr("HH"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("input8899"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg8440/providers/Microsoft.StreamAnalytics/streamingjobs/sj9597/inputs/input8899"),
		// 			Properties: &armstreamanalytics.StreamInputProperties{
		// 				Type: to.Ptr("Stream"),
		// 				Etag: to.Ptr("3b35d57c-02f4-4b41-8e1d-af02a86c2fa1"),
		// 				Serialization: &armstreamanalytics.CSVSerialization{
		// 					Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
		// 					Properties: &armstreamanalytics.CSVSerializationProperties{
		// 						Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
		// 						FieldDelimiter: to.Ptr("|"),
		// 					},
		// 				},
		// 				Datasource: &armstreamanalytics.BlobStreamInputDataSource{
		// 					Type: to.Ptr("Microsoft.Storage/Blob"),
		// 					Properties: &armstreamanalytics.BlobStreamInputDataSourceProperties{
		// 						Container: to.Ptr("state"),
		// 						DateFormat: to.Ptr("yyyy/MM/dd"),
		// 						PathPattern: to.Ptr("{date}/{time}"),
		// 						StorageAccounts: []*armstreamanalytics.StorageAccount{
		// 							{
		// 								AccountName: to.Ptr("someAccountName"),
		// 						}},
		// 						TimeFormat: to.Ptr("HH"),
		// 						SourcePartitionCount: to.Ptr[int32](32),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("input7425"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg8440/providers/Microsoft.StreamAnalytics/streamingjobs/sj9597/inputs/input7425"),
		// 			Properties: &armstreamanalytics.StreamInputProperties{
		// 				Type: to.Ptr("Stream"),
		// 				Etag: to.Ptr("7548f259-81b5-4ea7-b896-25c6717d98ba"),
		// 				Serialization: &armstreamanalytics.AvroSerialization{
		// 					Type: to.Ptr(armstreamanalytics.EventSerializationTypeAvro),
		// 					Properties: map[string]any{
		// 					},
		// 				},
		// 				Datasource: &armstreamanalytics.EventHubStreamInputDataSource{
		// 					Type: to.Ptr("Microsoft.ServiceBus/EventHub"),
		// 					Properties: &armstreamanalytics.EventHubStreamInputDataSourceProperties{
		// 						ServiceBusNamespace: to.Ptr("sdktest"),
		// 						SharedAccessPolicyName: to.Ptr("RootManageSharedAccessKey"),
		// 						EventHubName: to.Ptr("sdkeventhub"),
		// 						ConsumerGroupName: to.Ptr("differentConsumerGroupName"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("input7970"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg8440/providers/Microsoft.StreamAnalytics/streamingjobs/sj9597/inputs/input7970"),
		// 			Properties: &armstreamanalytics.StreamInputProperties{
		// 				Type: to.Ptr("Stream"),
		// 				Etag: to.Ptr("e2d847e0-c95b-48ef-9e14-1afc1f2270cb"),
		// 				Serialization: &armstreamanalytics.CSVSerialization{
		// 					Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
		// 					Properties: &armstreamanalytics.CSVSerializationProperties{
		// 						Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
		// 						FieldDelimiter: to.Ptr("|"),
		// 					},
		// 				},
		// 				Datasource: &armstreamanalytics.IoTHubStreamInputDataSource{
		// 					Type: to.Ptr("Microsoft.Devices/IotHubs"),
		// 					Properties: &armstreamanalytics.IoTHubStreamInputDataSourceProperties{
		// 						ConsumerGroupName: to.Ptr("sdkconsumergroup"),
		// 						Endpoint: to.Ptr("messages/operationsMonitoringEvents"),
		// 						IotHubNamespace: to.Ptr("iothub"),
		// 						SharedAccessPolicyName: to.Ptr("owner"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
