package armcostmanagement_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/Budgets/CreateOrUpdate/ReservationUtilization/EA/BillingAccountEA-AlertRule-ReservationIdFilter.json
func ExampleBudgetsClient_CreateOrUpdate_createOrUpdateReservationUtilizationBillingAccountEaAlertRuleReservationIdFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBudgetsClient().CreateOrUpdate(ctx, "providers/Microsoft.Billing/billingAccounts/123456", "TestAlertRule", armcostmanagement.Budget{
		ETag: to.Ptr(azcore.ETag("\"1d34d016a593709\"")),
		Properties: &armcostmanagement.BudgetProperties{
			Category: to.Ptr(armcostmanagement.CategoryTypeReservationUtilization),
			Filter: &armcostmanagement.BudgetFilter{
				Dimensions: &armcostmanagement.BudgetComparisonExpression{
					Name:     to.Ptr("ReservationId"),
					Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
					Values: []*string{
						to.Ptr("00000000-0000-0000-0000-000000000000"),
						to.Ptr("00000000-0000-0000-0000-000000000001"),
						to.Ptr("00000000-0000-0000-0000-000000000002"),
					},
				},
			},
			Notifications: map[string]*armcostmanagement.Notification{
				"Actual_LessThan_99_Percent": {
					ContactEmails: []*string{
						to.Ptr("johndoe@contoso.com"),
						to.Ptr("janesmith@contoso.com"),
					},
					Enabled:   to.Ptr(true),
					Frequency: to.Ptr(armcostmanagement.FrequencyWeekly),
					Locale:    to.Ptr(armcostmanagement.CultureCodeEnUs),
					Operator:  to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeLessThan),
					Threshold: to.Ptr[float32](99),
				},
			},
			TimeGrain: to.Ptr(armcostmanagement.TimeGrainTypeLast7Days),
			TimePeriod: &armcostmanagement.BudgetTimePeriod{
				EndDate:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00Z"); return t }()),
				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-01T00:00:00Z"); return t }()),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcostmanagement.BudgetsClientCreateOrUpdateResponse{
	// 	Budget: armcostmanagement.Budget{
	// 		Name: to.Ptr("TestAlertRule"),
	// 		Type: to.Ptr("Microsoft.CostManagement/budgets"),
	// 		ETag: to.Ptr(azcore.ETag("\"1d34d012214157f\"")),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/providers/Microsoft.CostManagement/budgets/TestAlertRule"),
	// 		Properties: &armcostmanagement.BudgetProperties{
	// 			Category: to.Ptr(armcostmanagement.CategoryTypeReservationUtilization),
	// 			Filter: &armcostmanagement.BudgetFilter{
	// 				Dimensions: &armcostmanagement.BudgetComparisonExpression{
	// 					Name: to.Ptr("ReservationId"),
	// 					Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
	// 					Values: []*string{
	// 						to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 						to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 						to.Ptr("00000000-0000-0000-0000-000000000002"),
	// 					},
	// 				},
	// 			},
	// 			Notifications: map[string]*armcostmanagement.Notification{
	// 				"Actual_LessThan_99_Percent": &armcostmanagement.Notification{
	// 					ContactEmails: []*string{
	// 						to.Ptr("johndoe@contoso.com"),
	// 						to.Ptr("janesmith@contoso.com"),
	// 					},
	// 					Enabled: to.Ptr(true),
	// 					Frequency: to.Ptr(armcostmanagement.FrequencyWeekly),
	// 					Locale: to.Ptr(armcostmanagement.CultureCodeEnUs),
	// 					Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeLessThan),
	// 					Threshold: to.Ptr[float32](99),
	// 				},
	// 			},
	// 			TimeGrain: to.Ptr(armcostmanagement.TimeGrainTypeLast7Days),
	// 			TimePeriod: &armcostmanagement.BudgetTimePeriod{
	// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00Z"); return t}()),
	// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-01T00:00:00Z"); return t}()),
	// 			},
	// 		},
	// 	},
	// }
}
