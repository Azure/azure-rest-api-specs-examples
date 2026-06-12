package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/ScheduledActions_Get_MaximumSet_Gen.json
func ExampleScheduledActionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("732116BD-AF31-4E74-9283-B387C44B4A44", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().Get(ctx, "rgcomputeschedule", "scheduled-action-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientGetResponse{
	// 	ScheduledAction: armcomputeschedule.ScheduledAction{
	// 		Properties: &armcomputeschedule.ScheduledActionProperties{
	// 			ResourceType: to.Ptr(armcomputeschedule.ResourceTypeVirtualMachine),
	// 			ActionType: to.Ptr(armcomputeschedule.ScheduledActionTypeStart),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:45.061Z"); return t}()),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:45.062Z"); return t}()),
	// 			Schedule: &armcomputeschedule.ScheduledActionsSchedule{
	// 				ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.TimeOnly, "12:00:00"); return t}()),
	// 				TimeZone: to.Ptr("America/Los_Angeles"),
	// 				RequestedWeekDays: []*armcomputeschedule.WeekDay{
	// 					to.Ptr(armcomputeschedule.WeekDayMonday),
	// 				},
	// 				RequestedMonths: []*armcomputeschedule.Month{
	// 					to.Ptr(armcomputeschedule.MonthJanuary),
	// 				},
	// 				RequestedDaysOfTheMonth: []*int32{
	// 					to.Ptr[int32](18),
	// 				},
	// 				ExecutionParameters: &armcomputeschedule.ExecutionParameters{
	// 					OptimizationPreference: to.Ptr(armcomputeschedule.OptimizationPreferenceCost),
	// 					RetryPolicy: &armcomputeschedule.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](3),
	// 						RetryWindowInMinutes: to.Ptr[int32](30),
	// 						OnFailureAction: to.Ptr(armcomputeschedule.ResourceOperationTypeUnknown),
	// 					},
	// 				},
	// 				DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeUnknown),
	// 			},
	// 			NotificationSettings: []*armcomputeschedule.NotificationProperties{
	// 				{
	// 					Destination: to.Ptr("admin@contoso.com"),
	// 					Type: to.Ptr(armcomputeschedule.NotificationTypeEmail),
	// 					Language: to.Ptr(armcomputeschedule.LanguageEnUs),
	// 					Disabled: to.Ptr(true),
	// 				},
	// 			},
	// 			Disabled: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armcomputeschedule.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 		},
	// 		Location: to.Ptr("eastus2"),
	// 		ID: to.Ptr("/subscriptions/83C27AB3-A7B9-498B-B165-D9440661474F/resourceGroups/myRg/providers/Microsoft.ComputeSchedule/scheduledActions/scheduled-action-01"),
	// 		Name: to.Ptr("scheduled-action-01"),
	// 		Type: to.Ptr("Microsoft.ComputeSchedule/scheduledActions"),
	// 		SystemData: &armcomputeschedule.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armcomputeschedule.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:41.641Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcomputeschedule.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:41.641Z"); return t}()),
	// 		},
	// 	},
	// }
}
