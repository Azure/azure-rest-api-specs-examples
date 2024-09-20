package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsGetByBillingProfile.json
func ExamplePaymentMethodsClient_GetByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPaymentMethodsClient().GetByBillingProfile(ctx, "00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31", "ABC1-A1CD-AB1-BP1", "ABCDABCDABC0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PaymentMethodLink = armbilling.PaymentMethodLink{
	// 	Name: to.Ptr("ABCDABCDABC0"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/paymentMethodLinks"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31/billingProfiles/ABC1-A1CD-AB1-BP1/paymentMethodLinks/ABCDABCDABC0"),
	// 	Properties: &armbilling.PaymentMethodLinkProperties{
	// 		AccountHolderName: to.Ptr("abc"),
	// 		DisplayName: to.Ptr("Master Card"),
	// 		Expiration: to.Ptr("1/2035"),
	// 		Family: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
	// 		LastFourDigits: to.Ptr("1270"),
	// 		Logos: []*armbilling.PaymentMethodLogo{
	// 			{
	// 				MimeType: to.Ptr("image/png"),
	// 				URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa_rect.png"),
	// 			},
	// 			{
	// 				MimeType: to.Ptr("image/svg+xml"),
	// 				URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa.svg"),
	// 		}},
	// 		PaymentMethodID: to.Ptr("/providers/Microsoft.Billing/paymentMethods/ABCDABCDABC0"),
	// 		PaymentMethodType: to.Ptr("mc"),
	// 		Status: to.Ptr(armbilling.PaymentMethodStatusActive),
	// 	},
	// }
}
