package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoicesGetByBillingSubscription.json
func ExampleInvoicesClient_GetByBillingSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInvoicesClient().GetByBillingSubscription(ctx, "E123456789", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Invoice = armbilling.Invoice{
	// 	Name: to.Ptr("E123456789"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptions/invoices"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/11111111-1111-1111-1111-111111111111/invoices/E123456789"),
	// 	Properties: &armbilling.InvoiceProperties{
	// 		AmountDue: &armbilling.InvoicePropertiesAmountDue{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](8.53),
	// 		},
	// 		BilledAmount: &armbilling.InvoicePropertiesBilledAmount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](33.99),
	// 		},
	// 		DueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-16T17:32:28.000Z"); return t}()),
	// 		InvoiceDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T17:32:28.000Z"); return t}()),
	// 		InvoicePeriodEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-15T17:32:28.000Z"); return t}()),
	// 		InvoicePeriodStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T17:32:28.000Z"); return t}()),
	// 		InvoiceType: to.Ptr(armbilling.InvoiceTypeAzureServices),
	// 		PurchaseOrderNumber: to.Ptr("123456"),
	// 		Status: to.Ptr(armbilling.InvoiceStatusDue),
	// 		SubscriptionDisplayName: to.Ptr("Contoso Operations Billing"),
	// 		SubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 	},
	// }
}
