package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountInvoicesList.json
func ExampleInvoicesClient_NewListByBillingAccountPager_billingAccountInvoicesList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInvoicesClient().NewListByBillingAccountPager("{billingAccountName}", "2018-01-01", "2018-06-30", nil)
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
		// page.InvoiceListResult = armbilling.InvoiceListResult{
		// 	Value: []*armbilling.Invoice{
		// 		{
		// 			Name: to.Ptr("1383723"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/invoices"),
		// 			ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/1383723"),
		// 			Properties: &armbilling.InvoiceProperties{
		// 				AmountDue: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](8.53),
		// 				},
		// 				AzurePrepaymentApplied: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](25.46),
		// 				},
		// 				BilledAmount: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](33.99),
		// 				},
		// 				BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 				CreditAmount: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](1),
		// 				},
		// 				DocumentType: to.Ptr(armbilling.InvoiceDocumentTypeInvoice),
		// 				Documents: []*armbilling.Document{
		// 					{
		// 						Kind: to.Ptr(armbilling.DocumentTypeInvoice),
		// 						Source: to.Ptr(armbilling.DocumentSourceDRS),
		// 						URL: to.Ptr("https://microsoft.com/invoice.pdf"),
		// 				}},
		// 				DueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-16T17:32:28.000Z"); return t}()),
		// 				FreeAzureCreditApplied: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](0),
		// 				},
		// 				InvoiceDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T17:32:28.000Z"); return t}()),
		// 				InvoicePeriodEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-15T17:32:28.000Z"); return t}()),
		// 				InvoicePeriodStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T17:32:28.000Z"); return t}()),
		// 				IsMonthlyInvoice: to.Ptr(false),
		// 				Payments: []*armbilling.PaymentProperties{
		// 					{
		// 						Amount: &armbilling.Amount{
		// 							Currency: to.Ptr("USD"),
		// 							Value: to.Ptr[float32](1),
		// 						},
		// 						Date: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-14T17:32:28.000Z"); return t}()),
		// 						PaymentMethodFamily: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
		// 						PaymentMethodType: to.Ptr("visa"),
		// 						PaymentType: to.Ptr("credited"),
		// 				}},
		// 				PurchaseOrderNumber: to.Ptr("123456"),
		// 				Status: to.Ptr(armbilling.InvoiceStatusDue),
		// 				SubTotal: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](33.99),
		// 				},
		// 				TaxAmount: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](0),
		// 				},
		// 				TotalAmount: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](7.53),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("1383724"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/invoices"),
		// 			ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/1383724"),
		// 			Properties: &armbilling.InvoiceProperties{
		// 				AmountDue: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](16.53),
		// 				},
		// 				AzurePrepaymentApplied: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](15.46),
		// 				},
		// 				BilledAmount: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](33.99),
		// 				},
		// 				BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 				CreditAmount: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](2),
		// 				},
		// 				DocumentType: to.Ptr(armbilling.InvoiceDocumentTypeInvoice),
		// 				Documents: []*armbilling.Document{
		// 					{
		// 						Kind: to.Ptr(armbilling.DocumentTypeTaxReceipt),
		// 						Source: to.Ptr(armbilling.DocumentSourceDRS),
		// 						URL: to.Ptr("https://microsoft.com/taxreceipt.pdf"),
		// 				}},
		// 				DueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-01T17:32:28.000Z"); return t}()),
		// 				FreeAzureCreditApplied: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](0),
		// 				},
		// 				InvoiceDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-01T17:32:28.000Z"); return t}()),
		// 				InvoicePeriodEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-28T17:32:28.000Z"); return t}()),
		// 				InvoicePeriodStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-01T17:32:28.000Z"); return t}()),
		// 				IsMonthlyInvoice: to.Ptr(true),
		// 				Payments: []*armbilling.PaymentProperties{
		// 					{
		// 						Amount: &armbilling.Amount{
		// 							Currency: to.Ptr("USD"),
		// 							Value: to.Ptr[float32](2),
		// 						},
		// 						Date: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-14T17:32:28.000Z"); return t}()),
		// 						PaymentMethodFamily: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
		// 						PaymentMethodType: to.Ptr("visa"),
		// 						PaymentType: to.Ptr("credited"),
		// 				}},
		// 				PurchaseOrderNumber: to.Ptr("123456"),
		// 				Status: to.Ptr(armbilling.InvoiceStatus("PastDue")),
		// 				SubTotal: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](33.99),
		// 				},
		// 				TaxAmount: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](0),
		// 				},
		// 				TotalAmount: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](16.53),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
