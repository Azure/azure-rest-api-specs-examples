Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flabservices%2Farmlabservices%2Fv0.1.0/sdk/resourcemanager/labservices/armlabservices/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Schedules/putSchedule.json
func ExampleSchedulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewSchedulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<schedule-name>",
		armlabservices.Schedule{
			Properties: &armlabservices.ScheduleProperties{
				ScheduleUpdateProperties: armlabservices.ScheduleUpdateProperties{
					Notes: to.StringPtr("<notes>"),
					RecurrencePattern: &armlabservices.RecurrencePattern{
						ExpirationDate: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-14"); return t }()),
						Frequency:      armlabservices.RecurrenceFrequencyDaily.ToPtr(),
						Interval:       to.Int32Ptr(2),
					},
					StartAt:    to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-26T12:00:00Z"); return t }()),
					StopAt:     to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-26T18:00:00Z"); return t }()),
					TimeZoneID: to.StringPtr("<time-zone-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Schedule.ID: %s\n", *res.ID)
}
```
