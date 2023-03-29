package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfilesListByBillingAccount.json
func ExampleProfilesClient_NewListByBillingAccountPager_billingProfilesListByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProfilesClient().NewListByBillingAccountPager("{billingAccountName}", &armbilling.ProfilesClientListByBillingAccountOptions{Expand: nil})
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
		// page.ProfileListResult = armbilling.ProfileListResult{
		// 	Value: []*armbilling.Profile{
		// 		{
		// 			Name: to.Ptr("11000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.ProfileProperties{
		// 				BillTo: &armbilling.AddressDetails{
		// 					AddressLine1: to.Ptr("Test Address1"),
		// 					AddressLine2: to.Ptr("Test Address2"),
		// 					AddressLine3: to.Ptr("Test Address3"),
		// 					City: to.Ptr("City"),
		// 					CompanyName: to.Ptr("Contoso"),
		// 					Country: to.Ptr("US"),
		// 					Email: to.Ptr("abc@contoso.com"),
		// 					FirstName: to.Ptr("Test"),
		// 					LastName: to.Ptr("User"),
		// 					PhoneNumber: to.Ptr("000-000-0000"),
		// 					PostalCode: to.Ptr("00000"),
		// 					Region: to.Ptr("WA"),
		// 				},
		// 				BillingRelationshipType: to.Ptr(armbilling.BillingRelationshipTypeDirect),
		// 				Currency: to.Ptr("USD"),
		// 				DisplayName: to.Ptr("BillingProfile1"),
		// 				EnabledAzurePlans: []*armbilling.AzurePlan{
		// 					{
		// 						SKUDescription: to.Ptr("Microsoft Azure Plan"),
		// 						SKUID: to.Ptr("0001"),
		// 					},
		// 					{
		// 						SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
		// 						SKUID: to.Ptr("0002"),
		// 				}},
		// 				HasReadAccess: to.Ptr(true),
		// 				InvoiceDay: to.Ptr[int32](5),
		// 				InvoiceEmailOptIn: to.Ptr(true),
		// 				PoNumber: to.Ptr("ABC12345"),
		// 				SpendingLimit: to.Ptr(armbilling.SpendingLimitOn),
		// 				Status: to.Ptr(armbilling.BillingProfileStatusWarned),
		// 				StatusReasonCode: to.Ptr(armbilling.StatusReasonCodePastDue),
		// 				SystemID: to.Ptr("1XXX-11XX-XX1-XXXX-XXX"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("11000000-0000-0000-0000-000000000001"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000001"),
		// 			Properties: &armbilling.ProfileProperties{
		// 				BillTo: &armbilling.AddressDetails{
		// 					AddressLine1: to.Ptr("Test Address1"),
		// 					AddressLine2: to.Ptr("Test Address2"),
		// 					AddressLine3: to.Ptr("Test Address3"),
		// 					City: to.Ptr("City"),
		// 					CompanyName: to.Ptr("Contoso"),
		// 					Country: to.Ptr("US"),
		// 					Email: to.Ptr("abc@contoso.com"),
		// 					FirstName: to.Ptr("Test"),
		// 					LastName: to.Ptr("User"),
		// 					PhoneNumber: to.Ptr("000-000-0000"),
		// 					PostalCode: to.Ptr("00000"),
		// 					Region: to.Ptr("WA"),
		// 				},
		// 				BillingRelationshipType: to.Ptr(armbilling.BillingRelationshipTypeIndirectCustomer),
		// 				Currency: to.Ptr("USD"),
		// 				DisplayName: to.Ptr("BillingProfile2"),
		// 				EnabledAzurePlans: []*armbilling.AzurePlan{
		// 					{
		// 						SKUDescription: to.Ptr("Microsoft Azure Plan"),
		// 						SKUID: to.Ptr("0001"),
		// 					},
		// 					{
		// 						SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
		// 						SKUID: to.Ptr("0002"),
		// 				}},
		// 				HasReadAccess: to.Ptr(true),
		// 				IndirectRelationshipInfo: &armbilling.IndirectRelationshipInfo{
		// 					BillingAccountName: to.Ptr("30000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000"),
		// 					BillingProfileName: to.Ptr("33000000-0000-0000-0000-000000000001"),
		// 					DisplayName: to.Ptr("Partner1"),
		// 				},
		// 				InvoiceDay: to.Ptr[int32](5),
		// 				InvoiceEmailOptIn: to.Ptr(true),
		// 				PoNumber: to.Ptr("ABC12345"),
		// 				SpendingLimit: to.Ptr(armbilling.SpendingLimitOff),
		// 				Status: to.Ptr(armbilling.BillingProfileStatusActive),
		// 				SystemID: to.Ptr("2XXX-22XX-XX1-XXXX-XXX"),
		// 				TargetClouds: []*armbilling.TargetCloud{
		// 				},
		// 			},
		// 	}},
		// }
	}
}
