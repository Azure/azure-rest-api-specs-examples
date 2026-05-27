package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/ApplicableDiscountsList.json
func ExampleDiscountsClient_NewScopeListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiscountsClient("<subscriptionID>").NewScopeListPager("providers/Microsoft.Billing/billingAccounts/{acctId}", nil)
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
		// page = armbillingbenefits.DiscountsClientScopeListResponse{
		// 	DiscountList: armbillingbenefits.DiscountList{
		// 		Value: []*armbillingbenefits.Discount{
		// 			{
		// 				Name: to.Ptr("13810867107109237"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/applicableDiscounts"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/providers/Microsoft.BillingBenefits/applicableDiscounts/13810867107109237"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armbillingbenefits.EntityTypePrimaryDiscount{
		// 					AppliedScopeType: to.Ptr(armbillingbenefits.DiscountAppliedScopeTypeBillingAccount),
		// 					BenefitResourceID: to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.BillingBenefits/discounts/testprimarydiscount"),
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
		// 					DiscountTypeProperties: &armbillingbenefits.DiscountTypeProductSKU{
		// 						ApplyDiscountOn: to.Ptr(armbillingbenefits.ApplyDiscountOnPurchase),
		// 						Conditions: []*armbillingbenefits.ConditionsItem{
		// 							{
		// 								Type: to.Ptr("equalAny"),
		// 								ConditionName: to.Ptr("Cloud"),
		// 								Value: []*string{
		// 									to.Ptr("US-Sec"),
		// 								},
		// 							},
		// 						},
		// 						DiscountPercentage: to.Ptr[float64](14),
		// 						DiscountType: to.Ptr(armbillingbenefits.DiscountTypeEnumSKU),
		// 						ProductFamilyName: to.Ptr("Azure"),
		// 						ProductID: to.Ptr("DZH318Z0BQ35"),
		// 						SKUID: to.Ptr("0001"),
		// 					},
		// 					DisplayName: to.Ptr("Virtual Machines D Series"),
		// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T23:59:59Z"); return t}()),
		// 					EntityType: to.Ptr(armbillingbenefits.DiscountEntityTypePrimary),
		// 					ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
		// 					ProvisioningState: to.Ptr(armbillingbenefits.DiscountProvisioningStateSucceeded),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T00:00:00Z"); return t}()),
		// 					Status: to.Ptr(armbillingbenefits.DiscountStatusActive),
		// 					SystemID: to.Ptr("13810867107109237"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 					"key2": to.Ptr("value2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("63810867107109237"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/applicableDiscounts"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/providers/Microsoft.BillingBenefits/applicableDiscounts/63810867107109237"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armbillingbenefits.EntityTypeAffiliateDiscount{
		// 					BenefitResourceID: to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.BillingBenefits/discounts/testprimarydiscount"),
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
		// 					BillingProfileResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
		// 					DisplayName: to.Ptr("Virtual Machines D Series"),
		// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T23:59:59Z"); return t}()),
		// 					EntityType: to.Ptr(armbillingbenefits.DiscountEntityTypeAffiliate),
		// 					PrimaryResourceID: to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.BillingBenefits/discounts/testaffiliatediscount"),
		// 					ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
		// 					ProvisioningState: to.Ptr(armbillingbenefits.DiscountProvisioningStateSucceeded),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T00:00:00Z"); return t}()),
		// 					Status: to.Ptr(armbillingbenefits.DiscountStatusActive),
		// 					SystemID: to.Ptr("63810867107109237"),
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
