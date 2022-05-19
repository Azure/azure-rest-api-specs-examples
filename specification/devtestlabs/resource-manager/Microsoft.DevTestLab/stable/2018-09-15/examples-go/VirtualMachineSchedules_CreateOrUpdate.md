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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_CreateOrUpdate.json
func ExampleVirtualMachineSchedulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewVirtualMachineSchedulesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourceGroupName",
		"{labName}",
		"{vmName}",
		"LabVmsShutdown",
		armdevtestlabs.Schedule{
			Location: to.Ptr("{location}"),
			Tags: map[string]*string{
				"tagName1": to.Ptr("tagValue1"),
			},
			Properties: &armdevtestlabs.ScheduleProperties{
				DailyRecurrence: &armdevtestlabs.DayDetails{
					Time: to.Ptr("1900"),
				},
				HourlyRecurrence: &armdevtestlabs.HourDetails{
					Minute: to.Ptr[int32](30),
				},
				NotificationSettings: &armdevtestlabs.NotificationSettings{
					EmailRecipient:     to.Ptr("{email}"),
					NotificationLocale: to.Ptr("EN"),
					Status:             to.Ptr(armdevtestlabs.EnableStatusEnabled),
					TimeInMinutes:      to.Ptr[int32](30),
					WebhookURL:         to.Ptr("{webhookUrl}"),
				},
				Status:           to.Ptr(armdevtestlabs.EnableStatusEnabled),
				TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}"),
				TaskType:         to.Ptr("LabVmsShutdownTask"),
				TimeZoneID:       to.Ptr("Pacific Standard Time"),
				WeeklyRecurrence: &armdevtestlabs.WeekDetails{
					Time: to.Ptr("1700"),
					Weekdays: []*string{
						to.Ptr("Friday"),
						to.Ptr("Saturday"),
						to.Ptr("Sunday")},
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
