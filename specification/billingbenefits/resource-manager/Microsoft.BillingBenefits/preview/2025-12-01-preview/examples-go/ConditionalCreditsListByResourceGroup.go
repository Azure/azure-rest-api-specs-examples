package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/ConditionalCreditsListByResourceGroup.json
func ExampleConditionalCreditsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConditionalCreditsClient("10000000-0000-0000-0000-000000000000").NewListByResourceGroupPager("resource_group_name_01", nil)
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
		// page = armbillingbenefits.ConditionalCreditsClientListByResourceGroupResponse{
		// 	ConditionalCreditList: armbillingbenefits.ConditionalCreditList{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/conditionalCredits?api-version=2025-05-01-preview&$skiptoken=abcd1234"),
		// 		Value: []*armbillingbenefits.ConditionalCredit{
		// 			{
		// 				Name: to.Ptr("conditionalCredit_20250801"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/conditionalCredits"),
		// 				ID: to.Ptr("/subscriptions/{primaryCloudSubId}/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/conditionalCredits/conditionalCredit_20250801"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armbillingbenefits.PrimaryConditionalCreditProperties{
		// 					AllowContributors: to.Ptr(armbillingbenefits.EnablementModeEnabled),
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:11111111-1111-1111-1111-111111111111_2025-10-28"),
		// 					DisplayName: to.Ptr("Conditional Credit 20250801"),
		// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
		// 					EntityType: to.Ptr(armbillingbenefits.ConditionalCreditEntityTypePrimary),
		// 					Milestones: []*armbillingbenefits.ConditionalCreditMilestone{
		// 						{
		// 							Award: &armbillingbenefits.Award{
		// 								BalanceVersion: to.Ptr[float32](13),
		// 								Credit: &armbillingbenefits.Commitment{
		// 									Amount: to.Ptr[float64](5000),
		// 									CurrencyCode: to.Ptr("USD"),
		// 									Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
		// 								},
		// 								Duration: to.Ptr(armbillingbenefits.Term("P3M")),
		// 								EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
		// 								StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-01T00:00:00Z"); return t}()),
		// 								SystemID: to.Ptr("credit98765432"),
		// 							},
		// 							EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-30T23:59:59Z"); return t}()),
		// 							MilestoneID: to.Ptr("550e8400-e29b-41d4-a716-446655440001"),
		// 							SpendTarget: &armbillingbenefits.Price{
		// 								Amount: to.Ptr[float64](50000),
		// 								CurrencyCode: to.Ptr("USD"),
		// 							},
		// 							Status: to.Ptr(armbillingbenefits.MilestoneStatusCompleted),
		// 						},
		// 						{
		// 							Award: &armbillingbenefits.Award{
		// 								Credit: &armbillingbenefits.Commitment{
		// 									Amount: to.Ptr[float64](10000),
		// 									CurrencyCode: to.Ptr("USD"),
		// 									Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
		// 								},
		// 								Duration: to.Ptr(armbillingbenefits.Term("P3M")),
		// 							},
		// 							EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
		// 							MilestoneID: to.Ptr("550e8400-e29b-41d4-a716-446655440002"),
		// 							SpendTarget: &armbillingbenefits.Price{
		// 								Amount: to.Ptr[float64](100000),
		// 								CurrencyCode: to.Ptr("USD"),
		// 							},
		// 							Status: to.Ptr(armbillingbenefits.MilestoneStatusActive),
		// 						},
		// 					},
		// 					ProductCode: to.Ptr("000187f7-0000-0260-ab43-b8473ce57f1d"),
		// 					ProvisioningState: to.Ptr(armbillingbenefits.ConditionalCreditsProvisioningStateSucceeded),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-01T00:00:00Z"); return t}()),
		// 					Status: to.Ptr(armbillingbenefits.ConditionalCreditStatusActive),
		// 					SystemID: to.Ptr("CACO-SYSTEM-20250801"),
		// 				},
		// 				SystemData: &armbillingbenefits.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-01T01:01:01.1075056Z"); return t}()),
		// 					CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-01T01:01:01.1075056Z"); return t}()),
		// 					LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 					"key2": to.Ptr("value2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("conditionalCredit_simple_20250801"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/conditionalCredits"),
		// 				ID: to.Ptr("/subscriptions/{primaryCloudSubId}/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/conditionalCredits/conditionalCredit_simple_20250801"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armbillingbenefits.PrimaryConditionalCreditProperties{
		// 					AllowContributors: to.Ptr(armbillingbenefits.EnablementModeDisabled),
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:11111111-1111-1111-1111-111111111111_2025-10-28"),
		// 					DisplayName: to.Ptr("Simple Conditional Credit"),
		// 					EntityType: to.Ptr(armbillingbenefits.ConditionalCreditEntityTypePrimary),
		// 					ProductCode: to.Ptr("000187f7-0000-0260-ab43-b8473ce57f1d"),
		// 					ProvisioningState: to.Ptr(armbillingbenefits.ConditionalCreditsProvisioningStateSucceeded),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-01T00:00:00Z"); return t}()),
		// 					Status: to.Ptr(armbillingbenefits.ConditionalCreditStatusActive),
		// 					SystemID: to.Ptr("CACO-SYSTEM-SIMPLE-20250801"),
		// 				},
		// 				SystemData: &armbillingbenefits.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-01T01:01:01.1075056Z"); return t}()),
		// 					CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-01T01:01:01.1075056Z"); return t}()),
		// 					LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
