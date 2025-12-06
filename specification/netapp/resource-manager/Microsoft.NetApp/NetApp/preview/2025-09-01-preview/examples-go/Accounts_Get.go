package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/Accounts_Get.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Get(ctx, "myRG", "account1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.AccountsClientGetResponse{
	// 	Account: &armnetapp.Account{
	// 		Name: to.Ptr("account1"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.AccountProperties{
	// 			ActiveDirectories: []*armnetapp.ActiveDirectory{
	// 				{
	// 					ActiveDirectoryID: to.Ptr("02da3711-6c58-2d64-098a-e3af7afaf936"),
	// 					AesEncryption: to.Ptr(true),
	// 					DNS: to.Ptr("10.10.10.3"),
	// 					Domain: to.Ptr("10.10.10.3"),
	// 					LdapSigning: to.Ptr(true),
	// 					OrganizationalUnit: to.Ptr("OU=Engineering"),
	// 					Site: to.Ptr("SiteName"),
	// 					SmbServerName: to.Ptr("SMBServer"),
	// 					Status: to.Ptr(armnetapp.ActiveDirectoryStatusInUse),
	// 					StatusDetails: to.Ptr("Status Details"),
	// 					Username: to.Ptr("ad_user_name"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	},
	// }
}
