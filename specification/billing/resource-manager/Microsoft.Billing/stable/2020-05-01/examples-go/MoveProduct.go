package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/MoveProduct.json
func ExampleProductsClient_Move() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProductsClient().Move(ctx, "{billingAccountName}", "{productName}", armbilling.TransferProductRequestProperties{
		DestinationInvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{newInvoiceSectionName}"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Product = armbilling.Product{
	// 	Name: to.Ptr("{productName}"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{newInvoiceSectionName}/products/{productName}"),
	// 	Properties: &armbilling.ProductProperties{
	// 		AutoRenew: to.Ptr(armbilling.AutoRenewOn),
	// 		AvailabilityID: to.Ptr("AvailabilityId1"),
	// 		BillingFrequency: to.Ptr(armbilling.BillingFrequencyMonthly),
	// 		BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
	// 		DisplayName: to.Ptr("Test Product"),
	// 		InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
	// 		InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{newBillingProfileName}/invoiceSections/{newInvoiceSectionName}"),
	// 		LastCharge: &armbilling.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](5000),
	// 		},
	// 		LastChargeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-15T17:32:28.000Z"); return t}()),
	// 		ProductType: to.Ptr("Subscription"),
	// 		ProductTypeID: to.Ptr("A12345"),
	// 		PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T17:32:28.000Z"); return t}()),
	// 		Quantity: to.Ptr[float32](4),
	// 		SKUDescription: to.Ptr("Enterprise Agreement Development"),
	// 		SKUID: to.Ptr("0001"),
	// 		Status: to.Ptr(armbilling.ProductStatusTypeActive),
	// 		TenantID: to.Ptr("515a6d36-aaf8-4ca2-a5e8-c45deb0c5cce"),
	// 	},
	// }
}
