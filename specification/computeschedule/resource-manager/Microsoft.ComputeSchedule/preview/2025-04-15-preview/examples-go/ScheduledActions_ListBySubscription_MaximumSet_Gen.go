package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: 2025-04-15-preview/ScheduledActions_ListBySubscription_MaximumSet_Gen.json
func ExampleScheduledActionsClient_NewListBySubscriptionPager_scheduledActionsListBySubscriptionMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
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
		// page = armcomputeschedule.ScheduledActionsClientListBySubscriptionResponse{
		// 	ScheduledActionListResult: armcomputeschedule.ScheduledActionListResult{
		// 		Value: []*armcomputeschedule.ScheduledAction{
		// 			{
		// 				Properties: &armcomputeschedule.ScheduledActionProperties{
		// 					ResourceType: to.Ptr(armcomputeschedule.ResourceTypeVirtualMachine),
		// 					ActionType: to.Ptr(armcomputeschedule.ScheduledActionTypeStart),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:55.281Z"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:55.286Z"); return t}()),
		// 					Schedule: &armcomputeschedule.ScheduledActionsSchedule{
		// 						ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse("15:04:05.999999999Z07:00", "19:00:00"); return t}()),
		// 						TimeZone: to.Ptr("g"),
		// 						RequestedWeekDays: []*armcomputeschedule.WeekDay{
		// 							to.Ptr(armcomputeschedule.WeekDayMonday),
		// 						},
		// 						RequestedMonths: []*armcomputeschedule.Month{
		// 							to.Ptr(armcomputeschedule.MonthJanuary),
		// 						},
		// 						RequestedDaysOfTheMonth: []*int32{
		// 							to.Ptr[int32](15),
		// 						},
		// 						ExecutionParameters: &armcomputeschedule.ExecutionParameters{
		// 							OptimizationPreference: to.Ptr(armcomputeschedule.OptimizationPreferenceCost),
		// 							RetryPolicy: &armcomputeschedule.RetryPolicy{
		// 								RetryCount: to.Ptr[int32](17),
		// 								RetryWindowInMinutes: to.Ptr[int32](29),
		// 							},
		// 						},
		// 						DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeUnknown),
		// 					},
		// 					NotificationSettings: []*armcomputeschedule.NotificationProperties{
		// 						{
		// 							Destination: to.Ptr("wbhryycyolvnypjxzlawwvb"),
		// 							Type: to.Ptr(armcomputeschedule.NotificationTypeEmail),
		// 							Language: to.Ptr(armcomputeschedule.LanguageEnUs),
		// 							Disabled: to.Ptr(true),
		// 						},
		// 					},
		// 					Disabled: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armcomputeschedule.ProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key2102": to.Ptr("obwsqwdydpkscnzceopxgkrhrxtdhv"),
		// 				},
		// 				Location: to.Ptr("vmuhgdgipeypkcv"),
		// 				ID: to.Ptr("/subscriptions/83C27AB3-A7B9-498B-B165-D9440661474F/resourceGroups/myRg/providers/Microsoft.ComputeSchedule/scheduledActions/myScheduledAction"),
		// 				Name: to.Ptr("a"),
		// 				Type: to.Ptr("obafnwpw"),
		// 				SystemData: &armcomputeschedule.SystemData{
		// 					CreatedBy: to.Ptr("cvryvreuvvjtiamcwhisrt"),
		// 					CreatedByType: to.Ptr(armcomputeschedule.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:55.288Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("supbnksztdbgulxgvfmqvriqdlpirh"),
		// 					LastModifiedByType: to.Ptr(armcomputeschedule.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:55.288Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
