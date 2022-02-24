Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flabservices%2Farmlabservices%2Fv0.2.1/sdk/resourcemanager/labservices/armlabservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armlabservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Schedules/patchSchedule.json
func ExampleSchedulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewSchedulesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<schedule-name>",
		armlabservices.ScheduleUpdate{
			Properties: &armlabservices.ScheduleUpdateProperties{
				RecurrencePattern: &armlabservices.RecurrencePattern{
					ExpirationDate: to.TimePtr(func() time.Time { t, _ := time.Parse("2006-01-02", "2020-08-14"); return t }()),
					Frequency:      armlabservices.RecurrenceFrequencyDaily.ToPtr(),
					Interval:       to.Int32Ptr(2),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SchedulesClientUpdateResult)
}
```
