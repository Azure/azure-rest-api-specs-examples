package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/ScheduledActions_ListBySubscription_PagedSuccess.json
func ExampleScheduledActionsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScheduledActionsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armbulkactions.ScheduledActionsClientListBySubscriptionResponse{
		// 	ScheduledActionListResult: armbulkactions.ScheduledActionListResult{
		// 		Value: []*armbulkactions.ScheduledAction{
		// 			{
		// 				Properties: &armbulkactions.ScheduledActionProperties{
		// 					ResourceType: to.Ptr(armbulkactions.ResourceTypeVirtualMachine),
		// 					ActionType: to.Ptr(armbulkactions.ScheduledActionTypeStart),
		// 					StartTime: to.Ptr(time.Date(2026, time.September, 15, 7, 0, 0, 0, time.FixedZone("", -25200))),
		// 					Schedule: &armbulkactions.ScheduledActionsSchedule{
		// 						ScheduledTime: to.Ptr(time.Date(0, time.January, 1, 7, 0, 0, 0, time.UTC)),
		// 						TimeZone: to.Ptr("America/Los_Angeles"),
		// 						RequestedWeekDays: []*armbulkactions.WeekDay{
		// 							to.Ptr(armbulkactions.WeekDayMonday),
		// 							to.Ptr(armbulkactions.WeekDayTuesday),
		// 							to.Ptr(armbulkactions.WeekDayWednesday),
		// 							to.Ptr(armbulkactions.WeekDayThursday),
		// 							to.Ptr(armbulkactions.WeekDayFriday),
		// 						},
		// 					},
		// 					NotificationSettings: []*armbulkactions.NotificationProperties{
		// 					},
		// 					Disabled: to.Ptr(false),
		// 					ProvisioningState: to.Ptr(armbulkactions.ScheduledActionsProvisioningStateSucceeded),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/web-rg/providers/Microsoft.Compute/scheduledActions/weekday-start"),
		// 				Name: to.Ptr("weekday-start"),
		// 				Type: to.Ptr("Microsoft.Compute/scheduledActions"),
		// 			},
		// 			{
		// 				Properties: &armbulkactions.ScheduledActionProperties{
		// 					ResourceType: to.Ptr(armbulkactions.ResourceTypeVirtualMachine),
		// 					ActionType: to.Ptr(armbulkactions.ScheduledActionTypeDeallocate),
		// 					StartTime: to.Ptr(time.Date(2026, time.September, 15, 19, 0, 0, 0, time.FixedZone("", -18000))),
		// 					Schedule: &armbulkactions.ScheduledActionsSchedule{
		// 						ScheduledTime: to.Ptr(time.Date(0, time.January, 1, 19, 0, 0, 0, time.UTC)),
		// 						TimeZone: to.Ptr("America/Chicago"),
		// 						RequestedWeekDays: []*armbulkactions.WeekDay{
		// 							to.Ptr(armbulkactions.WeekDayMonday),
		// 							to.Ptr(armbulkactions.WeekDayTuesday),
		// 							to.Ptr(armbulkactions.WeekDayWednesday),
		// 							to.Ptr(armbulkactions.WeekDayThursday),
		// 							to.Ptr(armbulkactions.WeekDayFriday),
		// 						},
		// 					},
		// 					NotificationSettings: []*armbulkactions.NotificationProperties{
		// 					},
		// 					Disabled: to.Ptr(false),
		// 					ProvisioningState: to.Ptr(armbulkactions.ScheduledActionsProvisioningStateSucceeded),
		// 				},
		// 				Location: to.Ptr("centralus"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/batch-rg/providers/Microsoft.Compute/scheduledActions/weekday-deallocate"),
		// 				Name: to.Ptr("weekday-deallocate"),
		// 				Type: to.Ptr("Microsoft.Compute/scheduledActions"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/scheduledActions?api-version=2026-10-06-preview&$skiptoken=page2"),
		// 	},
		// }
	}
}
