package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-08-06-preview/ScheduledActions_ListByResourceGroup_MaximumSet_Gen.json
func ExampleScheduledActionsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScheduledActionsClient().NewListByResourceGroupPager("rgcompute", nil)
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
		// page = armbulkactions.ScheduledActionsClientListByResourceGroupResponse{
		// 	ScheduledActionListResult: armbulkactions.ScheduledActionListResult{
		// 		Value: []*armbulkactions.ScheduledAction{
		// 			{
		// 				Properties: &armbulkactions.ScheduledActionProperties{
		// 					ResourceType: to.Ptr(armbulkactions.ResourceTypeVirtualMachine),
		// 					ActionType: to.Ptr(armbulkactions.ScheduledActionTypeStart),
		// 					StartTime: to.Ptr(time.Date(2025, time.April, 17, 0, 23, 55, 281000000, time.UTC)),
		// 					EndTime: to.Ptr(time.Date(2025, time.April, 17, 0, 23, 55, 286000000, time.UTC)),
		// 					Schedule: &armbulkactions.ScheduledActionsSchedule{
		// 						ScheduledTime: to.Ptr(time.Date(0, time.January, 1, 19, 0, 0, 0, time.UTC)),
		// 						TimeZone: to.Ptr("America/Los_Angeles"),
		// 						RequestedWeekDays: []*armbulkactions.WeekDay{
		// 							to.Ptr(armbulkactions.WeekDayMonday),
		// 						},
		// 						RequestedMonths: []*armbulkactions.Month{
		// 							to.Ptr(armbulkactions.MonthJanuary),
		// 						},
		// 						RequestedDaysOfTheMonth: []*int32{
		// 							to.Ptr[int32](15),
		// 						},
		// 						ExecutionParameters: &armbulkactions.ScheduledActionsExecutionParameters{
		// 							OptimizationPreference: to.Ptr(armbulkactions.OptimizationPreferenceCost),
		// 							RetryPolicy: &armbulkactions.ScheduledActionsRetryPolicy{
		// 								RetryCount: to.Ptr[int32](17),
		// 								RetryWindowInMinutes: to.Ptr[int32](29),
		// 							},
		// 						},
		// 						DeadlineType: to.Ptr(armbulkactions.ScheduledActionsDeadlineTypeUnknown),
		// 					},
		// 					NotificationSettings: []*armbulkactions.NotificationProperties{
		// 						{
		// 							Destination: to.Ptr("admin@contoso.com"),
		// 							Type: to.Ptr(armbulkactions.NotificationTypeEmail),
		// 							Language: to.Ptr(armbulkactions.LanguageEnUs),
		// 							Disabled: to.Ptr(true),
		// 						},
		// 					},
		// 					Disabled: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armbulkactions.ScheduledActionsProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key2102": to.Ptr("myTagValue"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/83C27AB3-A7B9-498B-B165-D9440661474F/resourceGroups/myRg/providers/Microsoft.Compute/scheduledActions/myScheduledAction"),
		// 				Name: to.Ptr("myScheduledAction"),
		// 				Type: to.Ptr("Microsoft.Compute/scheduledActions"),
		// 				SystemData: &armbulkactions.SystemData{
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2025, time.April, 17, 0, 23, 55, 288000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.April, 17, 0, 23, 55, 288000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
