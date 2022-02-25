Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv0.2.1/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_CreateOrUpdate.json
func ExampleServiceFabricSchedulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewServiceFabricSchedulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<user-name>",
		"<service-fabric-name>",
		"<name>",
		armdevtestlabs.Schedule{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tagName1": to.StringPtr("tagValue1"),
			},
			Properties: &armdevtestlabs.ScheduleProperties{
				DailyRecurrence: &armdevtestlabs.DayDetails{
					Time: to.StringPtr("<time>"),
				},
				HourlyRecurrence: &armdevtestlabs.HourDetails{
					Minute: to.Int32Ptr(0),
				},
				NotificationSettings: &armdevtestlabs.NotificationSettings{
					EmailRecipient:     to.StringPtr("<email-recipient>"),
					NotificationLocale: to.StringPtr("<notification-locale>"),
					Status:             armdevtestlabs.EnableStatus("{Enabled|Disabled}").ToPtr(),
					TimeInMinutes:      to.Int32Ptr(15),
					WebhookURL:         to.StringPtr("<webhook-url>"),
				},
				Status:           armdevtestlabs.EnableStatus("{Enabled|Disabled}").ToPtr(),
				TargetResourceID: to.StringPtr("<target-resource-id>"),
				TaskType:         to.StringPtr("<task-type>"),
				TimeZoneID:       to.StringPtr("<time-zone-id>"),
				WeeklyRecurrence: &armdevtestlabs.WeekDetails{
					Time: to.StringPtr("<time>"),
					Weekdays: []*string{
						to.StringPtr("Monday"),
						to.StringPtr("Tuesday"),
						to.StringPtr("Wednesday"),
						to.StringPtr("Thursday"),
						to.StringPtr("Friday"),
						to.StringPtr("Saturday"),
						to.StringPtr("Sunday")},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServiceFabricSchedulesClientCreateOrUpdateResult)
}
```
