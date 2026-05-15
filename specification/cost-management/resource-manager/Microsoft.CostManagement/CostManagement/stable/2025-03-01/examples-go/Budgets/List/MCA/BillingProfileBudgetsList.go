package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/Budgets/List/MCA/BillingProfileBudgetsList.json
func ExampleBudgetsClient_NewListPager_billingProfileBudgetsListMca() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBudgetsClient().NewListPager("providers/Microsoft.Billing/billingAccounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee:ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj_2023-04-01/billingProfiles/MYDEVTESTBP", nil)
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
		// page = armcostmanagement.BudgetsClientListResponse{
		// 	BudgetsListResult: armcostmanagement.BudgetsListResult{
		// 		Value: []*armcostmanagement.Budget{
		// 			{
		// 				Name: to.Ptr("TestBudget1"),
		// 				Type: to.Ptr("Microsoft.CostManagement/budgets"),
		// 				ETag: to.Ptr(azcore.ETag("\"1d34d012214157c\"")),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee:ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj_2023-04-01/billingProfiles/MYDEVTESTBP/providers/Microsoft.CostManagement/budgets/TestBudget1"),
		// 				Properties: &armcostmanagement.BudgetProperties{
		// 					Amount: to.Ptr[float32](200),
		// 					Category: to.Ptr(armcostmanagement.CategoryTypeCost),
		// 					CurrentSpend: &armcostmanagement.CurrentSpend{
		// 						Amount: to.Ptr[float32](30),
		// 						Unit: to.Ptr("USD"),
		// 					},
		// 					Notifications: map[string]*armcostmanagement.Notification{
		// 						"Actual_GreaterThanOrEqualTo_80_Percent": &armcostmanagement.Notification{
		// 							ContactEmails: []*string{
		// 								to.Ptr("johndoe@contoso.com"),
		// 								to.Ptr("janesmith@contoso.com"),
		// 							},
		// 							Enabled: to.Ptr(true),
		// 							Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeGreaterThanOrEqualTo),
		// 							Threshold: to.Ptr[float32](80),
		// 							ThresholdType: to.Ptr(armcostmanagement.ThresholdTypeActual),
		// 						},
		// 					},
		// 					TimeGrain: to.Ptr(armcostmanagement.TimeGrainTypeMonthly),
		// 					TimePeriod: &armcostmanagement.BudgetTimePeriod{
		// 						EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-31T00:00:00Z"); return t}()),
		// 						StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-01T00:00:00Z"); return t}()),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("TestBudget2"),
		// 				Type: to.Ptr("Microsoft.CostManagement/budgets"),
		// 				ETag: to.Ptr(azcore.ETag("\"1d34d012214157d\"")),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee:ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj_2023-04-01/billingProfiles/MYDEVTESTBP/providers/Microsoft.CostManagement/budgets/TestBudget2"),
		// 				Properties: &armcostmanagement.BudgetProperties{
		// 					Amount: to.Ptr[float32](600),
		// 					Category: to.Ptr(armcostmanagement.CategoryTypeCost),
		// 					CurrentSpend: &armcostmanagement.CurrentSpend{
		// 						Amount: to.Ptr[float32](20),
		// 						Unit: to.Ptr("USD"),
		// 					},
		// 					Notifications: map[string]*armcostmanagement.Notification{
		// 						"Actual_GreaterThanOrEqualTo_70_Percent": &armcostmanagement.Notification{
		// 							ContactEmails: []*string{
		// 								to.Ptr("johndoe@contoso.com"),
		// 								to.Ptr("janesmith@contoso.com"),
		// 							},
		// 							Enabled: to.Ptr(true),
		// 							Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeGreaterThanOrEqualTo),
		// 							Threshold: to.Ptr[float32](70),
		// 							ThresholdType: to.Ptr(armcostmanagement.ThresholdTypeActual),
		// 						},
		// 					},
		// 					TimeGrain: to.Ptr(armcostmanagement.TimeGrainTypeMonthly),
		// 					TimePeriod: &armcostmanagement.BudgetTimePeriod{
		// 						EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-31T00:00:00Z"); return t}()),
		// 						StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-01T00:00:00Z"); return t}()),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("TestRUAlert1"),
		// 				Type: to.Ptr("Microsoft.CostManagement/budgets"),
		// 				ETag: to.Ptr(azcore.ETag("\"1d34d012214157f\"")),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee:ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj_2023-04-01/billingProfiles/MYDEVTESTBP/providers/Microsoft.CostManagement/budgets/TestRUAlert1"),
		// 				Properties: &armcostmanagement.BudgetProperties{
		// 					Category: to.Ptr(armcostmanagement.CategoryTypeReservationUtilization),
		// 					Filter: &armcostmanagement.BudgetFilter{
		// 					},
		// 					Notifications: map[string]*armcostmanagement.Notification{
		// 						"Actual_LessThan_99_Percent": &armcostmanagement.Notification{
		// 							ContactEmails: []*string{
		// 								to.Ptr("johndoe@contoso.com"),
		// 								to.Ptr("janesmith@contoso.com"),
		// 							},
		// 							Enabled: to.Ptr(true),
		// 							Frequency: to.Ptr(armcostmanagement.FrequencyDaily),
		// 							Locale: to.Ptr(armcostmanagement.CultureCodeEnUs),
		// 							Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeLessThan),
		// 							Threshold: to.Ptr[float32](99),
		// 						},
		// 					},
		// 					TimeGrain: to.Ptr(armcostmanagement.TimeGrainTypeLast30Days),
		// 					TimePeriod: &armcostmanagement.BudgetTimePeriod{
		// 						EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00Z"); return t}()),
		// 						StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-01T00:00:00Z"); return t}()),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("TestRUAlert2"),
		// 				Type: to.Ptr("Microsoft.CostManagement/budgets"),
		// 				ETag: to.Ptr(azcore.ETag("\"1d34d012214157f\"")),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee:ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj_2023-04-01/billingProfiles/MYDEVTESTBP/providers/Microsoft.CostManagement/budgets/TestRUAlert2"),
		// 				Properties: &armcostmanagement.BudgetProperties{
		// 					Category: to.Ptr(armcostmanagement.CategoryTypeReservationUtilization),
		// 					Filter: &armcostmanagement.BudgetFilter{
		// 						Dimensions: &armcostmanagement.BudgetComparisonExpression{
		// 							Name: to.Ptr("ReservationId"),
		// 							Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
		// 							Values: []*string{
		// 								to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 								to.Ptr("00000000-0000-0000-0000-000000000001"),
		// 								to.Ptr("00000000-0000-0000-0000-000000000002"),
		// 							},
		// 						},
		// 					},
		// 					Notifications: map[string]*armcostmanagement.Notification{
		// 						"Actual_LessThan_99_Percent": &armcostmanagement.Notification{
		// 							ContactEmails: []*string{
		// 								to.Ptr("johndoe@contoso.com"),
		// 								to.Ptr("janesmith@contoso.com"),
		// 							},
		// 							Enabled: to.Ptr(true),
		// 							Frequency: to.Ptr(armcostmanagement.FrequencyDaily),
		// 							Locale: to.Ptr(armcostmanagement.CultureCodeEnUs),
		// 							Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeLessThan),
		// 							Threshold: to.Ptr[float32](80),
		// 						},
		// 					},
		// 					TimeGrain: to.Ptr(armcostmanagement.TimeGrainTypeLast7Days),
		// 					TimePeriod: &armcostmanagement.BudgetTimePeriod{
		// 						EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00Z"); return t}()),
		// 						StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-01T00:00:00Z"); return t}()),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
