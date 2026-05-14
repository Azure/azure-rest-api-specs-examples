package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/productsUpdate.json
func ExampleProductsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProductsClient().Update(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "11111111-1111-1111-1111-111111111111", armbilling.ProductPatch{
		Properties: &armbilling.ProductProperties{
			AutoRenew: to.Ptr(armbilling.AutoRenewOn),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.ProductsClientUpdateResponse{
	// 	Product: armbilling.Product{
	// 		Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/products/11111111-1111-1111-1111-111111111111"),
	// 		Properties: &armbilling.ProductProperties{
	// 			AutoRenew: to.Ptr(armbilling.AutoRenewOn),
	// 			BillingProfileDisplayName: to.Ptr("Billing Profile Display Name"),
	// 			BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
	// 			DisplayName: to.Ptr("Product Display Name"),
	// 			InvoiceSectionDisplayName: to.Ptr("Invoice Section Display Name"),
	// 			InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
	// 			ProductType: to.Ptr("Seat-Based Product Type"),
	// 			ProductTypeID: to.Ptr("XYZ56789"),
	// 			PurchaseDate: to.Ptr("2023-01-05T22:39:34.2606750Z"),
	// 			Quantity: to.Ptr[int64](1),
	// 			SKUDescription: to.Ptr("SKU Description"),
	// 			SKUID: to.Ptr("0001"),
	// 			Status: to.Ptr(armbilling.ProductStatusActive),
	// 		},
	// 	},
	// }
}
