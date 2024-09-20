package armbilling_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/transactionsListByInvoiceSection.json
func ExampleTransactionsClient_NewListByInvoiceSectionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTransactionsClient().NewListByInvoiceSectionPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "22000000-0000-0000-0000-000000000000", func() time.Time { t, _ := time.Parse("2006-01-02", "2024-04-01"); return t }(), func() time.Time { t, _ := time.Parse("2006-01-02", "2023-05-30"); return t }(), armbilling.TransactionTypeBilled, &armbilling.TransactionsClientListByInvoiceSectionOptions{Filter: to.Ptr("properties/date gt '2020-10-01'"),
		OrderBy: nil,
		Top:     nil,
		Skip:    nil,
		Count:   nil,
		Search:  to.Ptr("storage"),
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
		// page.TransactionListResult = armbilling.TransactionListResult{
		// 	Value: []*armbilling.Transaction{
		// 		{
		// 			Name: to.Ptr("41000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/transactions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/BillingProfiles/xxxx-xxxx-xxx-xxx/transactions/41000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.TransactionProperties{
		// 				AzureCreditApplied: &armbilling.TransactionPropertiesAzureCreditApplied{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](2000),
		// 				},
		// 				BillingCurrency: to.Ptr("USD"),
		// 				BillingProfileDisplayName: "Contoso operations billing",
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
		// 				ConsumptionCommitmentDecremented: &armbilling.TransactionPropertiesConsumptionCommitmentDecremented{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](100),
		// 				},
		// 				Date: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-01T00:00:00.000Z"); return t}()),
		// 				Discount: to.Ptr[float32](0.1),
		// 				EffectivePrice: &armbilling.TransactionPropertiesEffectivePrice{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](10),
		// 				},
		// 				ExchangeRate: to.Ptr[float32](1),
		// 				Invoice: to.Ptr("G123456789"),
		// 				InvoiceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/invoices/G123456789"),
		// 				InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/22000000-0000-0000-0000-000000000000"),
		// 				MarketPrice: &armbilling.TransactionPropertiesMarketPrice{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](20),
		// 				},
		// 				PartNumber: to.Ptr("0001"),
		// 				PricingCurrency: to.Ptr("USD"),
		// 				ProductDescription: to.Ptr("Standard D1, US West 3"),
		// 				ProductFamily: to.Ptr("Storage"),
		// 				ProductType: to.Ptr("VM Instance"),
		// 				ProductTypeID: to.Ptr("A12345"),
		// 				Quantity: to.Ptr[int32](1),
		// 				ServicePeriodEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-30T00:00:00.000Z"); return t}()),
		// 				ServicePeriodStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-01T00:00:00.000Z"); return t}()),
		// 				SubTotal: &armbilling.TransactionPropertiesSubTotal{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](4500),
		// 				},
		// 				Tax: &armbilling.TransactionPropertiesTax{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](500),
		// 				},
		// 				TransactionAmount: &armbilling.TransactionPropertiesTransactionAmount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](5000),
		// 				},
		// 				TransactionType: to.Ptr("Purchase"),
		// 				UnitOfMeasure: to.Ptr("1 Minute"),
		// 				UnitType: to.Ptr("1 Runtime Minute"),
		// 				Units: to.Ptr[float32](11.25),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("51000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/transactions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/BillingProfiles/xxxx-xxxx-xxx-xxx/transactions/51000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.TransactionProperties{
		// 				AzureCreditApplied: &armbilling.TransactionPropertiesAzureCreditApplied{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](20),
		// 				},
		// 				BillingCurrency: to.Ptr("USD"),
		// 				BillingProfileDisplayName: "Contoso operations billing",
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
		// 				ConsumptionCommitmentDecremented: &armbilling.TransactionPropertiesConsumptionCommitmentDecremented{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](50),
		// 				},
		// 				Date: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-01T00:00:00.000Z"); return t}()),
		// 				Discount: to.Ptr[float32](0.1),
		// 				EffectivePrice: &armbilling.TransactionPropertiesEffectivePrice{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](10),
		// 				},
		// 				ExchangeRate: to.Ptr[float32](1),
		// 				Invoice: to.Ptr("pending"),
		// 				InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/22000000-0000-0000-0000-000000000000"),
		// 				MarketPrice: &armbilling.TransactionPropertiesMarketPrice{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](20),
		// 				},
		// 				PartNumber: to.Ptr("0002"),
		// 				PricingCurrency: to.Ptr("USD"),
		// 				ProductDescription: to.Ptr("Standard Support"),
		// 				ProductFamily: to.Ptr("Storage"),
		// 				ProductType: to.Ptr("VM Instance"),
		// 				ProductTypeID: to.Ptr("A12345"),
		// 				Quantity: to.Ptr[int32](1),
		// 				ServicePeriodEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-30T00:00:00.000Z"); return t}()),
		// 				ServicePeriodStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-01T00:00:00.000Z"); return t}()),
		// 				SubTotal: &armbilling.TransactionPropertiesSubTotal{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](45),
		// 				},
		// 				Tax: &armbilling.TransactionPropertiesTax{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](5),
		// 				},
		// 				TransactionAmount: &armbilling.TransactionPropertiesTransactionAmount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](50),
		// 				},
		// 				TransactionType: to.Ptr("Cancel"),
		// 				UnitOfMeasure: to.Ptr("1 Minute"),
		// 				UnitType: to.Ptr("1 Runtime Minute"),
		// 				Units: to.Ptr[float32](1.25),
		// 			},
		// 	}},
		// }
	}
}
