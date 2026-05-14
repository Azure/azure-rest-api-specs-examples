package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/billingAccountForLegacyAccountDetails.json
func ExampleAccountsClient_NewListPager_billingAccountForLegacyAccountDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListPager(nil)
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
		// page = armbilling.AccountsClientListResponse{
		// 	AccountListResult: armbilling.AccountListResult{
		// 		Value: []*armbilling.Account{
		// 			{
		// 				Name: to.Ptr("20000000-0000-0000-0000-000000000001"),
		// 				Type: to.Ptr("Microsoft.Billing/billingAccounts"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-0000-0000-0000-000000000001"),
		// 				Properties: &armbilling.AccountProperties{
		// 					AccountStatus: to.Ptr(armbilling.AccountStatusActive),
		// 					AccountType: to.Ptr(armbilling.AccountTypeIndividual),
		// 					AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftOnlineServicesProgram),
		// 					DisplayName: to.Ptr("Individual Account 2"),
		// 					HasReadAccess: to.Ptr(true),
		// 					NotificationEmailAddress: to.Ptr("individual@domain.com"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
