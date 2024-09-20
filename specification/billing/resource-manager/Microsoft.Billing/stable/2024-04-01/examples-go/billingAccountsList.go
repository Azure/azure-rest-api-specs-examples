package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingAccountsList.json
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
	pager := clientFactory.NewAccountsClient().NewListPager(&armbilling.AccountsClientListOptions{IncludeAll: nil,
		IncludeAllWithoutBillingProfiles: nil,
		IncludeDeleted:                   nil,
		IncludePendingAgreement:          nil,
		IncludeResellee:                  nil,
		LegalOwnerTID:                    nil,
		LegalOwnerOID:                    nil,
		Filter:                           nil,
		Expand:                           nil,
		Top:                              nil,
		Skip:                             nil,
		Search:                           nil,
	})
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
		// 			Name: to.Ptr("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 			SystemData: &armbilling.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T22:39:34.260Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
		// 			},
		// 			Properties: &armbilling.AccountProperties{
		// 				AccountStatus: to.Ptr(armbilling.AccountStatusActive),
		// 				AccountSubType: to.Ptr(armbilling.AccountSubTypeEnterprise),
		// 				AccountType: to.Ptr(armbilling.AccountTypeBusiness),
		// 				AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
		// 				DisplayName: to.Ptr("Premier Business Account"),
		// 				HasReadAccess: to.Ptr(true),
		// 				PrimaryBillingTenantID: to.Ptr("20000000-0000-0000-0000-000000000001"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("20000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 			SystemData: &armbilling.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T22:39:34.260Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
		// 			},
		// 			Properties: &armbilling.AccountProperties{
		// 				AccountStatus: to.Ptr(armbilling.AccountStatusActive),
		// 				AccountSubType: to.Ptr(armbilling.AccountSubTypeProfessional),
		// 				AccountType: to.Ptr(armbilling.AccountTypeBusiness),
		// 				AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
		// 				DisplayName: to.Ptr("Standard Business Account"),
		// 				HasReadAccess: to.Ptr(true),
		// 				PrimaryBillingTenantID: to.Ptr("20000000-0000-0000-0000-000000000001"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("30000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/30000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 			SystemData: &armbilling.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T22:39:34.260Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
		// 			},
		// 			Properties: &armbilling.AccountProperties{
		// 				AccountStatus: to.Ptr(armbilling.AccountStatusActive),
		// 				AccountSubType: to.Ptr(armbilling.AccountSubTypeIndividual),
		// 				AccountType: to.Ptr(armbilling.AccountTypeIndividual),
		// 				AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
		// 				DisplayName: to.Ptr("Individual Account"),
		// 				HasReadAccess: to.Ptr(true),
		// 				PrimaryBillingTenantID: to.Ptr("20000000-0000-0000-0000-000000000001"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("40000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/40000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 			SystemData: &armbilling.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T22:39:34.260Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
		// 			},
		// 			Properties: &armbilling.AccountProperties{
		// 				AccountStatus: to.Ptr(armbilling.AccountStatusActive),
		// 				AccountSubType: to.Ptr(armbilling.AccountSubTypeEnterprise),
		// 				AccountType: to.Ptr(armbilling.AccountTypeBusiness),
		// 				AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftPartnerAgreement),
		// 				DisplayName: to.Ptr("Premier Business Account"),
		// 				HasReadAccess: to.Ptr(true),
		// 				PrimaryBillingTenantID: to.Ptr("20000000-0000-0000-0000-000000000001"),
		// 			},
		// 	}},
		// }
	}
}
