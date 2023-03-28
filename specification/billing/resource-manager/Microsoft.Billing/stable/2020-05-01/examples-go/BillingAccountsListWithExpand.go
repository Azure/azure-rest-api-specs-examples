package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountsListWithExpand.json
func ExampleAccountsClient_NewListPager_billingAccountsListWithExpand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListPager(&armbilling.AccountsClientListOptions{Expand: to.Ptr("soldTo,billingProfiles,billingProfiles/invoiceSections")})
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
		// page.AccountListResult = armbilling.AccountListResult{
		// 	Value: []*armbilling.Account{
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.AccountProperties{
		// 				AccountStatus: to.Ptr(armbilling.AccountStatusActive),
		// 				AccountType: to.Ptr(armbilling.AccountTypeEnterprise),
		// 				AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
		// 				BillingProfiles: &armbilling.ProfilesOnExpand{
		// 					HasMoreResults: to.Ptr(true),
		// 					Value: []*armbilling.Profile{
		// 						{
		// 							Name: to.Ptr("11000000-0000-0000-0000-000000000000"),
		// 							Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles"),
		// 							ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 							Properties: &armbilling.ProfileProperties{
		// 								BillTo: &armbilling.AddressDetails{
		// 									AddressLine1: to.Ptr("Test Address1"),
		// 									AddressLine2: to.Ptr("Test Address2"),
		// 									AddressLine3: to.Ptr("Test Address3"),
		// 									City: to.Ptr("City"),
		// 									CompanyName: to.Ptr("Contoso"),
		// 									Country: to.Ptr("US"),
		// 									Email: to.Ptr("abc@contoso.com"),
		// 									FirstName: to.Ptr("Test"),
		// 									LastName: to.Ptr("User"),
		// 									PhoneNumber: to.Ptr("000-000-0000"),
		// 									PostalCode: to.Ptr("00000"),
		// 									Region: to.Ptr("WA"),
		// 								},
		// 								BillingRelationshipType: to.Ptr(armbilling.BillingRelationshipTypeDirect),
		// 								Currency: to.Ptr("USD"),
		// 								DisplayName: to.Ptr("BillingProfile1"),
		// 								EnabledAzurePlans: []*armbilling.AzurePlan{
		// 									{
		// 										SKUDescription: to.Ptr("Microsoft Azure Plan"),
		// 										SKUID: to.Ptr("0001"),
		// 									},
		// 									{
		// 										SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
		// 										SKUID: to.Ptr("0002"),
		// 								}},
		// 								HasReadAccess: to.Ptr(true),
		// 								InvoiceDay: to.Ptr[int32](5),
		// 								InvoiceEmailOptIn: to.Ptr(true),
		// 								InvoiceSections: &armbilling.InvoiceSectionsOnExpand{
		// 									HasMoreResults: to.Ptr(false),
		// 									Value: []*armbilling.InvoiceSection{
		// 										{
		// 											Name: to.Ptr("invoiceSectionId1"),
		// 											Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 											ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000000/invoiceSections/invoiceSectionId1"),
		// 											Properties: &armbilling.InvoiceSectionProperties{
		// 												DisplayName: to.Ptr("invoiceSectionName1"),
		// 												Labels: map[string]*string{
		// 													"costCategory": to.Ptr("Support"),
		// 													"pcCode": to.Ptr("A123456"),
		// 												},
		// 												State: to.Ptr(armbilling.InvoiceSectionStateActive),
		// 												SystemID: to.Ptr("9XXX-11XX-XX1-XXXX-XXX"),
		// 											},
		// 										},
		// 										{
		// 											Name: to.Ptr("invoiceSectionId2"),
		// 											Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 											ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000000/invoiceSections/invoiceSectionId2"),
		// 											Properties: &armbilling.InvoiceSectionProperties{
		// 												DisplayName: to.Ptr("invoiceSectionName2"),
		// 												Labels: map[string]*string{
		// 													"costCategory": to.Ptr("Finance"),
		// 													"pcCode": to.Ptr("B223456"),
		// 												},
		// 												State: to.Ptr(armbilling.InvoiceSectionStateActive),
		// 												SystemID: to.Ptr("9XXX-22XX-XX1-XXXX-XXX"),
		// 											},
		// 									}},
		// 								},
		// 								PoNumber: to.Ptr("ABC12345"),
		// 								SpendingLimit: to.Ptr(armbilling.SpendingLimitOn),
		// 								Status: to.Ptr(armbilling.BillingProfileStatusWarned),
		// 								StatusReasonCode: to.Ptr(armbilling.StatusReasonCodePastDue),
		// 								SystemID: to.Ptr("1XXX-11XX-XX1-XXXX-XXX"),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("11000000-0000-0000-0000-000000000001"),
		// 							Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles"),
		// 							ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000001"),
		// 							Properties: &armbilling.ProfileProperties{
		// 								BillTo: &armbilling.AddressDetails{
		// 									AddressLine1: to.Ptr("Test Address1"),
		// 									AddressLine2: to.Ptr("Test Address2"),
		// 									AddressLine3: to.Ptr("Test Address3"),
		// 									City: to.Ptr("City"),
		// 									CompanyName: to.Ptr("Contoso"),
		// 									Country: to.Ptr("US"),
		// 									Email: to.Ptr("abc@contoso.com"),
		// 									FirstName: to.Ptr("Test"),
		// 									LastName: to.Ptr("User"),
		// 									PhoneNumber: to.Ptr("000-000-0000"),
		// 									PostalCode: to.Ptr("00000"),
		// 									Region: to.Ptr("WA"),
		// 								},
		// 								BillingRelationshipType: to.Ptr(armbilling.BillingRelationshipTypeCSPPartner),
		// 								Currency: to.Ptr("USD"),
		// 								DisplayName: to.Ptr("BillingProfile2"),
		// 								EnabledAzurePlans: []*armbilling.AzurePlan{
		// 									{
		// 										SKUDescription: to.Ptr("Microsoft Azure Plan"),
		// 										SKUID: to.Ptr("0001"),
		// 									},
		// 									{
		// 										SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
		// 										SKUID: to.Ptr("0002"),
		// 								}},
		// 								HasReadAccess: to.Ptr(true),
		// 								InvoiceDay: to.Ptr[int32](5),
		// 								InvoiceEmailOptIn: to.Ptr(true),
		// 								InvoiceSections: &armbilling.InvoiceSectionsOnExpand{
		// 									HasMoreResults: to.Ptr(false),
		// 									Value: []*armbilling.InvoiceSection{
		// 										{
		// 											Name: to.Ptr("invoiceSectionId11"),
		// 											Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 											ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000001/invoiceSections/invoiceSectionId11"),
		// 											Properties: &armbilling.InvoiceSectionProperties{
		// 												DisplayName: to.Ptr("invoiceSectionName11"),
		// 												Labels: map[string]*string{
		// 													"costCategory": to.Ptr("Marketing"),
		// 													"pcCode": to.Ptr("Z223456"),
		// 												},
		// 												State: to.Ptr(armbilling.InvoiceSectionStateActive),
		// 												SystemID: to.Ptr("9XXX-33XX-XX1-XXXX-XXX"),
		// 											},
		// 									}},
		// 								},
		// 								PoNumber: to.Ptr("ABC12345"),
		// 								SpendingLimit: to.Ptr(armbilling.SpendingLimitOff),
		// 								Status: to.Ptr(armbilling.BillingProfileStatusActive),
		// 								SystemID: to.Ptr("2XXX-22XX-XX1-XXXX-XXX"),
		// 							},
		// 					}},
		// 				},
		// 				DisplayName: to.Ptr("Test Account 1"),
		// 				HasReadAccess: to.Ptr(true),
		// 				SoldTo: &armbilling.AddressDetails{
		// 					AddressLine1: to.Ptr("Test Address"),
		// 					AddressLine2: to.Ptr("Test Address"),
		// 					AddressLine3: to.Ptr("Test Address"),
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
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("10000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.AccountProperties{
		// 				AccountStatus: to.Ptr(armbilling.AccountStatusActive),
		// 				AccountType: to.Ptr(armbilling.AccountTypeEnterprise),
		// 				AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
		// 				BillingProfiles: &armbilling.ProfilesOnExpand{
		// 					HasMoreResults: to.Ptr(true),
		// 					Value: []*armbilling.Profile{
		// 						{
		// 							Name: to.Ptr("11000000-0000-0000-0000-000000000004"),
		// 							Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles"),
		// 							ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000004"),
		// 							Properties: &armbilling.ProfileProperties{
		// 								BillTo: &armbilling.AddressDetails{
		// 									AddressLine1: to.Ptr("Test Address1"),
		// 									AddressLine2: to.Ptr("Test Address2"),
		// 									AddressLine3: to.Ptr("Test Address3"),
		// 									City: to.Ptr("City"),
		// 									CompanyName: to.Ptr("Contoso Test"),
		// 									Country: to.Ptr("US"),
		// 									Email: to.Ptr("abc@contoso.com"),
		// 									FirstName: to.Ptr("Test"),
		// 									LastName: to.Ptr("User"),
		// 									PhoneNumber: to.Ptr("000-000-0000"),
		// 									PostalCode: to.Ptr("00000"),
		// 									Region: to.Ptr("WA"),
		// 								},
		// 								BillingRelationshipType: to.Ptr(armbilling.BillingRelationshipTypeIndirectPartner),
		// 								Currency: to.Ptr("USD"),
		// 								DisplayName: to.Ptr("BillingProfile3"),
		// 								EnabledAzurePlans: []*armbilling.AzurePlan{
		// 									{
		// 										SKUDescription: to.Ptr("Microsoft Azure Plan"),
		// 										SKUID: to.Ptr("0001"),
		// 									},
		// 									{
		// 										SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
		// 										SKUID: to.Ptr("0002"),
		// 								}},
		// 								HasReadAccess: to.Ptr(true),
		// 								IndirectRelationshipInfo: &armbilling.IndirectRelationshipInfo{
		// 									BillingAccountName: to.Ptr("20000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000"),
		// 									BillingProfileName: to.Ptr("22000000-0000-0000-0000-000000000001"),
		// 									DisplayName: to.Ptr("Customer1"),
		// 								},
		// 								InvoiceDay: to.Ptr[int32](5),
		// 								InvoiceEmailOptIn: to.Ptr(true),
		// 								InvoiceSections: &armbilling.InvoiceSectionsOnExpand{
		// 									HasMoreResults: to.Ptr(true),
		// 									Value: []*armbilling.InvoiceSection{
		// 										{
		// 											Name: to.Ptr("invoiceSectionId3"),
		// 											Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 											ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000004/invoiceSections/invoiceSectionId3"),
		// 											Properties: &armbilling.InvoiceSectionProperties{
		// 												DisplayName: to.Ptr("invoiceSectionName3"),
		// 												Labels: map[string]*string{
		// 													"costCategory": to.Ptr("Support"),
		// 													"pcCode": to.Ptr("C123456"),
		// 												},
		// 												State: to.Ptr(armbilling.InvoiceSectionStateRestricted),
		// 												SystemID: to.Ptr("9XXX-44XX-XX1-XXXX-XXX"),
		// 												TargetCloud: to.Ptr(armbilling.TargetCloudUSNat),
		// 											},
		// 										},
		// 										{
		// 											Name: to.Ptr("invoiceSectionId4"),
		// 											Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 											ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000004/invoiceSections/invoiceSectionId4"),
		// 											Properties: &armbilling.InvoiceSectionProperties{
		// 												DisplayName: to.Ptr("invoiceSectionName4"),
		// 												Labels: map[string]*string{
		// 													"costCategory": to.Ptr("Marketing"),
		// 													"pcCode": to.Ptr("D123456"),
		// 												},
		// 												State: to.Ptr(armbilling.InvoiceSectionStateRestricted),
		// 												SystemID: to.Ptr("9XXX-55XX-XX1-XXXX-XXX"),
		// 												TargetCloud: to.Ptr(armbilling.TargetCloudUSSec),
		// 											},
		// 									}},
		// 								},
		// 								PoNumber: to.Ptr("ABC12345"),
		// 								SpendingLimit: to.Ptr(armbilling.SpendingLimitOn),
		// 								Status: to.Ptr(armbilling.BillingProfileStatusWarned),
		// 								StatusReasonCode: to.Ptr(armbilling.StatusReasonCodePastDue),
		// 								SystemID: to.Ptr("3XXX-33XX-XX1-XXXX-XXX"),
		// 								TargetClouds: []*armbilling.TargetCloud{
		// 									to.Ptr(armbilling.TargetCloudUSNat),
		// 									to.Ptr(armbilling.TargetCloudUSSec)},
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("11000000-0000-0000-0000-000000000005"),
		// 								Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles"),
		// 								ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000005"),
		// 								Properties: &armbilling.ProfileProperties{
		// 									BillTo: &armbilling.AddressDetails{
		// 										AddressLine1: to.Ptr("Test Address1"),
		// 										AddressLine2: to.Ptr("Test Address2"),
		// 										AddressLine3: to.Ptr("Test Address3"),
		// 										City: to.Ptr("City"),
		// 										CompanyName: to.Ptr("Contoso Test"),
		// 										Country: to.Ptr("US"),
		// 										Email: to.Ptr("abc@contoso.com"),
		// 										FirstName: to.Ptr("Test"),
		// 										LastName: to.Ptr("User"),
		// 										PhoneNumber: to.Ptr("000-000-0000"),
		// 										PostalCode: to.Ptr("00000"),
		// 										Region: to.Ptr("WA"),
		// 									},
		// 									BillingRelationshipType: to.Ptr(armbilling.BillingRelationshipTypeIndirectCustomer),
		// 									Currency: to.Ptr("USD"),
		// 									DisplayName: to.Ptr("BillingProfile4"),
		// 									EnabledAzurePlans: []*armbilling.AzurePlan{
		// 										{
		// 											SKUDescription: to.Ptr("Microsoft Azure Plan"),
		// 											SKUID: to.Ptr("0001"),
		// 										},
		// 										{
		// 											SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
		// 											SKUID: to.Ptr("0002"),
		// 									}},
		// 									HasReadAccess: to.Ptr(true),
		// 									IndirectRelationshipInfo: &armbilling.IndirectRelationshipInfo{
		// 										BillingAccountName: to.Ptr("30000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000"),
		// 										BillingProfileName: to.Ptr("33000000-0000-0000-0000-000000000001"),
		// 										DisplayName: to.Ptr("Partner1"),
		// 									},
		// 									InvoiceDay: to.Ptr[int32](5),
		// 									InvoiceEmailOptIn: to.Ptr(true),
		// 									InvoiceSections: &armbilling.InvoiceSectionsOnExpand{
		// 										HasMoreResults: to.Ptr(true),
		// 										Value: []*armbilling.InvoiceSection{
		// 											{
		// 												Name: to.Ptr("invoiceSectionId5"),
		// 												Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 												ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000005/invoiceSections/invoiceSectionId5"),
		// 												Properties: &armbilling.InvoiceSectionProperties{
		// 													DisplayName: to.Ptr("invoiceSectionName5"),
		// 													Labels: map[string]*string{
		// 														"costCategory": to.Ptr("Finance"),
		// 														"pcCode": to.Ptr("E123456"),
		// 													},
		// 													State: to.Ptr(armbilling.InvoiceSectionStateActive),
		// 													SystemID: to.Ptr("9XXX-66XX-XX1-XXXX-XXX"),
		// 												},
		// 											},
		// 											{
		// 												Name: to.Ptr("invoiceSectionId6"),
		// 												Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 												ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000001_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000005/invoiceSections/invoiceSectionId6"),
		// 												Properties: &armbilling.InvoiceSectionProperties{
		// 													DisplayName: to.Ptr("invoiceSectionName6"),
		// 													Labels: map[string]*string{
		// 														"costCategory": to.Ptr("Support"),
		// 														"pcCode": to.Ptr("O123456"),
		// 													},
		// 													State: to.Ptr(armbilling.InvoiceSectionStateRestricted),
		// 													SystemID: to.Ptr("9XXX-77XX-XX1-XXXX-XXX"),
		// 													TargetCloud: to.Ptr(armbilling.TargetCloudUSSec),
		// 												},
		// 										}},
		// 									},
		// 									PoNumber: to.Ptr("ABC12345"),
		// 									SpendingLimit: to.Ptr(armbilling.SpendingLimitOn),
		// 									Status: to.Ptr(armbilling.BillingProfileStatusWarned),
		// 									StatusReasonCode: to.Ptr(armbilling.StatusReasonCodePastDue),
		// 									SystemID: to.Ptr("4XXX-44XX-XX1-XXXX-XXX"),
		// 									TargetClouds: []*armbilling.TargetCloud{
		// 										to.Ptr(armbilling.TargetCloudUSSec)},
		// 									},
		// 							}},
		// 						},
		// 						DisplayName: to.Ptr("Test Account 2"),
		// 						HasReadAccess: to.Ptr(true),
		// 					},
		// 			}},
		// 		}
	}
}
