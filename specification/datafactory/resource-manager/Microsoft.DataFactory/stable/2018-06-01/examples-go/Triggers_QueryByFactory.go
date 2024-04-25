package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_QueryByFactory.json
func ExampleTriggersClient_QueryByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTriggersClient().QueryByFactory(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.TriggerFilterParameters{
		ParentTriggerName: to.Ptr("exampleTrigger"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TriggerQueryResponse = armdatafactory.TriggerQueryResponse{
	// 	Value: []*armdatafactory.TriggerResource{
	// 		{
	// 			Name: to.Ptr("exampleRerunTrigger"),
	// 			Type: to.Ptr("Microsoft.DataFactory/factories/triggers"),
	// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/triggers/exampleRerunTrigger"),
	// 			Properties: &armdatafactory.RerunTumblingWindowTrigger{
	// 				Type: to.Ptr("RerunTumblingWindowTrigger"),
	// 				Description: to.Ptr("Example description"),
	// 				TypeProperties: &armdatafactory.RerunTumblingWindowTriggerTypeProperties{
	// 					ParentTrigger: "exampleTrigger",
	// 					RequestedEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:14.905Z"); return t}()),
	// 					RequestedStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:14.905Z"); return t}()),
	// 					RerunConcurrency: to.Ptr[int32](4),
	// 				},
	// 			},
	// 	}},
	// }
}
