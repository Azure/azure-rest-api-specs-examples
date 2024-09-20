package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingAccountForLegacyAccountDetails.json
func ExampleAccountsClient_NewListPager_billingAccountForLegacyAccountDetails() {
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
		// 			Name: to.Ptr("20000000-0000-0000-0000-000000000001"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-0000-0000-0000-000000000001"),
		// 			Properties: &armbilling.AccountProperties{
		// 				AccountStatus: to.Ptr(armbilling.AccountStatusActive),
		// 				AccountType: to.Ptr(armbilling.AccountTypeIndividual),
		// 				AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftOnlineServicesProgram),
		// 				DisplayName: to.Ptr("Individual Account 2"),
		// 				HasReadAccess: to.Ptr(true),
		// 				NotificationEmailAddress: to.Ptr("individual@domain.com"),
		// 			},
		// 	}},
		// }
	}
}
