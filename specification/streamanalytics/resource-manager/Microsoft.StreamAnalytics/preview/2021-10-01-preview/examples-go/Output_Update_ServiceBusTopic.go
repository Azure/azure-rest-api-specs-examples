package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_Update_ServiceBusTopic.json
func ExampleOutputsClient_Update_updateAServiceBusTopicOutputWithCsvSerialization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOutputsClient().Update(ctx, "sjrg6450", "sj7094", "output7886", armstreamanalytics.Output{
		Properties: &armstreamanalytics.OutputProperties{
			Datasource: &armstreamanalytics.ServiceBusTopicOutputDataSource{
				Type: to.Ptr("Microsoft.ServiceBus/Topic"),
				Properties: &armstreamanalytics.ServiceBusTopicOutputDataSourceProperties{
					TopicName: to.Ptr("differentTopicName"),
				},
			},
			Serialization: &armstreamanalytics.CSVSerialization{
				Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
				Properties: &armstreamanalytics.CSVSerializationProperties{
					Encoding:       to.Ptr(armstreamanalytics.EncodingUTF8),
					FieldDelimiter: to.Ptr("|"),
				},
			},
		},
	}, &armstreamanalytics.OutputsClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Output = armstreamanalytics.Output{
	// 	Name: to.Ptr("output7886"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg6450/providers/Microsoft.StreamAnalytics/streamingjobs/sj7094/outputs/output7886"),
	// 	Properties: &armstreamanalytics.OutputProperties{
	// 		Datasource: &armstreamanalytics.ServiceBusTopicOutputDataSource{
	// 			Type: to.Ptr("Microsoft.ServiceBus/Topic"),
	// 			Properties: &armstreamanalytics.ServiceBusTopicOutputDataSourceProperties{
	// 				ServiceBusNamespace: to.Ptr("sdktest"),
	// 				SharedAccessPolicyName: to.Ptr("RootManageSharedAccessKey"),
	// 				PropertyColumns: []*string{
	// 					to.Ptr("column1"),
	// 					to.Ptr("column2")},
	// 					TopicName: to.Ptr("differentTopicName"),
	// 				},
	// 			},
	// 			Serialization: &armstreamanalytics.CSVSerialization{
	// 				Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
	// 				Properties: &armstreamanalytics.CSVSerializationProperties{
	// 					Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
	// 					FieldDelimiter: to.Ptr("|"),
	// 				},
	// 			},
	// 		},
	// 	}
}
