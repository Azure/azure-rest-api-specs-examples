package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
)

// Generated from example definition: 2025-05-01/updateSecret.json
func ExampleSecretsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecretsClient().Update(ctx, "sample-group", "sample-vault", "secret-name", armkeyvault.SecretPatchParameters{
		Properties: &armkeyvault.SecretPatchProperties{
			Value: to.Ptr("secret-value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkeyvault.SecretsClientUpdateResponse{
	// 	Secret: &armkeyvault.Secret{
	// 		Name: to.Ptr("secret-name"),
	// 		Type: to.Ptr("Microsoft.KeyVault/vaults/secrets"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault/secrets/secret-name"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armkeyvault.SecretProperties{
	// 			Attributes: &armkeyvault.SecretAttributes{
	// 				Created: to.Ptr(func() time.Time { t, _ := strconv.ParseInt(1514940684, 10, 64); return time.Unix(t, 0).UTC()}()),
	// 				Enabled: to.Ptr(true),
	// 				Updated: to.Ptr(func() time.Time { t, _ := strconv.ParseInt(1514940698, 10, 64); return time.Unix(t, 0).UTC()}()),
	// 			},
	// 			SecretURI: to.Ptr("https://sample-vault.vault.azure.net/secrets/secret-name"),
	// 			SecretURIWithVersion: to.Ptr("https://sample-vault.vault.azure.net/secrets/secret-name/b8c802f549764f2d97885d152f92ee9d"),
	// 		},
	// 	},
	// }
}
