package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountWithExpand.json
func ExampleAccountsClient_Get_billingAccountWithExpand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Get(ctx, "{billingAccountName}", &armbilling.AccountsClientGetOptions{Expand: to.Ptr("soldTo,billingProfiles,billingProfiles/invoiceSections")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armbilling.Account{
	// 	Name: to.Ptr("{billingAccountName}"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}"),
	// 	Properties: &armbilling.AccountProperties{
	// 		AccountStatus: to.Ptr(armbilling.AccountStatusActive),
	// 		AccountType: to.Ptr(armbilling.AccountTypeEnterprise),
	// 		AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
	// 		BillingProfiles: &armbilling.ProfilesOnExpand{
	// 			HasMoreResults: to.Ptr(true),
	// 			Value: []*armbilling.Profile{
	// 				{
	// 					Name: to.Ptr("11000000-0000-0000-0000-000000000000"),
	// 					Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles"),
	// 					ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
	// 					Properties: &armbilling.ProfileProperties{
	// 						BillTo: &armbilling.AddressDetails{
	// 							AddressLine1: to.Ptr("Test Address1"),
	// 							AddressLine2: to.Ptr("Test Address2"),
	// 							AddressLine3: to.Ptr("Test Address3"),
	// 							City: to.Ptr("City"),
	// 							CompanyName: to.Ptr("Contoso"),
	// 							Country: to.Ptr("US"),
	// 							Email: to.Ptr("abc@contoso.com"),
	// 							FirstName: to.Ptr("Test"),
	// 							LastName: to.Ptr("User"),
	// 							PhoneNumber: to.Ptr("000-000-0000"),
	// 							PostalCode: to.Ptr("00000"),
	// 							Region: to.Ptr("WA"),
	// 						},
	// 						BillingRelationshipType: to.Ptr(armbilling.BillingRelationshipTypeDirect),
	// 						Currency: to.Ptr("USD"),
	// 						DisplayName: to.Ptr("BillingProfile1"),
	// 						EnabledAzurePlans: []*armbilling.AzurePlan{
	// 							{
	// 								SKUDescription: to.Ptr("Microsoft Azure Plan"),
	// 								SKUID: to.Ptr("0001"),
	// 							},
	// 							{
	// 								SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
	// 								SKUID: to.Ptr("0002"),
	// 						}},
	// 						HasReadAccess: to.Ptr(true),
	// 						InvoiceDay: to.Ptr[int32](5),
	// 						InvoiceEmailOptIn: to.Ptr(true),
	// 						InvoiceSections: &armbilling.InvoiceSectionsOnExpand{
	// 							HasMoreResults: to.Ptr(true),
	// 							Value: []*armbilling.InvoiceSection{
	// 								{
	// 									Name: to.Ptr("invoiceSectionId1"),
	// 									Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
	// 									ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000/invoiceSections/invoiceSectionId1"),
	// 									Properties: &armbilling.InvoiceSectionProperties{
	// 										DisplayName: to.Ptr("invoiceSectionName1"),
	// 										Labels: map[string]*string{
	// 											"costCategory": to.Ptr("Support"),
	// 											"pcCode": to.Ptr("A123456"),
	// 										},
	// 										State: to.Ptr(armbilling.InvoiceSectionStateActive),
	// 										SystemID: to.Ptr("9XXX-11XX-XX1-XXXX-XXX"),
	// 									},
	// 							}},
	// 						},
	// 						PoNumber: to.Ptr("ABC12345"),
	// 						SpendingLimit: to.Ptr(armbilling.SpendingLimitOn),
	// 						Status: to.Ptr(armbilling.BillingProfileStatusWarned),
	// 						StatusReasonCode: to.Ptr(armbilling.StatusReasonCodePastDue),
	// 						SystemID: to.Ptr("1XXX-11XX-XX1-XXXX-XXX"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("11000000-0000-0000-0000-000000000001"),
	// 					Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles"),
	// 					ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000001"),
	// 					Properties: &armbilling.ProfileProperties{
	// 						BillTo: &armbilling.AddressDetails{
	// 							AddressLine1: to.Ptr("Test Address1"),
	// 							AddressLine2: to.Ptr("Test Address2"),
	// 							AddressLine3: to.Ptr("Test Address3"),
	// 							City: to.Ptr("City"),
	// 							CompanyName: to.Ptr("Contoso"),
	// 							Country: to.Ptr("US"),
	// 							Email: to.Ptr("abc@contoso.com"),
	// 							FirstName: to.Ptr("Test"),
	// 							LastName: to.Ptr("User"),
	// 							PhoneNumber: to.Ptr("000-000-0000"),
	// 							PostalCode: to.Ptr("00000"),
	// 							Region: to.Ptr("WA"),
	// 						},
	// 						BillingRelationshipType: to.Ptr(armbilling.BillingRelationshipTypeIndirectCustomer),
	// 						Currency: to.Ptr("USD"),
	// 						DisplayName: to.Ptr("BillingProfile2"),
	// 						EnabledAzurePlans: []*armbilling.AzurePlan{
	// 							{
	// 								SKUDescription: to.Ptr("Microsoft Azure Plan"),
	// 								SKUID: to.Ptr("0001"),
	// 							},
	// 							{
	// 								SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
	// 								SKUID: to.Ptr("0002"),
	// 						}},
	// 						HasReadAccess: to.Ptr(true),
	// 						IndirectRelationshipInfo: &armbilling.IndirectRelationshipInfo{
	// 							BillingAccountName: to.Ptr("30000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000"),
	// 							BillingProfileName: to.Ptr("33000000-0000-0000-0000-000000000001"),
	// 							DisplayName: to.Ptr("Partner1"),
	// 						},
	// 						InvoiceDay: to.Ptr[int32](5),
	// 						InvoiceEmailOptIn: to.Ptr(true),
	// 						InvoiceSections: &armbilling.InvoiceSectionsOnExpand{
	// 							HasMoreResults: to.Ptr(true),
	// 							Value: []*armbilling.InvoiceSection{
	// 								{
	// 									Name: to.Ptr("invoiceSectionId2"),
	// 									Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
	// 									ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000001/invoiceSections/invoiceSectionId2"),
	// 									Properties: &armbilling.InvoiceSectionProperties{
	// 										DisplayName: to.Ptr("invoiceSectionName2"),
	// 										Labels: map[string]*string{
	// 											"costCategory": to.Ptr("Marketing"),
	// 											"pcCode": to.Ptr("Z223456"),
	// 										},
	// 										State: to.Ptr(armbilling.InvoiceSectionStateActive),
	// 										SystemID: to.Ptr("9XXX-22XX-XX1-XXXX-XXX"),
	// 									},
	// 							}},
	// 						},
	// 						PoNumber: to.Ptr("ABC12345"),
	// 						SpendingLimit: to.Ptr(armbilling.SpendingLimitOff),
	// 						Status: to.Ptr(armbilling.BillingProfileStatusActive),
	// 						SystemID: to.Ptr("2XXX-22XX-XX1-XXXX-XXX"),
	// 					},
	// 			}},
	// 		},
	// 		DisplayName: to.Ptr("Test Account 1"),
	// 		HasReadAccess: to.Ptr(true),
	// 		SoldTo: &armbilling.AddressDetails{
	// 			AddressLine1: to.Ptr("Test Address"),
	// 			AddressLine2: to.Ptr("Test Address"),
	// 			AddressLine3: to.Ptr("Test Address"),
	// 			City: to.Ptr("City"),
	// 			CompanyName: to.Ptr("Contoso"),
	// 			Country: to.Ptr("US"),
	// 			Email: to.Ptr("abc@contoso.com"),
	// 			FirstName: to.Ptr("Test"),
	// 			LastName: to.Ptr("User"),
	// 			PhoneNumber: to.Ptr("000-000-0000"),
	// 			PostalCode: to.Ptr("00000"),
	// 			Region: to.Ptr("WA"),
	// 		},
	// 	},
	// }
}
