package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CreditNote.json
func ExampleInvoicesClient_Get_creditNote() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInvoicesClient().Get(ctx, "{billingAccountName}", "{invoiceName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Invoice = armbilling.Invoice{
	// 	Name: to.Ptr("{invoiceName}"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/invoices"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/{invoiceName}"),
	// 	Properties: &armbilling.InvoiceProperties{
	// 		AmountDue: &armbilling.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](16.53),
	// 		},
	// 		AzurePrepaymentApplied: &armbilling.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](15.46),
	// 		},
	// 		BilledAmount: &armbilling.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](33.99),
	// 		},
	// 		BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}"),
	// 		CreditAmount: &armbilling.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](2),
	// 		},
	// 		CreditForDocumentID: to.Ptr("providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/I000001"),
	// 		DocumentType: to.Ptr(armbilling.InvoiceDocumentTypeCreditNote),
	// 		Documents: []*armbilling.Document{
	// 			{
	// 				Kind: to.Ptr(armbilling.DocumentTypeInvoice),
	// 				Source: to.Ptr(armbilling.DocumentSourceDRS),
	// 				URL: to.Ptr("https://microsoft.com/invoice.pdf"),
	// 			},
	// 			{
	// 				Kind: to.Ptr(armbilling.DocumentTypeTaxReceipt),
	// 				Source: to.Ptr(armbilling.DocumentSourceDRS),
	// 				URL: to.Ptr("https://microsoft.com/taxDoc.pdf"),
	// 		}},
	// 		DueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-16T17:32:28.000Z"); return t}()),
	// 		FreeAzureCreditApplied: &armbilling.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](0),
	// 		},
	// 		InvoiceDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T17:32:28.000Z"); return t}()),
	// 		InvoicePeriodEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-15T17:32:28.000Z"); return t}()),
	// 		InvoicePeriodStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T17:32:28.000Z"); return t}()),
	// 		IsMonthlyInvoice: to.Ptr(false),
	// 		Payments: []*armbilling.PaymentProperties{
	// 			{
	// 				Amount: &armbilling.Amount{
	// 					Currency: to.Ptr("USD"),
	// 					Value: to.Ptr[float32](2),
	// 				},
	// 				Date: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-14T17:32:28.000Z"); return t}()),
	// 				PaymentMethodFamily: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
	// 				PaymentMethodType: to.Ptr("visa"),
	// 				PaymentType: to.Ptr("credited"),
	// 		}},
	// 		PurchaseOrderNumber: to.Ptr("123456"),
	// 		Status: to.Ptr(armbilling.InvoiceStatusPaid),
	// 		SubTotal: &armbilling.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](33.99),
	// 		},
	// 		TaxAmount: &armbilling.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](0),
	// 		},
	// 		TotalAmount: &armbilling.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](16.53),
	// 		},
	// 	},
	// }
}
