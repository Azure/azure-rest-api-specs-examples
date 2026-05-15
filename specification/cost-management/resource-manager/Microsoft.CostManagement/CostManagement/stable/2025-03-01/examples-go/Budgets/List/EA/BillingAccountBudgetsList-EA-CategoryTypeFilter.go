package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/Budgets/List/EA/BillingAccountBudgetsList-EA-CategoryTypeFilter.json
func ExampleBudgetsClient_NewListPager_billingAccountBudgetsListEaCategoryTypeFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBudgetsClient().NewListPager("providers/Microsoft.Billing/billingAccounts/123456", &armcostmanagement.BudgetsClientListOptions{
		Filter: to.Ptr("properties/category eq 'ReservationUtilization'")})
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
		// 				Name: to.Ptr("TestRUAlert"),
		// 				Type: to.Ptr("Microsoft.CostManagement/budgets"),
		// 				ETag: to.Ptr(azcore.ETag("\"1d34d012214157f\"")),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/providers/Microsoft.CostManagement/budgets/TestRUAlert"),
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
		// 							Frequency: to.Ptr(armcostmanagement.FrequencyWeekly),
		// 							Locale: to.Ptr(armcostmanagement.CultureCodeEnUs),
		// 							Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeLessThan),
		// 							Threshold: to.Ptr[float32](99),
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
