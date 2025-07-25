package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86c6306649b02e542117adb46c61e8019dbd78e9/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/StorageAccountPutEncryptionScopeWithInfrastructureEncryption.json
func ExampleEncryptionScopesClient_Put_storageAccountPutEncryptionScopeWithInfrastructureEncryption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEncryptionScopesClient().Put(ctx, "resource-group-name", "accountname", "{encryption-scope-name}", armstorage.EncryptionScope{
		EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
			RequireInfrastructureEncryption: to.Ptr(true),
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
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.763Z"); return t}()),
	// 		RequireInfrastructureEncryption: to.Ptr(true),
	// 		Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftStorage),
	// 		State: to.Ptr(armstorage.EncryptionScopeStateEnabled),
	// 	},
	// }
}
