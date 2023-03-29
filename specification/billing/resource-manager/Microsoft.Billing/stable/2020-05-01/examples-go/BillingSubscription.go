package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscription.json
func ExampleSubscriptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionsClient().Get(ctx, "{billingAccountName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Subscription = armbilling.Subscription{
	// 	Name: to.Ptr("{subscriptionId}"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptions"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/{billingAccountName}/billingSubscriptions/{subscriptionId}"),
	// 	Properties: &armbilling.SubscriptionProperties{
	// 		BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
	// 		CostCenter: to.Ptr("ABC1234"),
	// 		DisplayName: to.Ptr("My Subscription"),
	// 		InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
	// 		InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000/invoiceSections/22000000-0000-0000-0000-000000000000"),
	// 		LastMonthCharges: &armbilling.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](5000),
	// 		},
	// 		MonthToDateCharges: &armbilling.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](600),
	// 		},
	// 		SKUDescription: to.Ptr("Microsoft Azure Plan"),
	// 		SKUID: to.Ptr("0001"),
	// 		SubscriptionBillingStatus: to.Ptr(armbilling.BillingSubscriptionStatusTypeActive),
	// 		SubscriptionID: to.Ptr("6b96d3f2-9008-4a9d-912f-f87744185aa3"),
	// 	},
	// }
}
