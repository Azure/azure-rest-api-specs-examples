package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-06-01/StorageContextCacheCRUD/ContextCaches_Update_CustomerManagedKey.json
func ExampleContextCachesClient_BeginUpdate_updateAAzureContextCacheAccountSCustomerManagedKeyEncryptionSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContextCachesClient().BeginUpdate(ctx, "testrg", "testaccount", armstorage.ContextCacheUpdate{
		Tags: map[string]*string{
			"environment": to.Ptr("production"),
			"team":        to.Ptr("context-cache"),
		},
		Identity: &armstorage.ManagedServiceIdentity{
			Type: to.Ptr(armstorage.ManagedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armstorage.ContextCachePropertiesUpdate{
			Description: to.Ptr("Updated Prompt Service account description"),
			Encryption: &armstorage.ArmEncryption{
				CustomerManagedKeyEncryption: &armstorage.CustomerManagedKeyEncryption{
					KeyEncryptionKeyIdentity: &armstorage.KeyEncryptionKeyIdentity{
						IdentityType: to.Ptr(armstorage.KeyEncryptionKeyIdentityTypeSystemAssignedIdentity),
					},
					KeyEncryptionKeyURL: to.Ptr("https://mykeyvault.vault.azure.net/keys/newEncryptionKey"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.ContextCachesClientUpdateResponse{
	// 	ContextCache: armstorage.ContextCache{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/contextCaches/testaccount"),
	// 		Name: to.Ptr("testaccount"),
	// 		Type: to.Ptr("Microsoft.Storage/contextCaches"),
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 			"team": to.Ptr("context-cache"),
	// 		},
	// 		Identity: &armstorage.ManagedServiceIdentity{
	// 			Type: to.Ptr(armstorage.ManagedServiceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		Properties: &armstorage.ContextCacheProperties{
	// 			AccountKind: to.Ptr(armstorage.ContextCacheAccountKindRegional),
	// 			Description: to.Ptr("Updated Prompt Service account description"),
	// 			Encryption: &armstorage.ArmEncryption{
	// 				CustomerManagedKeyEncryption: &armstorage.CustomerManagedKeyEncryption{
	// 					KeyEncryptionKeyIdentity: &armstorage.KeyEncryptionKeyIdentity{
	// 						IdentityType: to.Ptr(armstorage.KeyEncryptionKeyIdentityTypeSystemAssignedIdentity),
	// 					},
	// 					KeyEncryptionKeyURL: to.Ptr("https://mykeyvault.vault.azure.net/keys/newEncryptionKey"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armstorage.ContextCacheProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armstorage.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armstorage.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armstorage.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
