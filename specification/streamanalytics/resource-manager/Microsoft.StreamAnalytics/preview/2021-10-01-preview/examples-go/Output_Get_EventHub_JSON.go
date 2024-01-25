package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_Get_EventHub_JSON.json
func ExampleOutputsClient_Get_getAnEventHubOutputWithJsonSerialization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOutputsClient().Get(ctx, "sjrg6912", "sj3310", "output5195", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Output = armstreamanalytics.Output{
	// 	Name: to.Ptr("output5195"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/outputs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg6912/providers/Microsoft.StreamAnalytics/streamingjobs/sj3310/outputs/output5195"),
	// 	Properties: &armstreamanalytics.OutputProperties{
	// 		Datasource: &armstreamanalytics.EventHubOutputDataSource{
	// 			Type: to.Ptr("Microsoft.ServiceBus/EventHub"),
	// 			Properties: &armstreamanalytics.EventHubOutputDataSourceProperties{
	// 				ServiceBusNamespace: to.Ptr("sdktest"),
	// 				SharedAccessPolicyName: to.Ptr("RootManageSharedAccessKey"),
	// 				EventHubName: to.Ptr("sdkeventhub"),
	// 				PartitionKey: to.Ptr("partitionKey"),
	// 			},
	// 		},
	// 		LastOutputEventTimestamps: []*armstreamanalytics.LastOutputEventTimestamp{
	// 			{
	// 				LastOutputEventTime: to.Ptr("2009-06-15T13:45:30"),
	// 				LastUpdateTime: to.Ptr("2009-06-15T13:45:30"),
	// 			},
	// 			{
	// 				LastOutputEventTime: to.Ptr("2009-07-15T13:45:30"),
	// 				LastUpdateTime: to.Ptr("2009-06-15T13:45:30"),
	// 		}},
	// 		Serialization: &armstreamanalytics.JSONSerialization{
	// 			Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
	// 			Properties: &armstreamanalytics.JSONSerializationProperties{
	// 				Format: to.Ptr(armstreamanalytics.JSONOutputSerializationFormatArray),
	// 				Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
	// 			},
	// 		},
	// 	},
	// }
}
