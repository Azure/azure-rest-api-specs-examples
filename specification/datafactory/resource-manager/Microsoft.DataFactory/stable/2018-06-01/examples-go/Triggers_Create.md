Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv0.3.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Create.json
func ExampleTriggersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewTriggersClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<trigger-name>",
		armdatafactory.TriggerResource{
			Properties: &armdatafactory.ScheduleTrigger{
				Type: to.StringPtr("<type>"),
				Pipelines: []*armdatafactory.TriggerPipelineReference{
					{
						Parameters: map[string]interface{}{
							"OutputBlobNameList": []interface{}{
								"exampleoutput.csv",
							},
						},
						PipelineReference: &armdatafactory.PipelineReference{
							Type:          armdatafactory.PipelineReferenceType("PipelineReference").ToPtr(),
							ReferenceName: to.StringPtr("<reference-name>"),
						},
					}},
				TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
					Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
						EndTime:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:13.8441801Z"); return t }()),
						Frequency: armdatafactory.RecurrenceFrequency("Minute").ToPtr(),
						Interval:  to.Int32Ptr(4),
						StartTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:13.8441801Z"); return t }()),
						TimeZone:  to.StringPtr("<time-zone>"),
					},
				},
			},
		},
		&armdatafactory.TriggersClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TriggersClientCreateOrUpdateResult)
}
```
