package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/97ee23a6db6078abcbec7b75bf9af8c503e9bb8b/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/StorageAccountEncryptionScopeList.json
func ExampleEncryptionScopesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEncryptionScopesClient().NewListPager("resource-group-name", "accountname", &armstorage.EncryptionScopesClientListOptions{Maxpagesize: nil,
		Filter:  nil,
		Include: nil,
	})
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
		// page.EncryptionScopeListResult = armstorage.EncryptionScopeListResult{
		// 	Value: []*armstorage.EncryptionScope{
		// 		{
		// 			Name: to.Ptr("scope-1"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/encryptionScopes"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resource-group-name/providers/Microsoft.Storage/storageAccounts/accountname/encryptionScopes/scope-1"),
		// 			EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.763Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.763Z"); return t}()),
		// 				Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftStorage),
		// 				State: to.Ptr(armstorage.EncryptionScopeStateEnabled),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("scope-2"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/encryptionScopes"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resource-group-name/providers/Microsoft.Storage/storageAccounts/accountname/encryptionScopes/scope-2"),
		// 			EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T04:32:14.335Z"); return t}()),
		// 				KeyVaultProperties: &armstorage.EncryptionScopeKeyVaultProperties{
		// 					KeyURI: to.Ptr("https://testvault.vault.core.windows.net/keys/key1/863425f1358359c"),
		// 				},
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-17T06:23:14.451Z"); return t}()),
		// 				Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftKeyVault),
		// 				State: to.Ptr(armstorage.EncryptionScopeStateEnabled),
		// 			},
		// 	}},
		// }
	}
}
