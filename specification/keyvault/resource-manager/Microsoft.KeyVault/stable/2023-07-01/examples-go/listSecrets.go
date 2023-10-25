package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/listSecrets.json
func ExampleSecretsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecretsClient().NewListPager("sample-group", "sample-vault", &armkeyvault.SecretsClientListOptions{Top: nil})
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
		// page.SecretListResult = armkeyvault.SecretListResult{
		// 	Value: []*armkeyvault.Secret{
		// 		{
		// 			Name: to.Ptr("secret-name"),
		// 			Type: to.Ptr("Microsoft.KeyVault/vaults/secrets"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault/secrets/secret-name"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkeyvault.SecretProperties{
		// 				Attributes: &armkeyvault.SecretAttributes{
		// 					Created: to.Ptr(time.Unix(1514941476, 0)),
		// 					Enabled: to.Ptr(true),
		// 					Updated: to.Ptr(time.Unix(1514941476, 0)),
		// 				},
		// 				SecretURI: to.Ptr("https://sample-vault.vault.azure.net/secrets/secret-name"),
		// 				SecretURIWithVersion: to.Ptr("https://sample-vault.vault.azure.net/secrets/secret-name/40af42fbc10047f8a756a73211492f56"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("secret-name2"),
		// 			Type: to.Ptr("Microsoft.KeyVault/vaults/secrets"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault/secrets/secret-name2"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkeyvault.SecretProperties{
		// 				Attributes: &armkeyvault.SecretAttributes{
		// 					Created: to.Ptr(time.Unix(1514941476, 0)),
		// 					Enabled: to.Ptr(true),
		// 					Updated: to.Ptr(time.Unix(1514941476, 0)),
		// 				},
		// 				SecretURI: to.Ptr("https://sample-vault.vault.azure.net/secrets/secret-name2"),
		// 				SecretURIWithVersion: to.Ptr("https://sample-vault.vault.azure.net/secrets/secret-name2/cd7264a6f56c44d1b594423c80609aae"),
		// 			},
		// 	}},
		// }
	}
}
