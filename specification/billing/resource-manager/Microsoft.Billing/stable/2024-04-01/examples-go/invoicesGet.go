package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoicesGet.json
func ExampleInvoicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInvoicesClient().Get(ctx, "G123456789", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Invoice = armbilling.Invoice{
	// 	Name: to.Ptr("G123456789"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/invoices"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/invoices/G123456789"),
	// 	Properties: &armbilling.InvoiceProperties{
	// 		AmountDue: &armbilling.InvoicePropertiesAmountDue{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](8.53),
	// 		},
	// 		AzurePrepaymentApplied: &armbilling.InvoicePropertiesAzurePrepaymentApplied{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](25.46),
	// 		},
	// 		BilledAmount: &armbilling.InvoicePropertiesBilledAmount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](33.99),
	// 		},
	// 		BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
	// 		CreditAmount: &armbilling.InvoicePropertiesCreditAmount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](0),
	// 		},
	// 		DocumentType: to.Ptr(armbilling.InvoiceDocumentTypeInvoice),
	// 		Documents: []*armbilling.InvoiceDocument{
	// 			{
	// 				Name: to.Ptr("12345678"),
	// 				Kind: to.Ptr(armbilling.InvoiceDocumentTypeInvoice),
	// 			},
	// 			{
	// 				Name: to.Ptr("22345678"),
	// 				Kind: to.Ptr(armbilling.InvoiceDocumentTypeTaxReceipt),
	// 		}},
	// 		DueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-16T17:32:28.000Z"); return t}()),
	// 		FreeAzureCreditApplied: &armbilling.InvoicePropertiesFreeAzureCreditApplied{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](0),
	// 		},
	// 		InvoiceDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T17:32:28.000Z"); return t}()),
	// 		InvoicePeriodEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-15T17:32:28.000Z"); return t}()),
	// 		InvoicePeriodStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T17:32:28.000Z"); return t}()),
	// 		IsMonthlyInvoice: to.Ptr(false),
	// 		PurchaseOrderNumber: to.Ptr("123456"),
	// 		RebillDetails: &armbilling.InvoicePropertiesRebillDetails{
	// 			CreditNoteDocumentID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/invoices/CreditNote2"),
	// 			InvoiceDocumentID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/invoices/I000002"),
	// 			RebillDetails: &armbilling.RebillDetails{
	// 				CreditNoteDocumentID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/invoices/CreditNote"),
	// 				InvoiceDocumentID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/invoices/I000001"),
	// 			},
	// 		},
	// 		SpecialTaxationType: to.Ptr(armbilling.SpecialTaxationTypeSubtotalLevel),
	// 		Status: to.Ptr(armbilling.InvoiceStatusDue),
	// 		SubTotal: &armbilling.InvoicePropertiesSubTotal{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](33.99),
	// 		},
	// 		TaxAmount: &armbilling.InvoicePropertiesTaxAmount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](1),
	// 		},
	// 		TotalAmount: &armbilling.InvoicePropertiesTotalAmount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](7.53),
	// 		},
	// 	},
	// }
}
