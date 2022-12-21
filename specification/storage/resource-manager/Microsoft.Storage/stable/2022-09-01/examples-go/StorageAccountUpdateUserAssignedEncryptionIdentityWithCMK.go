package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountUpdateUserAssignedEncryptionIdentityWithCMK.json
func ExampleAccountsClient_Update_storageAccountUpdateUserAssignedEncryptionIdentityWithCmk() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewAccountsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "res9101", "sto4445", armstorage.AccountUpdateParameters{
		Identity: &armstorage.Identity{
			Type: to.Ptr(armstorage.IdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armstorage.UserAssignedIdentity{
				"/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}": {},
			},
		},
		Kind: to.Ptr(armstorage.KindStorage),
		Properties: &armstorage.AccountPropertiesUpdateParameters{
			Encryption: &armstorage.Encryption{
				EncryptionIdentity: &armstorage.EncryptionIdentity{
					EncryptionUserAssignedIdentity: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}"),
				},
				KeySource: to.Ptr(armstorage.KeySourceMicrosoftKeyvault),
				KeyVaultProperties: &armstorage.KeyVaultProperties{
					KeyName:     to.Ptr("wrappingKey"),
					KeyVaultURI: to.Ptr("https://myvault8569.vault.azure.net"),
					KeyVersion:  to.Ptr(""),
				},
				Services: &armstorage.EncryptionServices{
					Blob: &armstorage.EncryptionService{
						Enabled: to.Ptr(true),
						KeyType: to.Ptr(armstorage.KeyTypeAccount),
					},
					File: &armstorage.EncryptionService{
						Enabled: to.Ptr(true),
						KeyType: to.Ptr(armstorage.KeyTypeAccount),
					},
				},
			},
		},
		SKU: &armstorage.SKU{
			Name: to.Ptr(armstorage.SKUNameStandardLRS),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
