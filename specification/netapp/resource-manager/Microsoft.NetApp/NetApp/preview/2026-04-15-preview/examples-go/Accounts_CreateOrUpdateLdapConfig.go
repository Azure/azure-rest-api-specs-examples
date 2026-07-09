package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v11"
)

// Generated from example definition: 2026-04-15-preview/Accounts_CreateOrUpdateLdapConfig.json
func ExampleAccountsClient_BeginCreateOrUpdate_accountsCreateOrUpdateLdapConfig() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreateOrUpdate(ctx, "myRG", "account1", armnetapp.Account{
		Location: to.Ptr("eastus"),
		Properties: &armnetapp.AccountProperties{
			LdapConfiguration: &armnetapp.LdapConfiguration{
				Domain: to.Ptr("example.com"),
				LdapServers: []*string{
					to.Ptr("192.0.2.1"),
					to.Ptr("192.0.2.2"),
				},
				LdapOverTLS:             to.Ptr(false),
				ServerCACertificate:     to.Ptr("<REDACTED>"),
				CertificateCNHost:       to.Ptr("ldap.krypton.com"),
				BindAuthenticationLevel: to.Ptr(armnetapp.BindAuthenticationLevelSimple),
				BindDN:                  to.Ptr("cn=user,dc=domain,dc=com"),
				BindPasswordAkvConfig: &armnetapp.BindPasswordAkvConfig{
					AzureKeyVaultURI:     to.Ptr("https://myvault.vault.azure.net/"),
					SecretName:           to.Ptr("bindPassword"),
					UserAssignedIdentity: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.AccountsClientCreateOrUpdateResponse{
	// 	Account: armnetapp.Account{
	// 		Name: to.Ptr("account1"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.AccountProperties{
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			LdapConfiguration: &armnetapp.LdapConfiguration{
	// 				Domain: to.Ptr("example.com"),
	// 				LdapServers: []*string{
	// 					to.Ptr("192.0.2.1"),
	// 					to.Ptr("192.0.2.2"),
	// 				},
	// 				LdapOverTLS: to.Ptr(false),
	// 				CertificateCNHost: to.Ptr("ldap.krypton.com"),
	// 				BindAuthenticationLevel: to.Ptr(armnetapp.BindAuthenticationLevelSimple),
	// 				BindDN: to.Ptr("cn=user,dc=domain,dc=com"),
	// 				BindPasswordAkvConfig: &armnetapp.BindPasswordAkvConfig{
	// 					AzureKeyVaultURI: to.Ptr("https://myvault.vault.azure.net/"),
	// 					SecretName: to.Ptr("bindPassword"),
	// 					UserAssignedIdentity: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
