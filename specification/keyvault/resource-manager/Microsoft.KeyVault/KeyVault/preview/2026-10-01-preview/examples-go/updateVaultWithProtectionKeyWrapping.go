package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
)

// Generated from example definition: 2026-10-01-preview/updateVaultWithProtectionKeyWrapping.json
func ExampleVaultsClient_Update_configureProtectionKeyWrappingOnAnExistingVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVaultsClient().Update(ctx, "sample-resource-group", "sample-vault", armkeyvault.VaultPatchParameters{
		Properties: &armkeyvault.VaultPatchProperties{
			ProtectionKeysWrappingKey: &armkeyvault.WrappingKeyConfiguration{
				KeyValue: &armkeyvault.WrappingKeyValue{
					URL: to.Ptr("https://sample-hsm.managedhsm.azure.net/keys/wrapping-key"),
					Kty: to.Ptr("oct"),
				},
				KeyAuthenticationData: &armkeyvault.WrappingKeyAuthenticationData{
					AuthType: to.Ptr("ServiceIdentity"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkeyvault.VaultsClientUpdateResponse{
	// 	Vault: armkeyvault.Vault{
	// 		Name: to.Ptr("sample-vault"),
	// 		Type: to.Ptr("Microsoft.KeyVault/vaults"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-resource-group/providers/Microsoft.KeyVault/vaults/sample-vault"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armkeyvault.VaultProperties{
	// 			ProtectionKeysWrappingKey: &armkeyvault.WrappingKeyConfiguration{
	// 				KeyValue: &armkeyvault.WrappingKeyValue{
	// 					URL: to.Ptr("https://sample-hsm.managedhsm.azure.net/keys/wrapping-key/0123456789abcdef0123456789abcdef"),
	// 					Kty: to.Ptr("oct"),
	// 				},
	// 				KeyAuthenticationData: &armkeyvault.WrappingKeyAuthenticationData{
	// 					AuthType: to.Ptr("ServiceIdentity"),
	// 				},
	// 			},
	// 			SKU: &armkeyvault.SKU{
	// 				Name: to.Ptr(armkeyvault.SKUNameStandard),
	// 				Family: to.Ptr(armkeyvault.SKUFamilyA),
	// 			},
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			VaultURI: to.Ptr("https://sample-vault.vault.azure.net"),
	// 		},
	// 	},
	// }
}
