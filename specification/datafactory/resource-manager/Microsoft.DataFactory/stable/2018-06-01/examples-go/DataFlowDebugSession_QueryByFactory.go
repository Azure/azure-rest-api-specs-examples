package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_QueryByFactory.json
func ExampleDataFlowDebugSessionClient_NewQueryByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataFlowDebugSessionClient().NewQueryByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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
		// page.QueryDataFlowDebugSessionsResponse = armdatafactory.QueryDataFlowDebugSessionsResponse{
		// 	Value: []*armdatafactory.DataFlowDebugSessionInfo{
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"dataflowName": "DebugSession-0a7e0d6e-f2b7-48cc-8cd8-618326f5662f",
		// 				"userObjectId": "0a7e0d6e-f2b7-48cc-8cd8-618326f5662f",
		// 			},
		// 			ComputeType: to.Ptr("General"),
		// 			CoreCount: to.Ptr[int32](48),
		// 			LastActivityTime: to.Ptr("2019-09-05T18:28:00.9459674+00:00"),
		// 			SessionID: to.Ptr("229c688c-944c-44ac-b31a-82d50f347154"),
		// 			StartTime: to.Ptr("2019-09-05T18:23:20.3257799+00:00"),
		// 			TimeToLiveInMinutes: to.Ptr[int32](60),
		// 	}},
		// }
	}
}
