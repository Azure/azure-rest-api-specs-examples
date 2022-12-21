package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountPutEncryptionScopeWithInfrastructureEncryption.json
func ExampleEncryptionScopesClient_Put_storageAccountPutEncryptionScopeWithInfrastructureEncryption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewEncryptionScopesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Put(ctx, "resource-group-name", "{storage-account-name}", "{encryption-scope-name}", armstorage.EncryptionScope{
		EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
			RequireInfrastructureEncryption: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
