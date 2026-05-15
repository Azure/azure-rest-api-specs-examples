package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/Budgets/Get/ReservationUtilization/Get-ReservationUtilization-AlertRule.json
func ExampleBudgetsClient_Get_getReservationUtilizationAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBudgetsClient().Get(ctx, "providers/Microsoft.Billing/billingAccounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee:ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj_2023-04-01/billingProfiles/KKKK-LLLL-MMM-NNN", "TestAlertRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcostmanagement.BudgetsClientGetResponse{
	// 	Budget: armcostmanagement.Budget{
	// 		Name: to.Ptr("TestAlertRule"),
	// 		Type: to.Ptr("Microsoft.CostManagement/budgets"),
	// 		ETag: to.Ptr(azcore.ETag("\"1d34d012214157f\"")),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee:ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj_2023-04-01/billingProfiles/KKKK-LLLL-MMM-NNN/providers/Microsoft.CostManagement/budgets/TestAlertRule"),
	// 		Properties: &armcostmanagement.BudgetProperties{
	// 			Category: to.Ptr(armcostmanagement.CategoryTypeReservationUtilization),
	// 			Filter: &armcostmanagement.BudgetFilter{
	// 				Dimensions: &armcostmanagement.BudgetComparisonExpression{
	// 					Name: to.Ptr("ReservedResourceType"),
	// 					Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
	// 					Values: []*string{
	// 						to.Ptr("VirtualMachines"),
	// 						to.Ptr("SqlDatabases"),
	// 						to.Ptr("CosmosDb"),
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
	// 					Frequency: to.Ptr(armcostmanagement.FrequencyDaily),
	// 					Locale: to.Ptr(armcostmanagement.CultureCodeEnUs),
	// 					Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeLessThan),
	// 					Threshold: to.Ptr[float32](99),
	// 				},
	// 			},
	// 			TimeGrain: to.Ptr(armcostmanagement.TimeGrainTypeLast30Days),
	// 			TimePeriod: &armcostmanagement.BudgetTimePeriod{
	// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T00:00:00Z"); return t}()),
	// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-01T00:00:00Z"); return t}()),
	// 			},
	// 		},
	// 	},
	// }
}
