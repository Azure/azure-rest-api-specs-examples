package armcostmanagement_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/scheduledActions/scheduledAction-createOrUpdate-shared.json
func ExampleScheduledActionsClient_CreateOrUpdateByScope_createOrUpdateScheduledActionByScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().CreateOrUpdateByScope(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "monthlyCostByResource", armcostmanagement.ScheduledAction{
		Kind: to.Ptr(armcostmanagement.ScheduledActionKindEmail),
		Properties: &armcostmanagement.ScheduledActionProperties{
			DisplayName: to.Ptr("Monthly Cost By Resource"),
			FileDestination: &armcostmanagement.FileDestination{
				FileFormats: []*armcostmanagement.FileFormat{
					to.Ptr(armcostmanagement.FileFormatCSV),
				},
			},
			Notification: &armcostmanagement.NotificationProperties{
				Subject: to.Ptr("Cost by resource this month"),
				To: []*string{
					to.Ptr("user@gmail.com"),
					to.Ptr("team@gmail.com"),
				},
			},
			Schedule: &armcostmanagement.ScheduleProperties{
				DaysOfWeek: []*armcostmanagement.DaysOfWeek{
					to.Ptr(armcostmanagement.DaysOfWeekMonday),
				},
				EndDate:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-19T22:21:51.1287144Z"); return t }()),
				Frequency: to.Ptr(armcostmanagement.ScheduleFrequencyMonthly),
				HourOfDay: to.Ptr[int32](10),
				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-19T22:21:51.1287144Z"); return t }()),
				WeeksOfMonth: []*armcostmanagement.WeeksOfMonth{
					to.Ptr(armcostmanagement.WeeksOfMonthFirst),
					to.Ptr(armcostmanagement.WeeksOfMonthThird),
				},
			},
			Status: to.Ptr(armcostmanagement.ScheduledActionStatusEnabled),
			ViewID: to.Ptr("/providers/Microsoft.CostManagement/views/swaggerExample"),
		},
	}, &armcostmanagement.ScheduledActionsClientCreateOrUpdateByScopeOptions{
		IfMatch: to.Ptr("")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcostmanagement.ScheduledActionsClientCreateOrUpdateByScopeResponse{
	// 	ScheduledAction: armcostmanagement.ScheduledAction{
	// 		Name: to.Ptr("monthlyCostByResource"),
	// 		Type: to.Ptr("Microsoft.CostManagement/ScheduledActions"),
	// 		ETag: to.Ptr(azcore.ETag("\"1d4ff9fe66f1d10\"")),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/scheduledActions/monthlyCostByResource"),
	// 		Kind: to.Ptr(armcostmanagement.ScheduledActionKindEmail),
	// 		Properties: &armcostmanagement.ScheduledActionProperties{
	// 			DisplayName: to.Ptr("Monthly Cost By Resource"),
	// 			Notification: &armcostmanagement.NotificationProperties{
	// 				Subject: to.Ptr("Cost by resource this month"),
	// 				To: []*string{
	// 					to.Ptr("user@gmail.com"),
	// 					to.Ptr("team@gmail.com"),
	// 				},
	// 			},
	// 			Schedule: &armcostmanagement.ScheduleProperties{
	// 				DaysOfWeek: []*armcostmanagement.DaysOfWeek{
	// 					to.Ptr(armcostmanagement.DaysOfWeekMonday),
	// 				},
	// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-19T22:21:51.1287144Z"); return t}()),
	// 				Frequency: to.Ptr(armcostmanagement.ScheduleFrequencyMonthly),
	// 				HourOfDay: to.Ptr[int32](10),
	// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-19T22:21:51.1287144Z"); return t}()),
	// 				WeeksOfMonth: []*armcostmanagement.WeeksOfMonth{
	// 					to.Ptr(armcostmanagement.WeeksOfMonthFirst),
	// 					to.Ptr(armcostmanagement.WeeksOfMonthThird),
	// 				},
	// 			},
	// 			Scope: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000"),
	// 			Status: to.Ptr(armcostmanagement.ScheduledActionStatusEnabled),
	// 			ViewID: to.Ptr("/providers/Microsoft.CostManagement/views/swaggerExample"),
	// 		},
	// 		SystemData: &armcostmanagement.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
	// 			CreatedBy: to.Ptr("testuser"),
	// 			CreatedByType: to.Ptr(armcostmanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-19T22:21:51.1287144Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("testuser"),
	// 			LastModifiedByType: to.Ptr(armcostmanagement.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
