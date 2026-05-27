package armbillingbenefits_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/DiscountsCreateAffiliate.json
func ExampleDiscountsClient_BeginCreate_discountsCreateAffiliate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiscountsClient("30000000-0000-0000-0000-000000000000").BeginCreate(ctx, "testrg", "testaffiliatediscount", armbillingbenefits.Discount{
		Location: to.Ptr("global"),
		Properties: &armbillingbenefits.EntityTypeAffiliateDiscount{
			DisplayName: to.Ptr("Virtual Machines D Series"),
			EntityType:  to.Ptr(armbillingbenefits.DiscountEntityTypeAffiliate),
			ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
			StartAt:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T00:00:00Z"); return t }()),
			SystemID:    to.Ptr("13810867107109237"),
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
	// res = armbillingbenefits.DiscountsClientCreateResponse{
	// 	Discount: armbillingbenefits.Discount{
	// 		Name: to.Ptr("testprimarydiscount"),
	// 		Type: to.Ptr("Microsoft.BillingBenefits/discounts"),
	// 		ID: to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.BillingBenefits/discounts/testprimarydiscount"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armbillingbenefits.EntityTypeAffiliateDiscount{
	// 			BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
	// 			BillingProfileResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
	// 			CustomerResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/customers/40000000-0000-0000-0000-000000000000"),
	// 			DisplayName: to.Ptr("Virtual Machines D Series"),
	// 			EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T23:59:59Z"); return t}()),
	// 			EntityType: to.Ptr(armbillingbenefits.DiscountEntityTypeAffiliate),
	// 			PrimaryResourceID: to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.BillingBenefits/discounts/testaffiliatediscount"),
	// 			ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
	// 			ProvisioningState: to.Ptr(armbillingbenefits.DiscountProvisioningStateSucceeded),
	// 			StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T00:00:00Z"); return t}()),
	// 			Status: to.Ptr(armbillingbenefits.DiscountStatusActive),
	// 			SystemID: to.Ptr("13810867107109237"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
