package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledActions-listWithFilter-private.json
func ExampleScheduledActionsClient_NewListPager_privateScheduledActionsListFilterByViewId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScheduledActionsClient().NewListPager(&armcostmanagement.ScheduledActionsClientListOptions{Filter: to.Ptr("properties/viewId eq '/providers/Microsoft.CostManagement/views/swaggerExample'")})
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
		// page.ScheduledActionListResult = armcostmanagement.ScheduledActionListResult{
		// 	Value: []*armcostmanagement.ScheduledAction{
		// 		{
		// 			Name: to.Ptr("monthlyCostByResource"),
		// 			Type: to.Ptr("Microsoft.CostManagement/ScheduledActions"),
		// 			ID: to.Ptr("providers/Microsoft.CostManagement/scheduledActions/monthlyCostByResource"),
		// 			ETag: to.Ptr("\"1d4ff9fe66f1d10\""),
		// 			Kind: to.Ptr(armcostmanagement.ScheduledActionKindEmail),
		// 			SystemData: &armcostmanagement.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
		// 				CreatedBy: to.Ptr("testuser"),
		// 				CreatedByType: to.Ptr(armcostmanagement.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("testuser"),
		// 				LastModifiedByType: to.Ptr(armcostmanagement.CreatedByTypeUser),
		// 			},
		// 			Properties: &armcostmanagement.ScheduledActionProperties{
		// 				DisplayName: to.Ptr("Monthly Cost By Resource"),
		// 				Notification: &armcostmanagement.NotificationProperties{
		// 					Subject: to.Ptr("Cost by resource this month"),
		// 					To: []*string{
		// 						to.Ptr("user@gmail.com"),
		// 						to.Ptr("team@gmail.com")},
		// 					},
		// 					Schedule: &armcostmanagement.ScheduleProperties{
		// 						DaysOfWeek: []*armcostmanagement.DaysOfWeek{
		// 							to.Ptr(armcostmanagement.DaysOfWeekMonday)},
		// 							EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-19T22:21:51.1287144Z"); return t}()),
		// 							Frequency: to.Ptr(armcostmanagement.ScheduleFrequencyMonthly),
		// 							HourOfDay: to.Ptr[int32](10),
		// 							StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-19T22:21:51.1287144Z"); return t}()),
		// 							WeeksOfMonth: []*armcostmanagement.WeeksOfMonth{
		// 								to.Ptr(armcostmanagement.WeeksOfMonthFirst),
		// 								to.Ptr(armcostmanagement.WeeksOfMonthThird)},
		// 							},
		// 							Scope: to.Ptr(""),
		// 							Status: to.Ptr(armcostmanagement.ScheduledActionStatusEnabled),
		// 							ViewID: to.Ptr("/providers/Microsoft.CostManagement/views/swaggerExample"),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("dailyCostByResource"),
		// 						Type: to.Ptr("Microsoft.CostManagement/ScheduledActions"),
		// 						ID: to.Ptr("providers/Microsoft.CostManagement/scheduledActions/dailyCostByResource"),
		// 						ETag: to.Ptr("\"1d4ff9fe66f1d15\""),
		// 						Kind: to.Ptr(armcostmanagement.ScheduledActionKindEmail),
		// 						SystemData: &armcostmanagement.SystemData{
		// 							CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
		// 							CreatedBy: to.Ptr("testuser"),
		// 							CreatedByType: to.Ptr(armcostmanagement.CreatedByTypeUser),
		// 							LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
		// 							LastModifiedBy: to.Ptr("testuser"),
		// 							LastModifiedByType: to.Ptr(armcostmanagement.CreatedByTypeUser),
		// 						},
		// 						Properties: &armcostmanagement.ScheduledActionProperties{
		// 							DisplayName: to.Ptr("Daily Cost By Resource"),
		// 							Notification: &armcostmanagement.NotificationProperties{
		// 								Subject: to.Ptr("Daily Cost By Resource"),
		// 								To: []*string{
		// 									to.Ptr("user@gmail.com"),
		// 									to.Ptr("team@gmail.com")},
		// 								},
		// 								Schedule: &armcostmanagement.ScheduleProperties{
		// 									EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-19T22:21:51.1287144Z"); return t}()),
		// 									Frequency: to.Ptr(armcostmanagement.ScheduleFrequencyDaily),
		// 									HourOfDay: to.Ptr[int32](12),
		// 									StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-19T22:21:51.1287144Z"); return t}()),
		// 								},
		// 								Scope: to.Ptr(""),
		// 								Status: to.Ptr(armcostmanagement.ScheduledActionStatusDisabled),
		// 								ViewID: to.Ptr("/providers/Microsoft.CostManagement/views/swaggerExample"),
		// 							},
		// 					}},
		// 				}
	}
}
