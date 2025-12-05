package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/ElasticAccounts_CreateOrUpdate.json
func ExampleElasticAccountsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewElasticAccountsClient().BeginCreateOrUpdate(ctx, "myRG", "account1", armnetapp.ElasticAccount{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"ac-tag1": to.Ptr("account1"),
		},
		Properties: &armnetapp.ElasticAccountProperties{},
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
	// res = armnetapp.ElasticAccountsClientCreateOrUpdateResponse{
	// 	ElasticAccount: &armnetapp.ElasticAccount{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1"),
	// 		Name: to.Ptr("account1"),
	// 		Type: to.Ptr("Microsoft.NetApp/elasticAccounts"),
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 			"ac-tag1": to.Ptr("account1"),
	// 		},
	// 		Properties: &armnetapp.ElasticAccountProperties{
	// 			ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
	// 			Encryption: &armnetapp.ElasticEncryption{
	// 				KeySource: to.Ptr(armnetapp.KeySource("NetApp")),
	// 				KeyVaultProperties: &armnetapp.ElasticKeyVaultProperties{
	// 					KeyVaultURI: to.Ptr("https://myVault.vault.azure.net/"),
	// 					KeyName: to.Ptr("string"),
	// 					KeyVaultResourceID: to.Ptr("/subscriptions/a1b2c3d4-e5f6-7890-1234-567890abcdef/resourceGroups/my-keyvault-rg/providers/Microsoft.KeyVault/vaults/my-prod-keyvault"),
	// 					Status: to.Ptr(armnetapp.ElasticKeyVaultStatusCreated),
	// 				},
	// 				Identity: &armnetapp.ElasticEncryptionIdentity{
	// 					PrincipalID: to.Ptr("70ab1774-1488-4f9c-a283-2db682f91b98"),
	// 					UserAssignedIdentity: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/snaebjor-rotterdam-westcentralus/providers/Microsoft.ManagedIdentity/userAssignedIdentities/snaebjor-wcu-identity"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
