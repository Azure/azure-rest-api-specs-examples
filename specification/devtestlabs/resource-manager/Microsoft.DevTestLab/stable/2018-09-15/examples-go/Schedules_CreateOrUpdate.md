Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv0.4.0/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Schedules_CreateOrUpdate.json
func ExampleSchedulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewSchedulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		armdevtestlabs.Schedule{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"tagName1": to.Ptr("tagValue1"),
			},
			Properties: &armdevtestlabs.ScheduleProperties{
				DailyRecurrence: &armdevtestlabs.DayDetails{
					Time: to.Ptr("<time>"),
				},
				HourlyRecurrence: &armdevtestlabs.HourDetails{
					Minute: to.Ptr[int32](30),
				},
				NotificationSettings: &armdevtestlabs.NotificationSettings{
					EmailRecipient:     to.Ptr("<email-recipient>"),
					NotificationLocale: to.Ptr("<notification-locale>"),
					Status:             to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
					TimeInMinutes:      to.Ptr[int32](15),
					WebhookURL:         to.Ptr("<webhook-url>"),
				},
				Status:           to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
				TargetResourceID: to.Ptr("<target-resource-id>"),
				TaskType:         to.Ptr("<task-type>"),
				TimeZoneID:       to.Ptr("<time-zone-id>"),
				WeeklyRecurrence: &armdevtestlabs.WeekDetails{
					Time: to.Ptr("<time>"),
					Weekdays: []*string{
						to.Ptr("Monday"),
						to.Ptr("Wednesday"),
						to.Ptr("Friday")},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
