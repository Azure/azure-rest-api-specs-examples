package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/ConditionalCreditContributorsListFromApplicableConditionalCredit.json
func ExampleConditionalCreditContributorsClient_NewListFromApplicableConditionalCreditPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConditionalCreditContributorsClient("<subscriptionID>").NewListFromApplicableConditionalCreditPager("20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31", "C20251028000000000000", nil)
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
		// page = armbillingbenefits.ConditionalCreditContributorsClientListFromApplicableConditionalCreditResponse{
		// 	ConditionalCreditContributorList: armbillingbenefits.ConditionalCreditContributorList{
		// 		Value: []*armbillingbenefits.ConditionalCreditContributor{
		// 			{
		// 				Name: to.Ptr("contributor_caco_20250801"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/applicableContributors"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}/providers/Microsoft.BillingBenefits/applicableConditionalCredits/C20251028000000000000/providers/Microsoft.BillingBenefits/applicableContributors/contributor_caco_20250801"),
		// 				Properties: &armbillingbenefits.ContributorConditionalCreditProperties{
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}"),
		// 					DisplayName: to.Ptr("Contributor Conditional Credit 20250801"),
		// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
		// 					EntityType: to.Ptr(armbillingbenefits.ConditionalCreditEntityTypeContributor),
		// 					Milestones: []*armbillingbenefits.ContributorConditionalCreditMilestone{
		// 						{
		// 							Name: to.Ptr("Q4 2025 Milestone"),
		// 							Award: &armbillingbenefits.Award{
		// 								BalanceVersion: to.Ptr[float32](13),
		// 								Credit: &armbillingbenefits.Commitment{
		// 									Amount: to.Ptr[float64](5000),
		// 									CurrencyCode: to.Ptr("USD"),
		// 									Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
		// 								},
		// 								Duration: to.Ptr(armbillingbenefits.Term("P3M")),
		// 								SystemID: to.Ptr("credit98765432"),
		// 							},
		// 							EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
		// 							MilestoneID: to.Ptr("550e8400-e29b-41d4-a716-446655440001"),
		// 							SpendTarget: &armbillingbenefits.Price{
		// 								Amount: to.Ptr[float64](50000),
		// 								CurrencyCode: to.Ptr("USD"),
		// 							},
		// 							Status: to.Ptr(armbillingbenefits.MilestoneStatusActive),
		// 						},
		// 					},
		// 					PrimaryBillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}"),
		// 					PrimaryResourceID: to.Ptr("/subscriptions/{primaryCloudSubId}/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/conditionalCredits/conditionalCredit_20250801"),
		// 					ProductCode: to.Ptr("000187f7-0000-0260-ab43-b8473ce57f1d"),
		// 					ProvisioningState: to.Ptr(armbillingbenefits.ConditionalCreditsProvisioningStateSucceeded),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-01T00:00:00Z"); return t}()),
		// 					Status: to.Ptr(armbillingbenefits.ConditionalCreditStatusActive),
		// 					SystemID: to.Ptr("C20251028000000000000"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
