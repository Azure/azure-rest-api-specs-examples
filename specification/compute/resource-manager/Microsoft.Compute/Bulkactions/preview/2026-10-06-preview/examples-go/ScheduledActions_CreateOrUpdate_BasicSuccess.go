package armbulkactions_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/ScheduledActions_CreateOrUpdate_BasicSuccess.json
func ExampleScheduledActionsClient_BeginCreateOrUpdate_oneCreateANewRecurringScheduledAction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScheduledActionsClient().BeginCreateOrUpdate(ctx, "example-rg", "weekday-start", armbulkactions.ScheduledAction{
		Properties: &armbulkactions.ScheduledActionProperties{
			ResourceType: to.Ptr(armbulkactions.ResourceTypeVirtualMachine),
			ActionType:   to.Ptr(armbulkactions.ScheduledActionTypeStart),
			StartTime:    to.Ptr(time.Date(2026, time.September, 15, 7, 0, 0, 0, time.FixedZone("", -25200))),
			Schedule: &armbulkactions.ScheduledActionsSchedule{
				ScheduledTime: to.Ptr(time.Date(0, time.January, 1, 7, 0, 0, 0, time.UTC)),
				TimeZone:      to.Ptr("America/Los_Angeles"),
				RequestedWeekDays: []*armbulkactions.WeekDay{
					to.Ptr(armbulkactions.WeekDayMonday),
					to.Ptr(armbulkactions.WeekDayTuesday),
					to.Ptr(armbulkactions.WeekDayWednesday),
					to.Ptr(armbulkactions.WeekDayThursday),
					to.Ptr(armbulkactions.WeekDayFriday),
				},
			},
			NotificationSettings: []*armbulkactions.NotificationProperties{},
		},
		Location: to.Ptr("eastus"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.ScheduledActionsClientCreateOrUpdateResponse{
	// 	ScheduledAction: armbulkactions.ScheduledAction{
	// 		Properties: &armbulkactions.ScheduledActionProperties{
	// 			ResourceType: to.Ptr(armbulkactions.ResourceTypeVirtualMachine),
	// 			ActionType: to.Ptr(armbulkactions.ScheduledActionTypeStart),
	// 			StartTime: to.Ptr(time.Date(2026, time.September, 15, 7, 0, 0, 0, time.FixedZone("", -25200))),
	// 			Schedule: &armbulkactions.ScheduledActionsSchedule{
	// 				ScheduledTime: to.Ptr(time.Date(0, time.January, 1, 7, 0, 0, 0, time.UTC)),
	// 				TimeZone: to.Ptr("America/Los_Angeles"),
	// 				RequestedWeekDays: []*armbulkactions.WeekDay{
	// 					to.Ptr(armbulkactions.WeekDayMonday),
	// 					to.Ptr(armbulkactions.WeekDayTuesday),
	// 					to.Ptr(armbulkactions.WeekDayWednesday),
	// 					to.Ptr(armbulkactions.WeekDayThursday),
	// 					to.Ptr(armbulkactions.WeekDayFriday),
	// 				},
	// 				RequestedMonths: []*armbulkactions.Month{
	// 					to.Ptr(armbulkactions.MonthAll),
	// 				},
	// 				RequestedDaysOfTheMonth: []*int32{
	// 				},
	// 			},
	// 			NotificationSettings: []*armbulkactions.NotificationProperties{
	// 			},
	// 			ProvisioningState: to.Ptr(armbulkactions.ScheduledActionsProvisioningStateSucceeded),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/scheduledActions/weekday-start"),
	// 		Name: to.Ptr("weekday-start"),
	// 		Type: to.Ptr("Microsoft.Compute/scheduledActions"),
	// 	},
	// }
}
