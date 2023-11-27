package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/BenefitUtilizationSummaries/SavingsPlan-BillingProfile.json
func ExampleBenefitUtilizationSummariesClient_NewListByBillingProfileIDPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBenefitUtilizationSummariesClient().NewListByBillingProfileIDPager("c0a00000-0e04-5ee3-000e-f0c6e00000ec:c0a00000-0e04-5ee3-000e-f0c6e00000ec", "200e5e90-000e-4960-8dcd-8d00a02db000", &armcostmanagement.BenefitUtilizationSummariesClientListByBillingProfileIDOptions{GrainParameter: nil,
		Filter: to.Ptr("properties/usageDate ge 2022-10-15 and properties/usageDate le 2022-10-18"),
	})
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
		// page.BenefitUtilizationSummariesListResult = armcostmanagement.BenefitUtilizationSummariesListResult{
		// 	Value: []armcostmanagement.BenefitUtilizationSummaryClassification{
		// 		&armcostmanagement.SavingsPlanUtilizationSummary{
		// 			Name: to.Ptr("66cccc66-6ccc-6c66-666c-66cc6c6c66c6_222d22dd-d2d2-2dd2-222d-2dd2222ddddd_20211116"),
		// 			Type: to.Ptr("Microsoft.CostManagement/benefitUtilizationSummaries"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/c0a00000-0e04-5ee3-000e-f0c6e00000ec:c0a00000-0e04-5ee3-000e-f0c6e00000ec/billingProfiles/200e5e90-000e-4960-8dcd-8d00a02db000/providers/Microsoft.CostManagement/benefitUtilizationSummaries/66cccc66-6ccc-6c66-666c-66cc6c6c66c6_222d22dd-d2d2-2dd2-222d-2dd2222ddddd_20211116"),
		// 			Kind: to.Ptr(armcostmanagement.BenefitKindSavingsPlan),
		// 			Properties: &armcostmanagement.SavingsPlanUtilizationSummaryProperties{
		// 				ArmSKUName: to.Ptr("Compute_Savings_Plan"),
		// 				BenefitID: to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/66cccc66-6ccc-6c66-666c-66cc6c6c66c6/savingsPlans/222d22dd-d2d2-2dd2-222d-2dd2222ddddd"),
		// 				BenefitOrderID: to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/66cccc66-6ccc-6c66-666c-66cc6c6c66c6"),
		// 				BenefitType: to.Ptr(armcostmanagement.BenefitKindSavingsPlan),
		// 				UsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-16T00:00:00.000Z"); return t}()),
		// 				AvgUtilizationPercentage: to.Ptr[float64](90),
		// 				MaxUtilizationPercentage: to.Ptr[float64](100),
		// 				MinUtilizationPercentage: to.Ptr[float64](80),
		// 			},
		// 	}},
		// }
	}
}
