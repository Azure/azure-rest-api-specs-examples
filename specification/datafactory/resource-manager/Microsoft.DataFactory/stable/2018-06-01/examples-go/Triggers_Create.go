package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Create.json
func ExampleTriggersClient_CreateOrUpdate_triggersCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTriggersClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleTrigger", armdatafactory.TriggerResource{
		Properties: &armdatafactory.ScheduleTrigger{
			Type: to.Ptr("ScheduleTrigger"),
			Pipelines: []*armdatafactory.TriggerPipelineReference{
				{
					Parameters: map[string]any{
						"OutputBlobNameList": []any{
							"exampleoutput.csv",
						},
					},
					PipelineReference: &armdatafactory.PipelineReference{
						Type:          to.Ptr(armdatafactory.PipelineReferenceTypePipelineReference),
						ReferenceName: to.Ptr("examplePipeline"),
					},
				}},
			TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
				Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
					EndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:13.8441801Z"); return t }()),
					Frequency: to.Ptr(armdatafactory.RecurrenceFrequencyMinute),
					Interval:  to.Ptr[int32](4),
					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:13.8441801Z"); return t }()),
					TimeZone:  to.Ptr("UTC"),
				},
			},
		},
	}, &armdatafactory.TriggersClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TriggerResource = armdatafactory.TriggerResource{
	// 	Name: to.Ptr("exampleTrigger"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/triggers"),
	// 	Etag: to.Ptr("0a008ad4-0000-0000-0000-5b245c6e0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/triggers/exampleTrigger"),
	// 	Properties: &armdatafactory.ScheduleTrigger{
	// 		Type: to.Ptr("ScheduleTrigger"),
	// 		RuntimeState: to.Ptr(armdatafactory.TriggerRuntimeStateStopped),
	// 		Pipelines: []*armdatafactory.TriggerPipelineReference{
	// 			{
	// 				Parameters: map[string]any{
	// 					"OutputBlobNameList": []any{
	// 						"exampleoutput.csv",
	// 					},
	// 				},
	// 				PipelineReference: &armdatafactory.PipelineReference{
	// 					Type: to.Ptr(armdatafactory.PipelineReferenceTypePipelineReference),
	// 					ReferenceName: to.Ptr("examplePipeline"),
	// 				},
	// 		}},
	// 		TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
	// 			Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:13.8441801Z"); return t}()),
	// 				Frequency: to.Ptr(armdatafactory.RecurrenceFrequencyMinute),
	// 				Interval: to.Ptr[int32](4),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:13.8441801Z"); return t}()),
	// 				TimeZone: to.Ptr("UTC"),
	// 			},
	// 		},
	// 	},
	// }
}
