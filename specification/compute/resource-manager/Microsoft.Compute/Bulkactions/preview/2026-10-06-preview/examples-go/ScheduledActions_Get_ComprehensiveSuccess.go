package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/ScheduledActions_Get_ComprehensiveSuccess.json
func ExampleScheduledActionsClient_Get_twoGetARecurringScheduledActionWithCompleteConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().Get(ctx, "example-rg", "weekday-start", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.ScheduledActionsClientGetResponse{
	// 	ScheduledAction: armbulkactions.ScheduledAction{
	// 		Properties: &armbulkactions.ScheduledActionProperties{
	// 			ResourceType: to.Ptr(armbulkactions.ResourceTypeVirtualMachine),
	// 			ActionType: to.Ptr(armbulkactions.ScheduledActionTypeStart),
	// 			StartTime: to.Ptr(time.Date(2026, time.September, 1, 19, 0, 0, 0, time.FixedZone("", -25200))),
	// 			EndTime: to.Ptr(time.Date(2027, time.September, 1, 19, 0, 0, 0, time.FixedZone("", -25200))),
	// 			Schedule: &armbulkactions.ScheduledActionsSchedule{
	// 				ScheduledTime: to.Ptr(time.Date(0, time.January, 1, 19, 0, 0, 0, time.UTC)),
	// 				TimeZone: to.Ptr("America/Los_Angeles"),
	// 				RequestedWeekDays: []*armbulkactions.WeekDay{
	// 					to.Ptr(armbulkactions.WeekDayMonday),
	// 				},
	// 				RequestedMonths: []*armbulkactions.Month{
	// 					to.Ptr(armbulkactions.MonthJanuary),
	// 				},
	// 				RequestedDaysOfTheMonth: []*int32{
	// 					to.Ptr[int32](15),
	// 				},
	// 			},
	// 			NotificationSettings: []*armbulkactions.NotificationProperties{
	// 				{
	// 					Destination: to.Ptr("admin@contoso.com"),
	// 					Type: to.Ptr(armbulkactions.NotificationTypeEmail),
	// 					Language: to.Ptr(armbulkactions.LanguageEnUs),
	// 					Disabled: to.Ptr(true),
	// 				},
	// 			},
	// 			Disabled: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armbulkactions.ScheduledActionsProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/scheduledActions/weekday-start"),
	// 		Name: to.Ptr("weekday-start"),
	// 		Type: to.Ptr("Microsoft.Compute/scheduledActions"),
	// 		SystemData: &armbulkactions.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2025, time.April, 17, 0, 23, 55, 288000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2025, time.April, 17, 0, 23, 55, 288000000, time.UTC)),
	// 		},
	// 	},
	// }
}
