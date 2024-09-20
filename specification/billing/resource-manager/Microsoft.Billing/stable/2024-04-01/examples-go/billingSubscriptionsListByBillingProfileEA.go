package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsListByBillingProfileEA.json
func ExampleSubscriptionsClient_NewListByBillingProfilePager_billingSubscriptionsListByBillingProfileEa() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionsClient().NewListByBillingProfilePager("pcn.94077792", "6478903", &armbilling.SubscriptionsClientListByBillingProfileOptions{IncludeDeleted: nil,
		Expand:  nil,
		Filter:  nil,
		OrderBy: nil,
		Top:     nil,
		Skip:    nil,
		Count:   nil,
		Search:  nil,
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
		// page.SubscriptionListResult = armbilling.SubscriptionListResult{
		// 	Value: []*armbilling.Subscription{
		// 		{
		// 			Name: to.Ptr("6b96d3f2-9008-4a9d-912f-f87744185aa3"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/billingSubscriptions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/pcn.94077792/billingProfiles/6478903/billingSubscriptions/6b96d3f2-9008-4a9d-912f-f87744185aa3"),
		// 			Properties: &armbilling.SubscriptionProperties{
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/pcn.94077792/billingProfiles/6478903"),
		// 				BillingProfileName: to.Ptr("47268432"),
		// 				DisplayName: to.Ptr("My Subscription"),
		// 				EnrollmentAccountDisplayName: to.Ptr("billtest332211@outlook.com"),
		// 				EnrollmentAccountID: to.Ptr("172988"),
		// 				Status: to.Ptr(armbilling.BillingSubscriptionStatusActive),
		// 				SubscriptionID: to.Ptr("6b96d3f2-9008-4a9d-912f-f87744185aa3"),
		// 			},
		// 	}},
		// }
	}
}
