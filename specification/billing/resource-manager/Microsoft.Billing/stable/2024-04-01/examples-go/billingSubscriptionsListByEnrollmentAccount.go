package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsListByEnrollmentAccount.json
func ExampleSubscriptionsClient_NewListByEnrollmentAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionsClient().NewListByEnrollmentAccountPager("6564892", "172988", &armbilling.SubscriptionsClientListByEnrollmentAccountOptions{Filter: nil,
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
		// 			Name: to.Ptr("90000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/6564892/billingSubscriptions/90000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.SubscriptionProperties{
		// 				ConsumptionCostCenter: to.Ptr("123"),
		// 				DisplayName: to.Ptr("TestSubscription"),
		// 				EnrollmentAccountDisplayName: to.Ptr("billtest332211@outlook.com"),
		// 				EnrollmentAccountID: to.Ptr("172988"),
		// 				EnrollmentAccountSubscriptionDetails: &armbilling.EnrollmentAccountSubscriptionDetails{
		// 					EnrollmentAccountStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T00:00:00.000Z"); return t}()),
		// 				},
		// 				SubscriptionID: to.Ptr("90000000-0000-0000-0000-000000000000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("90000000-0000-0000-0000-000000000001"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/6564892/billingSubscriptions/90000000-0000-0000-0000-000000000001"),
		// 			Properties: &armbilling.SubscriptionProperties{
		// 				ConsumptionCostCenter: to.Ptr(""),
		// 				DisplayName: to.Ptr("Microsoft Azure Enterprise-1"),
		// 				EnrollmentAccountDisplayName: to.Ptr("billtest332211@outlook.com"),
		// 				EnrollmentAccountID: to.Ptr("172988"),
		// 				EnrollmentAccountSubscriptionDetails: &armbilling.EnrollmentAccountSubscriptionDetails{
		// 					EnrollmentAccountStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T00:00:00.000Z"); return t}()),
		// 					SubscriptionEnrollmentAccountStatus: to.Ptr(armbilling.SubscriptionEnrollmentAccountStatusDeleted),
		// 				},
		// 				SubscriptionID: to.Ptr("90000000-0000-0000-0000-000000000001"),
		// 			},
		// 	}},
		// }
	}
}
