package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountPatchEncryptionScope.json
func ExampleEncryptionScopesClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEncryptionScopesClient().Patch(ctx, "resource-group-name", "accountname", "{encryption-scope-name}", armstorage.EncryptionScope{
		EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
			KeyVaultProperties: &armstorage.EncryptionScopeKeyVaultProperties{
				KeyURI: to.Ptr("https://testvault.vault.core.windows.net/keys/key1/863425f1358359c"),
			},
			Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftKeyVault),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EncryptionScope = armstorage.EncryptionScope{
	// 	Name: to.Ptr("{encryption-scope-name}"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/encryptionScopes"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resource-group-name/providers/Microsoft.Storage/storageAccounts/accountname/encryptionScopes/{encryption-scope-name}"),
	// 	EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.763Z"); return t}()),
	// 		KeyVaultProperties: &armstorage.EncryptionScopeKeyVaultProperties{
	// 			CurrentVersionedKeyIdentifier: to.Ptr("https://testvault.vault.core.windows.net/keys/key1/863425f1358359c"),
	// 			KeyURI: to.Ptr("https://testvault.vault.core.windows.net/keys/key1/863425f1358359c"),
	// 			LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-13T20:36:23.702Z"); return t}()),
	// 		},
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-17T06:23:14.451Z"); return t}()),
	// 		Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftKeyVault),
	// 		State: to.Ptr(armstorage.EncryptionScopeStateEnabled),
	// 	},
	// }
}
