package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/ElasticAccounts_ListByResourceGroup.json
func ExampleElasticAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticAccountsClient().NewListByResourceGroupPager("myRG", nil)
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
		// page = armnetapp.ElasticAccountsClientListByResourceGroupResponse{
		// 	ElasticAccountListResult: armnetapp.ElasticAccountListResult{
		// 		Value: []*armnetapp.ElasticAccount{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1"),
		// 				Name: to.Ptr("account1"),
		// 				Type: to.Ptr("Microsoft.NetApp/elasticAccounts"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"ac-tag1": to.Ptr("account1"),
		// 				},
		// 				Properties: &armnetapp.ElasticAccountProperties{
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 					Encryption: &armnetapp.ElasticEncryption{
		// 						KeySource: to.Ptr(armnetapp.KeySource("NetApp")),
		// 						KeyVaultProperties: &armnetapp.ElasticKeyVaultProperties{
		// 							KeyVaultURI: to.Ptr("https://myVault.vault.azure.net/"),
		// 							KeyName: to.Ptr("string"),
		// 							KeyVaultResourceID: to.Ptr("/subscriptions/a1b2c3d4-e5f6-7890-1234-567890abcdef/resourceGroups/my-keyvault-rg/providers/Microsoft.KeyVault/vaults/my-prod-keyvault"),
		// 							Status: to.Ptr(armnetapp.ElasticKeyVaultStatusCreated),
		// 						},
		// 						Identity: &armnetapp.ElasticEncryptionIdentity{
		// 							PrincipalID: to.Ptr("70ab1774-1488-4f9c-a283-2db682f91b98"),
		// 							UserAssignedIdentity: to.Ptr("/subscriptions/ccdce6ae-b7b3-4a53-b9c5-48e2caa01800/resourcegroups/snaebjor-rotterdam-westcentralus/providers/Microsoft.ManagedIdentity/userAssignedIdentities/snaebjor-wcu-identity"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
