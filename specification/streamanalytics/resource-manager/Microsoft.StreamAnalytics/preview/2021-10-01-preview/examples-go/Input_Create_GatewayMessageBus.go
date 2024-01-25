package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_Create_GatewayMessageBus.json
func ExampleInputsClient_CreateOrReplace_createAGatewayMessageBusInput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInputsClient().CreateOrReplace(ctx, "sjrg3467", "sj9742", "input7970", armstreamanalytics.Input{
		Properties: &armstreamanalytics.StreamInputProperties{
			Type: to.Ptr("Stream"),
			Datasource: &armstreamanalytics.GatewayMessageBusStreamInputDataSource{
				Type: to.Ptr("GatewayMessageBus"),
				Properties: &armstreamanalytics.GatewayMessageBusStreamInputDataSourceProperties{
					Topic: to.Ptr("EdgeTopic1"),
				},
			},
		},
	}, &armstreamanalytics.InputsClientCreateOrReplaceOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Input = armstreamanalytics.Input{
	// 	Name: to.Ptr("input7970"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/inputs"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg3467/providers/Microsoft.StreamAnalytics/streamingjobs/sj9742/inputs/input7970"),
	// 	Properties: &armstreamanalytics.StreamInputProperties{
	// 		Type: to.Ptr("Stream"),
	// 		Datasource: &armstreamanalytics.GatewayMessageBusStreamInputDataSource{
	// 			Type: to.Ptr("GatewayMessageBus"),
	// 			Properties: &armstreamanalytics.GatewayMessageBusStreamInputDataSourceProperties{
	// 				Topic: to.Ptr("EdgeTopic1"),
	// 			},
	// 		},
	// 	},
	// }
}
