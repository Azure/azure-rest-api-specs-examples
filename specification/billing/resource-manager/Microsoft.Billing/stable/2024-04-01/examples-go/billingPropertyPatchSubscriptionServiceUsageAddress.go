package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPropertyPatchSubscriptionServiceUsageAddress.json
func ExamplePropertyClient_Update_billingPropertyPatchSubscriptionServiceUsageAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPropertyClient().Update(ctx, armbilling.Property{
		Properties: &armbilling.PropertyProperties{
			SubscriptionServiceUsageAddress: &armbilling.PropertyPropertiesSubscriptionServiceUsageAddress{
				AddressLine1: to.Ptr("Address line 1"),
				AddressLine2: to.Ptr("Address line 2"),
				City:         to.Ptr("City"),
				Country:      to.Ptr("US"),
				FirstName:    to.Ptr("Jenny"),
				LastName:     to.Ptr("Doe"),
				MiddleName:   to.Ptr("Ann"),
				PostalCode:   to.Ptr("12345"),
				Region:       to.Ptr("State"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Property = armbilling.Property{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Billing/billingProperty"),
	// 	ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/providers/Microsoft.Billing/billingProperty/default"),
	// 	Properties: &armbilling.PropertyProperties{
	// 		BillingAccountAgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftOnlineServicesProgram),
	// 		BillingAccountDisplayName: to.Ptr("My Account"),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000"),
	// 		BillingAccountSoldToCountry: to.Ptr("US"),
	// 		BillingAccountType: to.Ptr(armbilling.AccountTypeIndividual),
	// 		SubscriptionBillingType: to.Ptr(armbilling.SubscriptionBillingTypeFree),
	// 		SubscriptionServiceUsageAddress: &armbilling.PropertyPropertiesSubscriptionServiceUsageAddress{
	// 			AddressLine1: to.Ptr("Address line 1"),
	// 			AddressLine2: to.Ptr("Address line 2"),
	// 			City: to.Ptr("City"),
	// 			Country: to.Ptr("US"),
	// 			FirstName: to.Ptr("Jenny"),
	// 			LastName: to.Ptr("Doe"),
	// 			MiddleName: to.Ptr("Ann"),
	// 			PostalCode: to.Ptr("12345"),
	// 			Region: to.Ptr("State"),
	// 		},
	// 		SubscriptionWorkloadType: to.Ptr(armbilling.SubscriptionWorkloadTypeProduction),
	// 	},
	// }
}
