package armbillingbenefits_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/EndCustomerInvestmentFundCreditCreate.json
func ExampleCreditsClient_BeginCreate_endCustomerInvestmentFundCreditCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCreditsClient("97ee05f2-07d5-494d-908c-081a197f4277").BeginCreate(ctx, "resource_group_name_01", "credit_20231212", armbillingbenefits.Credit{
		Location: to.Ptr("global"),
		Properties: &armbillingbenefits.CreditProperties{
			Breakdown: []*armbillingbenefits.CreditBreakdownItem{
				{
					Allocation: &armbillingbenefits.Commitment{
						Amount:       to.Ptr[float64](25000),
						CurrencyCode: to.Ptr("USD"),
						Grain:        to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
					},
					Dimensions: []*armbillingbenefits.CreditDimension{
						{
							Key:   to.Ptr("productFamily"),
							Value: to.Ptr("Azure"),
						},
						{
							Key:   to.Ptr("productCode"),
							Value: to.Ptr("00002b30-0000-0260-d348-f3ffcd565473"),
						},
						{
							Key:   to.Ptr("creditType"),
							Value: to.Ptr("EndCustomerInvestmentCredit"),
						},
						{
							Key:   to.Ptr("supplierType"),
							Value: to.Ptr("External"),
						},
					},
					EndAt:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-31T23:59:59.999Z"); return t }()),
					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T00:00:00Z"); return t }()),
				},
			},
			Credit: &armbillingbenefits.Commitment{
				Amount:       to.Ptr[float64](25000),
				CurrencyCode: to.Ptr("USD"),
				Grain:        to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
			},
			EndAt:       to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-31T23:59:59.999Z"); return t }()),
			ProductCode: to.Ptr("00002b30-0000-0260-d348-f3ffcd565473"),
			StartAt:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T00:00:00Z"); return t }()),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
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
	// res = armbillingbenefits.CreditsClientCreateResponse{
	// 	Credit: armbillingbenefits.Credit{
	// 		Name: to.Ptr("credit_20231212"),
	// 		Type: to.Ptr("Microsoft.BillingBenefits/credits"),
	// 		ID: to.Ptr("/subscriptions/97ee05f2-07d5-494d-908c-081a197f4277/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/credits/credit_20231212"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armbillingbenefits.CreditProperties{
	// 			BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}"),
	// 			BillingProfileResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}/billingProfiles/{billingProfileId}"),
	// 			Breakdown: []*armbillingbenefits.CreditBreakdownItem{
	// 				{
	// 					Allocation: &armbillingbenefits.Commitment{
	// 						Amount: to.Ptr[float64](25000),
	// 						CurrencyCode: to.Ptr("USD"),
	// 						Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
	// 					},
	// 					Dimensions: []*armbillingbenefits.CreditDimension{
	// 						{
	// 							Key: to.Ptr("productFamily"),
	// 							Value: to.Ptr("Azure"),
	// 						},
	// 						{
	// 							Key: to.Ptr("productCode"),
	// 							Value: to.Ptr("00002b30-0000-0260-d348-f3ffcd565473"),
	// 						},
	// 						{
	// 							Key: to.Ptr("creditType"),
	// 							Value: to.Ptr("EndCustomerInvestmentCredit"),
	// 						},
	// 						{
	// 							Key: to.Ptr("supplierType"),
	// 							Value: to.Ptr("External"),
	// 						},
	// 					},
	// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-31T23:59:59.999Z"); return t}()),
	// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T00:00:00Z"); return t}()),
	// 				},
	// 			},
	// 			Credit: &armbillingbenefits.Commitment{
	// 				Amount: to.Ptr[float64](25000),
	// 				CurrencyCode: to.Ptr("USD"),
	// 				Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
	// 			},
	// 			CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}/customers/{customerId}"),
	// 			EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-31T23:59:59.999Z"); return t}()),
	// 			ProductCode: to.Ptr("00002b30-0000-0260-d348-f3ffcd565473"),
	// 			StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T00:00:00Z"); return t}()),
	// 			Status: to.Ptr(armbillingbenefits.CreditStatusSucceeded),
	// 		},
	// 		SystemData: &armbillingbenefits.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-12T01:01:01.1075056Z"); return t}()),
	// 			CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-12T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
