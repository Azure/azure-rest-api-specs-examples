package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutBillingProfile.json
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
	poller, err := clientFactory.NewProfilesClient().BeginCreateOrUpdate(ctx, "{billingAccountName}", "{billingProfileName}", armbilling.Profile{
		Properties: &armbilling.ProfileProperties{
			BillTo: &armbilling.AddressDetails{
				AddressLine1: to.Ptr("Test Address 1"),
				City:         to.Ptr("Redmond"),
				Country:      to.Ptr("US"),
				FirstName:    to.Ptr("Test"),
				LastName:     to.Ptr("User"),
				PostalCode:   to.Ptr("12345"),
				Region:       to.Ptr("WA"),
			},
			DisplayName: to.Ptr("Finance"),
			EnabledAzurePlans: []*armbilling.AzurePlan{
				{
					SKUID: to.Ptr("0001"),
				},
				{
					SKUID: to.Ptr("0002"),
				}},
			InvoiceEmailOptIn: to.Ptr(true),
			PoNumber:          to.Ptr("ABC12345"),
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
	// 	Name: to.Ptr("billingProfileId1"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/billingProfileId1"),
	// 	Properties: &armbilling.ProfileProperties{
	// 		BillTo: &armbilling.AddressDetails{
	// 			AddressLine1: to.Ptr("Test Address 1"),
	// 			City: to.Ptr("Redmond"),
	// 			Country: to.Ptr("US"),
	// 			FirstName: to.Ptr("Test"),
	// 			LastName: to.Ptr("User"),
	// 			PostalCode: to.Ptr("12345"),
	// 			Region: to.Ptr("WA"),
	// 		},
	// 		Currency: to.Ptr("USD"),
	// 		DisplayName: to.Ptr("Finance"),
	// 		EnabledAzurePlans: []*armbilling.AzurePlan{
	// 			{
	// 				SKUDescription: to.Ptr("Microsoft Azure Plan"),
	// 				SKUID: to.Ptr("0001"),
	// 			},
	// 			{
	// 				SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
	// 				SKUID: to.Ptr("0002"),
	// 		}},
	// 		HasReadAccess: to.Ptr(true),
	// 		InvoiceDay: to.Ptr[int32](5),
	// 		InvoiceEmailOptIn: to.Ptr(true),
	// 		PoNumber: to.Ptr("ABC12345"),
	// 		SpendingLimit: to.Ptr(armbilling.SpendingLimitOff),
	// 		Status: to.Ptr(armbilling.BillingProfileStatusActive),
	// 		SystemID: to.Ptr("1XXX-11XX-XX1-XXXX-XXX"),
	// 	},
	// }
}
