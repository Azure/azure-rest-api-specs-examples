package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/paymentMethodsGetByUser.json
func ExamplePaymentMethodsClient_GetByUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPaymentMethodsClient().GetByUser(ctx, "ABCDABCDABC0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.PaymentMethodsClientGetByUserResponse{
	// 	PaymentMethod: armbilling.PaymentMethod{
	// 		Name: to.Ptr("ABCDABCDABC0"),
	// 		Type: to.Ptr("Microsoft.Billing/paymentMethods"),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/paymentMethods/ABCDABCDABC0"),
	// 		Properties: &armbilling.PaymentMethodProperties{
	// 			AccountHolderName: to.Ptr("abc"),
	// 			DisplayName: to.Ptr("Master Card"),
	// 			Expiration: to.Ptr("1/2035"),
	// 			Family: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
	// 			LastFourDigits: to.Ptr("1270"),
	// 			Logos: []*armbilling.PaymentMethodLogo{
	// 				{
	// 					MimeType: to.Ptr("image/png"),
	// 					URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa_rect.png"),
	// 				},
	// 				{
	// 					MimeType: to.Ptr("image/svg+xml"),
	// 					URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa.svg"),
	// 				},
	// 			},
	// 			PaymentMethodType: to.Ptr("mc"),
	// 			Status: to.Ptr(armbilling.PaymentMethodStatusActive),
	// 		},
	// 	},
	// }
}
