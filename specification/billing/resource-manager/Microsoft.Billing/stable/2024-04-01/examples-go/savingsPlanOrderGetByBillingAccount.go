package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/savingsPlanOrderGetByBillingAccount.json
func ExampleSavingsPlanOrdersClient_GetByBillingAccount_savingsPlanOrderGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlanOrdersClient().GetByBillingAccount(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "20000000-0000-0000-0000-000000000000", &armbilling.SavingsPlanOrdersClientGetByBillingAccountOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SavingsPlanOrderModel = armbilling.SavingsPlanOrderModel{
	// 	Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("microsoft.billing/billingAccounts/savingsPlanOrders"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/savingsPlanOrders/20000000-0000-0000-0000-000000000000"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armbilling.SavingsPlanOrderModelProperties{
	// 		BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-20T02:36:22.339Z"); return t}()),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
	// 		BillingPlan: to.Ptr(armbilling.BillingPlanP1M),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-BBBB-CCC-DDD"),
	// 		BillingScopeID: to.Ptr("10000000-0000-0000-0000-000000000000"),
	// 		DisplayName: to.Ptr("SP1"),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-20T02:36:22.339Z"); return t}()),
	// 		ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SavingsPlans: []*string{
	// 			to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000")},
	// 			Term: to.Ptr(armbilling.SavingsPlanTermP3Y),
	// 		},
	// 		SKU: &armbilling.SKU{
	// 			Name: to.Ptr("Compute_Savings_Plan"),
	// 		},
	// 	}
}
