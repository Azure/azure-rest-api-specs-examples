package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fb3217991ff57b5760525aeba1a0670bfe0880fa/specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/Accounts_CreateOrUpdate.json
func ExampleAccountsClient_BeginCreateOrUpdate_accountsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreateOrUpdate(ctx, "myRG", "account1", armnetapp.Account{
		Location:   to.Ptr("eastus"),
		Properties: &armnetapp.AccountProperties{},
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
	// res.Account = armnetapp.Account{
	// 	Name: to.Ptr("account1"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetapp.AccountProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
