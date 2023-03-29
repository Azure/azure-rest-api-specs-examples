package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountsList.json
func ExampleAccountsClient_NewListPager_billingAccountsList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListPager(&armbilling.AccountsClientListOptions{Expand: nil})
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
		// 				DisplayName: to.Ptr("Test Account 1"),
		// 				HasReadAccess: to.Ptr(true),
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
		// 				DisplayName: to.Ptr("Test Account 2"),
		// 				HasReadAccess: to.Ptr(true),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("20000000-0000-0000-0000-000000000002_00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-0000-0000-0000-000000000002_00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.AccountProperties{
		// 				AccountStatus: to.Ptr(armbilling.AccountStatusActive),
		// 				AccountType: to.Ptr(armbilling.AccountTypeEnterprise),
		// 				AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
		// 				DisplayName: to.Ptr("Test Account 3"),
		// 				HasReadAccess: to.Ptr(true),
		// 			},
		// 	}},
		// }
	}
}
