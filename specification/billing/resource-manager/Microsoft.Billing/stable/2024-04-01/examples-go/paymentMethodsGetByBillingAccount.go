package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsGetByBillingAccount.json
func ExamplePaymentMethodsClient_GetByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPaymentMethodsClient().GetByBillingAccount(ctx, "00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31", "21dd9edc-af71-4d62-80ce-37151d475326", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PaymentMethod = armbilling.PaymentMethod{
	// 	Name: to.Ptr("21dd9edc-af71-4d62-80ce-37151d475326"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/paymentMethods"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31/paymentMethods/21dd9edc-af71-4d62-80ce-37151d475326"),
	// 	Properties: &armbilling.PaymentMethodProperties{
	// 		DisplayName: to.Ptr("Check/wire transfer"),
	// 		Family: to.Ptr(armbilling.PaymentMethodFamilyCheckWire),
	// 		PaymentMethodType: to.Ptr("check"),
	// 	},
	// }
}
