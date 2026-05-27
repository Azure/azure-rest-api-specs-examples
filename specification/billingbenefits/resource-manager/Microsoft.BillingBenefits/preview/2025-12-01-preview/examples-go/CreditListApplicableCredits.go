package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/CreditListApplicableCredits.json
func ExampleCreditsClient_NewListApplicablePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCreditsClient("<subscriptionID>").NewListApplicablePager("providers/Microsoft.Billing/billingAccounts/{acctId}", nil)
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
		// page = armbillingbenefits.CreditsClientListApplicableResponse{
		// 	CreditsList: armbillingbenefits.CreditsList{
		// 		Value: []*armbillingbenefits.Credit{
		// 			{
		// 				Name: to.Ptr("credit_20231212"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/applicableCredits"),
		// 				ID: to.Ptr("/subscriptions/97ee05f2-07d5-494d-908c-081a197f4277/providers/Microsoft.BillingBenefits/applicableCredits/{systemId}"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armbillingbenefits.CreditProperties{
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/12345678"),
		// 					BillingProfileResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1185c37e-d180-5aea-03a3-07e2489791a8:3c9f6394-2029-4f58-bf47-76ebbb03136e_2019-05-31/billingProfiles/6EB6-D7AQ-BG7-TGB"),
		// 					Credit: &armbillingbenefits.Commitment{
		// 						Amount: to.Ptr[float64](20000),
		// 						CurrencyCode: to.Ptr("USD"),
		// 						Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
		// 					},
		// 					CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1185c37e-d180-5aea-03a3-07e2489791a8:3c9f6394-2029-4f58-bf47-76ebbb03136e_2019-05-31/customers/5334c5c9-e4b0-4170-a015-fa8fff5c226a"),
		// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T23:59:59Z"); return t}()),
		// 					ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
		// 					Reason: &armbillingbenefits.CreditReason{
		// 						Description: to.Ptr("Outage credit (SLA)"),
		// 						Code: to.Ptr("125"),
		// 					},
		// 					ResourceID: to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.BillingBenefits/credits/testprimarydiscount"),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T00:00:00Z"); return t}()),
		// 					Status: to.Ptr(armbillingbenefits.CreditStatusSucceeded),
		// 					SystemID: to.Ptr("13810867107109237"),
		// 				},
		// 				SystemData: &armbillingbenefits.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-12T01:01:01.1075056Z"); return t}()),
		// 					CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-12T01:01:01.1075056Z"); return t}()),
		// 					LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 					"key2": to.Ptr("value2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
