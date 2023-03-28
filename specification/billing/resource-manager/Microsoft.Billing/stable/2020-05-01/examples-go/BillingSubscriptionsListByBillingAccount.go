package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByBillingAccount.json
func ExampleSubscriptionsClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionsClient().NewListByBillingAccountPager("{billingAccountName}", nil)
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
		// 			Name: to.Ptr("90000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/{billingAccountName}/billingSubscriptions/90000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.SubscriptionProperties{
		// 				BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 				CostCenter: to.Ptr("ABC1234"),
		// 				DisplayName: to.Ptr("My subscription"),
		// 				InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/22000000-0000-0000-0000-000000000000"),
		// 				LastMonthCharges: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](5000),
		// 				},
		// 				MonthToDateCharges: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](600),
		// 				},
		// 				SKUDescription: to.Ptr("Microsoft Azure Plan"),
		// 				SKUID: to.Ptr("0001"),
		// 				SubscriptionBillingStatus: to.Ptr(armbilling.BillingSubscriptionStatusTypeActive),
		// 				SubscriptionID: to.Ptr("6b96d3f2-9008-4a9d-912f-f87744185aa3"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("90000000-0000-0000-0000-000000000001"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/{billingAccountName}/billingSubscriptions/90000000-0000-0000-0000-000000000001"),
		// 			Properties: &armbilling.SubscriptionProperties{
		// 				BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 				CostCenter: to.Ptr("ABC1234"),
		// 				DisplayName: to.Ptr("Test subscription"),
		// 				InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/22000000-0000-0000-0000-000000000000"),
		// 				LastMonthCharges: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](1000),
		// 				},
		// 				MonthToDateCharges: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](400),
		// 				},
		// 				SKUDescription: to.Ptr("Microsoft Azure Plan"),
		// 				SKUID: to.Ptr("0001"),
		// 				SubscriptionBillingStatus: to.Ptr(armbilling.BillingSubscriptionStatusTypeActive),
		// 				SubscriptionID: to.Ptr("6b96d3f2-9008-4a9d-912f-6b96d3f2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("90000000-0000-0000-0000-000000000002"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/{billingAccountName}/billingSubscriptions/90000000-0000-0000-0000-000000000002"),
		// 			Properties: &armbilling.SubscriptionProperties{
		// 				BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 				CostCenter: to.Ptr("ABC1234"),
		// 				DisplayName: to.Ptr("Dev Subscription"),
		// 				InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/22000000-0000-0000-0000-000000000000"),
		// 				LastMonthCharges: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](6000),
		// 				},
		// 				MonthToDateCharges: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](900),
		// 				},
		// 				SKUDescription: to.Ptr("Microsoft Azure Plan"),
		// 				SKUID: to.Ptr("0001"),
		// 				SubscriptionBillingStatus: to.Ptr(armbilling.BillingSubscriptionStatusTypeActive),
		// 				SubscriptionID: to.Ptr("6b96d3f2-4a9d-9008-912f-f87744185aa3"),
		// 			},
		// 	}},
		// }
	}
}
