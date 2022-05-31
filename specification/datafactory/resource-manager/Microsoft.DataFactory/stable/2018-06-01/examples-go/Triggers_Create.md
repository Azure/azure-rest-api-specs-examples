Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv1.1.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Create.json
func ExampleTriggersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewTriggersClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"exampleResourceGroup",
		"exampleFactoryName",
		"exampleTrigger",
		armdatafactory.TriggerResource{
			Properties: &armdatafactory.ScheduleTrigger{
				Type: to.Ptr("ScheduleTrigger"),
				Pipelines: []*armdatafactory.TriggerPipelineReference{
					{
						Parameters: map[string]interface{}{
							"OutputBlobNameList": []interface{}{
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
		},
		&armdatafactory.TriggersClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
