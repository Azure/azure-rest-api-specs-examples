package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/billingAccountsGet.json
func ExampleAccountsClient_Get_billingAccountsGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Get(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.AccountsClientGetResponse{
	// 	Account: armbilling.Account{
	// 		Name: to.Ptr("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
	// 		Type: to.Ptr("Microsoft.Billing/billingAccounts"),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
	// 		Properties: &armbilling.AccountProperties{
	// 			AccountStatus: to.Ptr(armbilling.AccountStatusActive),
	// 			AccountSubType: to.Ptr(armbilling.AccountSubTypeEnterprise),
	// 			AccountType: to.Ptr(armbilling.AccountTypeBusiness),
	// 			AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
	// 			DisplayName: to.Ptr("Premier Business Account"),
	// 			HasReadAccess: to.Ptr(true),
	// 			PrimaryBillingTenantID: to.Ptr("20000000-0000-0000-0000-000000000001"),
	// 		},
	// 		SystemData: &armbilling.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T22:39:34.2606750Z"); return t}()),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.2606750Z"); return t}()),
	// 		},
	// 	},
	// }
}
