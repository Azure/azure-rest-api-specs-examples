package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/listKeyVersions.json
func ExampleKeysClient_NewListVersionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKeysClient().NewListVersionsPager("sample-group", "sample-vault-name", "sample-key-name", nil)
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
		// page.KeyListResult = armkeyvault.KeyListResult{
		// 	Value: []*armkeyvault.Key{
		// 		{
		// 			Name: to.Ptr("c2296aa24acf4daf86942bff5aca73dd"),
		// 			Type: to.Ptr("Microsoft.KeyVault/vaults/keys/versions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault-name/keys/sample-key-name/versions/c2296aa24acf4daf86942bff5aca73dd"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkeyvault.KeyProperties{
		// 				Attributes: &armkeyvault.KeyAttributes{
		// 					Created: to.Ptr[int64](1598641074),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1598641074),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name"),
		// 				KeyURIWithVersion: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name/c2296aa24acf4daf86942bff5aca73dd"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("d5a04667b6f44b0ca62825f5eae93da6"),
		// 			Type: to.Ptr("Microsoft.KeyVault/vaults/keys/versions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault-name/keys/sample-key-name/versions/d5a04667b6f44b0ca62825f5eae93da6"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkeyvault.KeyProperties{
		// 				Attributes: &armkeyvault.KeyAttributes{
		// 					Created: to.Ptr[int64](1598641295),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1598641295),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name"),
		// 				KeyURIWithVersion: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name/d5a04667b6f44b0ca62825f5eae93da6"),
		// 			},
		// 	}},
		// }
	}
}
