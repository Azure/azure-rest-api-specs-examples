package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/BenefitRecommendationsByBillingAccount.json
func ExampleBenefitRecommendationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBenefitRecommendationsClient().NewListPager("providers/Microsoft.Billing/billingAccounts/123456", &armcostmanagement.BenefitRecommendationsClientListOptions{Filter: to.Ptr("properties/lookBackPeriod eq 'Last7Days' AND properties/term eq 'P1Y'"),
		Orderby: nil,
		Expand:  to.Ptr("properties/usage,properties/allRecommendationDetails"),
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
		// page.BenefitRecommendationsListResult = armcostmanagement.BenefitRecommendationsListResult{
		// 	Value: []*armcostmanagement.BenefitRecommendationModel{
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.CostManagement/benefitRecommendations"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/providers/Microsoft.CostManagement/benefitRecommendations/00000000-0000-0000-0000-000000000000"),
		// 			Kind: to.Ptr(armcostmanagement.BenefitKindSavingsPlan),
		// 			Properties: &armcostmanagement.SharedScopeBenefitRecommendationProperties{
		// 				AllRecommendationDetails: &armcostmanagement.AllSavingsList{
		// 					Value: []*armcostmanagement.AllSavingsBenefitDetails{
		// 						{
		// 							AverageUtilizationPercentage: to.Ptr[float64](99.33),
		// 							BenefitCost: to.Ptr[float64](52.002),
		// 							CommitmentAmount: to.Ptr[float64](0.164),
		// 							CoveragePercentage: to.Ptr[float64](54.609),
		// 							OverageCost: to.Ptr[float64](144.841),
		// 							SavingsAmount: to.Ptr[float64](21.424),
		// 							SavingsPercentage: to.Ptr[float64](9.815),
		// 							TotalCost: to.Ptr[float64](196.843),
		// 							WastageCost: to.Ptr[float64](0.035),
		// 						},
		// 						{
		// 							AverageUtilizationPercentage: to.Ptr[float64](81.474),
		// 							BenefitCost: to.Ptr[float64](83.754),
		// 							CommitmentAmount: to.Ptr[float64](0.161),
		// 							CoveragePercentage: to.Ptr[float64](56.748),
		// 							OverageCost: to.Ptr[float64](120.389),
		// 							SavingsAmount: to.Ptr[float64](14.124),
		// 							SavingsPercentage: to.Ptr[float64](6.47),
		// 							TotalCost: to.Ptr[float64](204.143),
		// 							WastageCost: to.Ptr[float64](0.1),
		// 					}},
		// 				},
		// 				ArmSKUName: to.Ptr("Compute_Savings_Plan"),
		// 				CommitmentGranularity: to.Ptr(armcostmanagement.GrainHourly),
		// 				CostWithoutBenefit: to.Ptr[float64](218.267),
		// 				CurrencyCode: to.Ptr("USD"),
		// 				FirstConsumptionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-18T00:00:00.000Z"); return t}()),
		// 				LastConsumptionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-25T00:00:00.000Z"); return t}()),
		// 				LookBackPeriod: to.Ptr(armcostmanagement.LookBackPeriodLast7Days),
		// 				RecommendationDetails: &armcostmanagement.AllSavingsBenefitDetails{
		// 					AverageUtilizationPercentage: to.Ptr[float64](99.33),
		// 					BenefitCost: to.Ptr[float64](52.002),
		// 					CommitmentAmount: to.Ptr[float64](0.164),
		// 					CoveragePercentage: to.Ptr[float64](54.609),
		// 					OverageCost: to.Ptr[float64](144.841),
		// 					SavingsAmount: to.Ptr[float64](21.424),
		// 					SavingsPercentage: to.Ptr[float64](9.815),
		// 					TotalCost: to.Ptr[float64](196.843),
		// 					WastageCost: to.Ptr[float64](0.035),
		// 				},
		// 				Scope: to.Ptr(armcostmanagement.ScopeShared),
		// 				Term: to.Ptr(armcostmanagement.TermP1Y),
		// 				TotalHours: to.Ptr[int32](168),
		// 				Usage: &armcostmanagement.RecommendationUsageDetails{
		// 					Charges: []*float64{
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](0),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](1),
		// 						to.Ptr[float64](2),
		// 						to.Ptr[float64](2),
		// 						to.Ptr[float64](2),
		// 						to.Ptr[float64](2)},
		// 						UsageGrain: to.Ptr(armcostmanagement.GrainHourly),
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
