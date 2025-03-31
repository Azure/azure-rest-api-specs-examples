package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/createSecret.json
func ExampleSecretsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecretsClient().CreateOrUpdate(ctx, "sample-group", "sample-vault", "secret-name", armkeyvault.SecretCreateOrUpdateParameters{
		Properties: &armkeyvault.SecretProperties{
			Value: to.Ptr("secret-value"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Secret = armkeyvault.Secret{
	// 	Name: to.Ptr("secret-name"),
	// 	Type: to.Ptr("Microsoft.KeyVault/vaults/secrets"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault/secrets/secret-name"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armkeyvault.SecretProperties{
	// 		Attributes: &armkeyvault.SecretAttributes{
	// 			Created: to.Ptr(time.Unix(1514938738, 0)),
	// 			Enabled: to.Ptr(true),
	// 			Updated: to.Ptr(time.Unix(1514938738, 0)),
	// 		},
	// 		SecretURI: to.Ptr("https://sample-vault.vault.azure.net/secrets/secret-name"),
	// 		SecretURIWithVersion: to.Ptr("https:/sample-vault.vault.azure.net/secrets/secret-name/baf6de32c4774c7c81345f6476cf90a4"),
	// 	},
	// }
}
