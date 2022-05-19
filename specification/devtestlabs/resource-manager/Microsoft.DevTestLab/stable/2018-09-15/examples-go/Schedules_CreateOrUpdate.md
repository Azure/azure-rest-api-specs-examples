Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv1.0.0/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

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
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewSchedulesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourceGroupName",
		"{labName}",
		"{scheduleName}",
		armdevtestlabs.Schedule{
			Location: to.Ptr("{location}"),
			Tags: map[string]*string{
				"tagName1": to.Ptr("tagValue1"),
			},
			Properties: &armdevtestlabs.ScheduleProperties{
				DailyRecurrence: &armdevtestlabs.DayDetails{
					Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
				},
				HourlyRecurrence: &armdevtestlabs.HourDetails{
					Minute: to.Ptr[int32](30),
				},
				NotificationSettings: &armdevtestlabs.NotificationSettings{
					EmailRecipient:     to.Ptr("{email}"),
					NotificationLocale: to.Ptr("EN"),
					Status:             to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
					TimeInMinutes:      to.Ptr[int32](15),
					WebhookURL:         to.Ptr("{webhookUrl}"),
				},
				Status:           to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
				TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
				TaskType:         to.Ptr("{myLabVmTaskType}"),
				TimeZoneID:       to.Ptr("Pacific Standard Time"),
				WeeklyRecurrence: &armdevtestlabs.WeekDetails{
					Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
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
	}
	// TODO: use response item
	_ = res
}
```
