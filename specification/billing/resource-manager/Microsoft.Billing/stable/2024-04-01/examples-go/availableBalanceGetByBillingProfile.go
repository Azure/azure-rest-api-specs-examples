package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/availableBalanceGetByBillingProfile.json
func ExampleAvailableBalancesClient_GetByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailableBalancesClient().GetByBillingProfile(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvailableBalance = armbilling.AvailableBalance{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/availableBalance"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/availableBalance/default"),
	// 	Properties: &armbilling.AvailableBalanceProperties{
	// 		Amount: &armbilling.AvailableBalancePropertiesAmount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](500),
	// 		},
	// 		PaymentsOnAccount: []*armbilling.PaymentOnAccount{
	// 			{
	// 				Amount: &armbilling.PaymentOnAccountAmount{
	// 					Currency: to.Ptr("USD"),
	// 					Value: to.Ptr[float32](50),
	// 				},
	// 				BillingProfileDisplayName: to.Ptr("Contoso Finance"),
	// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
	// 				Date: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-01T00:00:00.000Z"); return t}()),
	// 				InvoiceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/invoices/G123456789"),
	// 				InvoiceName: to.Ptr("G123456789"),
	// 				PaymentMethodType: to.Ptr(armbilling.PaymentMethodFamilyCheckWire),
	// 		}},
	// 		TotalPaymentsOnAccount: &armbilling.AvailableBalancePropertiesTotalPaymentsOnAccount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](200),
	// 		},
	// 	},
	// }
}
