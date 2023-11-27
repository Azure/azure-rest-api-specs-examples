package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/69ece3818b8b0929b43a07c3fe25716427734882/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_ListByFactory.json
func ExampleTriggersClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTriggersClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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
		// page.TriggerListResponse = armdatafactory.TriggerListResponse{
		// 	Value: []*armdatafactory.TriggerResource{
		// 		{
		// 			Name: to.Ptr("exampleTrigger"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/triggers"),
		// 			Etag: to.Ptr("0a008ed4-0000-0000-0000-5b245c740000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/triggers/exampleTrigger"),
		// 			Properties: &armdatafactory.ScheduleTrigger{
		// 				Type: to.Ptr("ScheduleTrigger"),
		// 				Description: to.Ptr("Example description"),
		// 				RuntimeState: to.Ptr(armdatafactory.TriggerRuntimeStateStarted),
		// 				Pipelines: []*armdatafactory.TriggerPipelineReference{
		// 					{
		// 						Parameters: map[string]any{
		// 							"OutputBlobNameList": []any{
		// 								"exampleoutput.csv",
		// 							},
		// 						},
		// 						PipelineReference: &armdatafactory.PipelineReference{
		// 							Type: to.Ptr(armdatafactory.PipelineReferenceTypePipelineReference),
		// 							ReferenceName: to.Ptr("examplePipeline"),
		// 						},
		// 				}},
		// 				TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
		// 					Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
		// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:14.905Z"); return t}()),
		// 						Frequency: to.Ptr(armdatafactory.RecurrenceFrequencyMinute),
		// 						Interval: to.Ptr[int32](4),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:14.905Z"); return t}()),
		// 						TimeZone: to.Ptr("UTC"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
