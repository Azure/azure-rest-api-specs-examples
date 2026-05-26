package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/CreditsListByResourceGroup.json
func ExampleCreditsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCreditsClient("97ee05f2-07d5-494d-908c-081a197f4277").NewListByResourceGroupPager("resource_group_name_01", nil)
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
		// page = armbillingbenefits.CreditsClientListByResourceGroupResponse{
		// 	CreditsList: armbillingbenefits.CreditsList{
		// 		Value: []*armbillingbenefits.Credit{
		// 			{
		// 				Name: to.Ptr("credit_20231212"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/credits"),
		// 				ID: to.Ptr("/subscriptions/97ee05f2-07d5-494d-908c-081a197f4277/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/credits/credit_20231212"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armbillingbenefits.CreditProperties{
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{enrollmentId}"),
		// 					Credit: &armbillingbenefits.Commitment{
		// 						Amount: to.Ptr[float64](20000),
		// 						CurrencyCode: to.Ptr("USD"),
		// 						Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
		// 					},
		// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-12T00:00:00Z"); return t}()),
		// 					ProductCode: to.Ptr("0002d726-0000-0160-330f-a0b98cdbbdc4"),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-12T00:00:00Z"); return t}()),
		// 					Status: to.Ptr(armbillingbenefits.CreditStatusPending),
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
