package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingProfilesCreateOrUpdate.json
func ExampleProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProfilesClient().BeginCreateOrUpdate(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", armbilling.Profile{
		Properties: &armbilling.ProfileProperties{
			BillTo: &armbilling.ProfilePropertiesBillTo{
				AddressLine1:   to.Ptr("Test Address1"),
				AddressLine2:   to.Ptr("Test Address2"),
				AddressLine3:   to.Ptr("Test Address3"),
				City:           to.Ptr("City"),
				CompanyName:    to.Ptr("Contoso"),
				Country:        to.Ptr("US"),
				Email:          to.Ptr("abc@contoso.com"),
				FirstName:      to.Ptr("Test"),
				IsValidAddress: to.Ptr(true),
				LastName:       to.Ptr("User"),
				PhoneNumber:    to.Ptr("000-000-0000"),
				PostalCode:     to.Ptr("00000"),
				Region:         to.Ptr("WA"),
			},
			DisplayName: to.Ptr("Billing Profile 1"),
			EnabledAzurePlans: []*armbilling.AzurePlan{
				{
					SKUID: to.Ptr("0001"),
				},
				{
					SKUID: to.Ptr("0002"),
				}},
			InvoiceEmailOptIn: to.Ptr(true),
			PoNumber:          to.Ptr("ABC12345"),
			ShipTo: &armbilling.ProfilePropertiesShipTo{
				AddressLine1:   to.Ptr("Test Address1"),
				AddressLine2:   to.Ptr("Test Address2"),
				AddressLine3:   to.Ptr("Test Address3"),
				City:           to.Ptr("City"),
				CompanyName:    to.Ptr("Contoso"),
				Country:        to.Ptr("US"),
				Email:          to.Ptr("abc@contoso.com"),
				FirstName:      to.Ptr("Test"),
				IsValidAddress: to.Ptr(true),
				LastName:       to.Ptr("User"),
				PhoneNumber:    to.Ptr("000-000-0000"),
				PostalCode:     to.Ptr("00000"),
				Region:         to.Ptr("WA"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Profile = armbilling.Profile{
	// 	Name: to.Ptr("xxxx-xxxx-xxx-xxx"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
	// 	Properties: &armbilling.ProfileProperties{
	// 		BillTo: &armbilling.ProfilePropertiesBillTo{
	// 			AddressLine1: to.Ptr("Test Address1"),
	// 			AddressLine2: to.Ptr("Test Address2"),
	// 			AddressLine3: to.Ptr("Test Address3"),
	// 			City: to.Ptr("City"),
	// 			CompanyName: to.Ptr("Contoso"),
	// 			Country: to.Ptr("US"),
	// 			Email: to.Ptr("abc@contoso.com"),
	// 			FirstName: to.Ptr("Test"),
	// 			IsValidAddress: to.Ptr(true),
	// 			LastName: to.Ptr("User"),
	// 			PhoneNumber: to.Ptr("000-000-0000"),
	// 			PostalCode: to.Ptr("00000"),
	// 			Region: to.Ptr("WA"),
	// 		},
	// 		Currency: to.Ptr("USD"),
	// 		DisplayName: to.Ptr("Billing Profile 1"),
	// 		HasReadAccess: to.Ptr(true),
	// 		InvoiceDay: to.Ptr[int32](5),
	// 		InvoiceEmailOptIn: to.Ptr(true),
	// 		PoNumber: to.Ptr("ABC12345"),
	// 		ShipTo: &armbilling.ProfilePropertiesShipTo{
	// 			AddressLine1: to.Ptr("Test Address1"),
	// 			AddressLine2: to.Ptr("Test Address2"),
	// 			AddressLine3: to.Ptr("Test Address3"),
	// 			City: to.Ptr("City"),
	// 			CompanyName: to.Ptr("Contoso"),
	// 			Country: to.Ptr("US"),
	// 			Email: to.Ptr("abc@contoso.com"),
	// 			FirstName: to.Ptr("Test"),
	// 			IsValidAddress: to.Ptr(true),
	// 			LastName: to.Ptr("User"),
	// 			PhoneNumber: to.Ptr("000-000-0000"),
	// 			PostalCode: to.Ptr("00000"),
	// 			Region: to.Ptr("WA"),
	// 		},
	// 		Status: to.Ptr(armbilling.BillingProfileStatusActive),
	// 		SystemID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 	},
	// }
}
