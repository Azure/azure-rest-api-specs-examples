package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Schedules_CreateOrUpdate.json
func ExampleSchedulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSchedulesClient().CreateOrUpdate(ctx, "resourceGroupName", "{labName}", "{scheduleName}", armdevtestlabs.Schedule{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Schedule = armdevtestlabs.Schedule{
	// 	Name: to.Ptr("{scheduleName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/schedules"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/{scheduleName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.ScheduleProperties{
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-02T01:40:48.173Z"); return t}()),
	// 		DailyRecurrence: &armdevtestlabs.DayDetails{
	// 			Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
	// 		},
	// 		HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 			Minute: to.Ptr[int32](30),
	// 		},
	// 		NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 			EmailRecipient: to.Ptr("{email}"),
	// 			NotificationLocale: to.Ptr("EN"),
	// 			Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 			TimeInMinutes: to.Ptr[int32](15),
	// 			WebhookURL: to.Ptr("{webhookUrl}"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 		TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
	// 		TaskType: to.Ptr("{myLabVmTaskType}"),
	// 		TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 		UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 		WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 			Time: to.Ptr("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
	// 			Weekdays: []*string{
	// 				to.Ptr("Monday"),
	// 				to.Ptr("Wednesday"),
	// 				to.Ptr("Friday")},
	// 			},
	// 		},
	// 	}
}
