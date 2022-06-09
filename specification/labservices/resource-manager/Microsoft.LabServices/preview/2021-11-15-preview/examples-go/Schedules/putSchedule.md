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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Schedules/putSchedule.json
func ExampleSchedulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlabservices.NewSchedulesClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"testrg123",
		"testlab",
		"schedule1",
		armlabservices.Schedule{
			Properties: &armlabservices.ScheduleProperties{
				Notes: to.Ptr("Schedule 1 for students"),
				RecurrencePattern: &armlabservices.RecurrencePattern{
					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2020-08-14"); return t }()),
					Frequency:      to.Ptr(armlabservices.RecurrenceFrequencyDaily),
					Interval:       to.Ptr[int32](2),
				},
				StartAt:    to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-26T12:00:00Z"); return t }()),
				StopAt:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-26T18:00:00Z"); return t }()),
				TimeZoneID: to.Ptr("America/Los_Angeles"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flabservices%2Farmlabservices%2Fv0.5.0/sdk/resourcemanager/labservices/armlabservices/README.md) on how to add the SDK to your project and authenticate.
