package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsListByUser.json
func ExamplePaymentMethodsClient_NewListByUserPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPaymentMethodsClient().NewListByUserPager(nil)
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
		// page.PaymentMethodsListResult = armbilling.PaymentMethodsListResult{
		// 	Value: []*armbilling.PaymentMethod{
		// 		{
		// 			Name: to.Ptr("ABCDABCDABC0"),
		// 			Type: to.Ptr("Microsoft.Billing/paymentMethods"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/paymentMethods/ABCDABCDABC0"),
		// 			Properties: &armbilling.PaymentMethodProperties{
		// 				AccountHolderName: to.Ptr("abc"),
		// 				DisplayName: to.Ptr("Master Card"),
		// 				Expiration: to.Ptr("1/2035"),
		// 				Family: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
		// 				LastFourDigits: to.Ptr("1270"),
		// 				Logos: []*armbilling.PaymentMethodLogo{
		// 					{
		// 						MimeType: to.Ptr("image/png"),
		// 						URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_mc_rect.png"),
		// 					},
		// 					{
		// 						MimeType: to.Ptr("image/svg+xml"),
		// 						URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_mc.svg"),
		// 				}},
		// 				PaymentMethodType: to.Ptr("mc"),
		// 				Status: to.Ptr(armbilling.PaymentMethodStatusActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ABCDABCDABC1"),
		// 			Type: to.Ptr("Microsoft.Billing/paymentMethods"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/paymentMethods/ABCDABCDABC1"),
		// 			Properties: &armbilling.PaymentMethodProperties{
		// 				AccountHolderName: to.Ptr("abc"),
		// 				DisplayName: to.Ptr("Visa"),
		// 				Expiration: to.Ptr("1/2025"),
		// 				Family: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
		// 				LastFourDigits: to.Ptr("7373"),
		// 				Logos: []*armbilling.PaymentMethodLogo{
		// 					{
		// 						MimeType: to.Ptr("image/png"),
		// 						URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa_rect.png"),
		// 					},
		// 					{
		// 						MimeType: to.Ptr("image/svg+xml"),
		// 						URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa.svg"),
		// 				}},
		// 				PaymentMethodType: to.Ptr("visa"),
		// 				Status: to.Ptr(armbilling.PaymentMethodStatusActive),
		// 			},
		// 	}},
		// }
	}
}
