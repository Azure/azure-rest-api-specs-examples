package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/ConditionalCreditUpdateContributor.json
func ExampleConditionalCreditsClient_BeginUpdate_conditionalCreditUpdateContributor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConditionalCreditsClient("20000000-0000-0000-0000-000000000000").BeginUpdate(ctx, "resource_group_name_02", "conditionalCreditContributor_20250801", armbillingbenefits.ConditionalCreditPatchRequest{
		Properties: &armbillingbenefits.ConditionalCreditPatchRequestProperties{
			DisplayName: to.Ptr("Updated Contributor Conditional Credit 20250801"),
		},
		Tags: map[string]*string{
			"environment": to.Ptr("test"),
			"team":        to.Ptr("finance"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbillingbenefits.ConditionalCreditsClientUpdateResponse{
	// 	ConditionalCredit: armbillingbenefits.ConditionalCredit{
	// 		Name: to.Ptr("conditionalCreditContributor_20250801"),
	// 		Type: to.Ptr("Microsoft.BillingBenefits/conditionalCredits"),
	// 		ID: to.Ptr("/subscriptions/{contributorCloudSubId}/resourceGroups/resource_group_name_02/providers/Microsoft.BillingBenefits/conditionalCredits/conditionalCreditContributor_20250801"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armbillingbenefits.ContributorConditionalCreditProperties{
	// 			BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:11111111-1111-1111-1111-111111111111_2025-10-28"),
	// 			DisplayName: to.Ptr("Updated Contributor Conditional Credit 20250801"),
	// 			EntityType: to.Ptr(armbillingbenefits.ConditionalCreditEntityTypeContributor),
	// 			Milestones: []*armbillingbenefits.ContributorConditionalCreditMilestone{
	// 				{
	// 					Award: &armbillingbenefits.Award{
	// 						BalanceVersion: to.Ptr[float32](13),
	// 						Credit: &armbillingbenefits.Commitment{
	// 							Amount: to.Ptr[float64](5000),
	// 							CurrencyCode: to.Ptr("USD"),
	// 							Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
	// 						},
	// 						Duration: to.Ptr(armbillingbenefits.Term("P3M")),
	// 						EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
	// 						StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-01T00:00:00Z"); return t}()),
	// 						SystemID: to.Ptr("credit98765432"),
	// 					},
	// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-30T23:59:59Z"); return t}()),
	// 					MilestoneID: to.Ptr("550e8400-e29b-41d4-a716-446655440001"),
	// 					SpendTarget: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](50000),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.MilestoneStatusCompleted),
	// 				},
	// 				{
	// 					Award: &armbillingbenefits.Award{
	// 						Credit: &armbillingbenefits.Commitment{
	// 							Amount: to.Ptr[float64](5000),
	// 							CurrencyCode: to.Ptr("USD"),
	// 							Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
	// 						},
	// 						Duration: to.Ptr(armbillingbenefits.Term("P3M")),
	// 					},
	// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
	// 					MilestoneID: to.Ptr("550e8400-e29b-41d4-a716-446655440002"),
	// 					SpendTarget: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](100000),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.MilestoneStatusActive),
	// 				},
	// 			},
	// 			PrimaryBillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}"),
	// 			PrimaryResourceID: to.Ptr("/subscriptions/{primaryCloudSubId}/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/conditionalCredits/conditionalCredit_20250801"),
	// 			ProductCode: to.Ptr("000187f7-0000-0260-ab43-b8473ce57f1d"),
	// 			ProvisioningState: to.Ptr(armbillingbenefits.ConditionalCreditsProvisioningStateSucceeded),
	// 			StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-01T00:00:00Z"); return t}()),
	// 			Status: to.Ptr(armbillingbenefits.ConditionalCreditStatusActive),
	// 			SystemID: to.Ptr("CONDITIONALCREDITS-SYSTEM-20250801"),
	// 		},
	// 		SystemData: &armbillingbenefits.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-01T01:01:01.1075056Z"); return t}()),
	// 			CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-01T02:15:30.2080047Z"); return t}()),
	// 			LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("test"),
	// 			"team": to.Ptr("finance"),
	// 		},
	// 	},
	// }
}
