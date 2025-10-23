package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
)

// Generated from example definition: 2025-05-01/getKey.json
func ExampleKeysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKeysClient().Get(ctx, "sample-group", "sample-vault-name", "sample-key-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkeyvault.KeysClientGetResponse{
	// 	Key: &armkeyvault.Key{
	// 		Name: to.Ptr("sample-key-name"),
	// 		Type: to.Ptr("Microsoft.KeyVault/vaults/keys"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault-name/keys/sample-key-name"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armkeyvault.KeyProperties{
	// 			Attributes: &armkeyvault.KeyAttributes{
	// 				Created: to.Ptr[int64](1598533051),
	// 				Enabled: to.Ptr(true),
	// 				RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
	// 				Updated: to.Ptr[int64](1598533051),
	// 			},
	// 			KeyOps: []*armkeyvault.JSONWebKeyOperation{
	// 				to.Ptr(armkeyvault.JSONWebKeyOperationEncrypt),
	// 				to.Ptr(armkeyvault.JSONWebKeyOperationDecrypt),
	// 				to.Ptr(armkeyvault.JSONWebKeyOperationSign),
	// 				to.Ptr(armkeyvault.JSONWebKeyOperationVerify),
	// 				to.Ptr(armkeyvault.JSONWebKeyOperationWrapKey),
	// 				to.Ptr(armkeyvault.JSONWebKeyOperationUnwrapKey),
	// 			},
	// 			KeySize: to.Ptr[int32](2048),
	// 			KeyURI: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name"),
	// 			KeyURIWithVersion: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name/fd618d9519b74f9aae94ade66b876acc"),
	// 			Kty: to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
	// 		},
	// 	},
	// }
}
