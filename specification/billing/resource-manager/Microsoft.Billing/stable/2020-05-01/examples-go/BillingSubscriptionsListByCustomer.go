package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByCustomer.json
func ExampleSubscriptionsClient_NewListByCustomerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionsClient().NewListByCustomerPager("{billingAccountName}", "{customerName}", nil)
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
		// page.SubscriptionsListResult = armbilling.SubscriptionsListResult{
		// 	Value: []*armbilling.Subscription{
		// 		{
		// 			Name: to.Ptr("billingSubscriptionId1"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/{billingAccountName}/billingSubscriptions/billinSubscriptionId1"),
		// 			Properties: &armbilling.SubscriptionProperties{
		// 				CostCenter: to.Ptr("ABC1234"),
		// 				CustomerDisplayName: to.Ptr("Customer1"),
		// 				CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}"),
		// 				DisplayName: to.Ptr("My subscription"),
		// 				LastMonthCharges: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](5000),
		// 				},
		// 				MonthToDateCharges: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](600),
		// 				},
		// 				Reseller: &armbilling.Reseller{
		// 					Description: to.Ptr("Reseller1"),
		// 					ResellerID: to.Ptr("89e87bdf-a2a2-4687-925f-4c18b27bccfd"),
		// 				},
		// 				SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
		// 				SKUID: to.Ptr("0002"),
		// 				SubscriptionBillingStatus: to.Ptr(armbilling.BillingSubscriptionStatusTypeActive),
		// 				SubscriptionID: to.Ptr("6b96d3f2-9008-4a9d-912f-f87744185aa3"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("billingSubscriptionId2"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/{billingAccountName}/billingSubscriptions/billinSubscriptionId2"),
		// 			Properties: &armbilling.SubscriptionProperties{
		// 				CostCenter: to.Ptr("ABC1234"),
		// 				CustomerDisplayName: to.Ptr("Customer1"),
		// 				CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}"),
		// 				DisplayName: to.Ptr("Test subscription"),
		// 				LastMonthCharges: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](1000),
		// 				},
		// 				MonthToDateCharges: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](400),
		// 				},
		// 				Reseller: &armbilling.Reseller{
		// 					Description: to.Ptr("Reseller3"),
		// 					ResellerID: to.Ptr("3b65b5a8-bd4f-4084-90e9-e1bd667a2b19"),
		// 				},
		// 				SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
		// 				SKUID: to.Ptr("0002"),
		// 				SubscriptionBillingStatus: to.Ptr(armbilling.BillingSubscriptionStatusTypeActive),
		// 				SubscriptionID: to.Ptr("6b96d3f2-9008-4a9d-912f-6b96d3f2"),
		// 			},
		// 	}},
		// }
	}
}
