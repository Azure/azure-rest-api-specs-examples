package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-06-01/StorageAccountCreateUserAssignedIdentityWithFederatedIdentityClientId.json
func ExampleAccountsClient_BeginCreate_storageAccountCreateUserAssignedIdentityWithFederatedIdentityClientId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "res131918", "sto131918", armstorage.AccountCreateParameters{
		Identity: &armstorage.Identity{
			Type: to.Ptr(armstorage.IdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armstorage.UserAssignedIdentity{
				"/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}": {},
			},
		},
		Kind:     to.Ptr(armstorage.KindStorage),
		Location: to.Ptr("eastus"),
		Properties: &armstorage.AccountPropertiesCreateParameters{
			Encryption: &armstorage.Encryption{
				EncryptionIdentity: &armstorage.EncryptionIdentity{
					EncryptionFederatedIdentityClientID: to.Ptr("f83c6b1b-4d34-47e4-bb34-9d83df58b540"),
					EncryptionUserAssignedIdentity:      to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}"),
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
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.AccountsClientCreateResponse{
	// 	Account: armstorage.Account{
	// 		Name: to.Ptr("sto4445"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Storage/storageAccounts/sto4445"),
	// 		Identity: &armstorage.Identity{
	// 			Type: to.Ptr(armstorage.IdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armstorage.UserAssignedIdentity{
	// 				"/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}": &armstorage.UserAssignedIdentity{
	// 					ClientID: to.Ptr("fbaa6278-1ecc-415c-819f-6e2058d3acb5"),
	// 					PrincipalID: to.Ptr("8d823284-1060-42a5-9ec4-ed3d831e24d7"),
	// 				},
	// 			},
	// 		},
	// 		Kind: to.Ptr(armstorage.KindStorageV2),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armstorage.AccountProperties{
	// 			AccessTier: to.Ptr(armstorage.AccessTierHot),
	// 			CreationTime: to.Ptr(time.Date(2020, time.December, 15, 0, 43, 14, 83909300, time.UTC)),
	// 			Encryption: &armstorage.Encryption{
	// 				EncryptionIdentity: &armstorage.EncryptionIdentity{
	// 					EncryptionFederatedIdentityClientID: to.Ptr("f83c6b1b-4d34-47e4-bb34-9d83df58b540"),
	// 					EncryptionUserAssignedIdentity: to.Ptr("/subscriptions/{subscription-id}/resourcegroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}"),
	// 				},
	// 				KeySource: to.Ptr(armstorage.KeySourceMicrosoftKeyvault),
	// 				KeyVaultProperties: &armstorage.KeyVaultProperties{
	// 					CurrentVersionedKeyIdentifier: to.Ptr("https://myvault8569.vault.azure.net/keys/wrappingKey/0682afdd9c104f4285df20107e956cad"),
	// 					KeyName: to.Ptr("wrappingKey"),
	// 					KeyVaultURI: to.Ptr("https://myvault8569.vault.azure.net"),
	// 					KeyVersion: to.Ptr(""),
	// 					LastKeyRotationTimestamp: to.Ptr(time.Date(2019, time.December, 13, 20, 36, 23, 702329000, time.UTC)),
	// 				},
	// 				Services: &armstorage.EncryptionServices{
	// 					Blob: &armstorage.EncryptionService{
	// 						Enabled: to.Ptr(true),
	// 						KeyType: to.Ptr(armstorage.KeyTypeAccount),
	// 						LastEnabledTime: to.Ptr(time.Date(2020, time.December, 15, 0, 43, 14, 173958700, time.UTC)),
	// 					},
	// 					File: &armstorage.EncryptionService{
	// 						Enabled: to.Ptr(true),
	// 						KeyType: to.Ptr(armstorage.KeyTypeAccount),
	// 						LastEnabledTime: to.Ptr(time.Date(2020, time.December, 15, 0, 43, 14, 173958700, time.UTC)),
	// 					},
	// 				},
	// 			},
	// 			NetworkRuleSet: &armstorage.NetworkRuleSet{
	// 				Bypass: to.Ptr(armstorage.BypassAzureServices),
	// 				DefaultAction: to.Ptr(armstorage.DefaultActionAllow),
	// 				IPRules: []*armstorage.IPRule{
	// 				},
	// 				VirtualNetworkRules: []*armstorage.VirtualNetworkRule{
	// 				},
	// 			},
	// 			PrimaryEndpoints: &armstorage.Endpoints{
	// 				Blob: to.Ptr("https://sto4445.blob.core.windows.net/"),
	// 				Dfs: to.Ptr("https://sto4445.dfs.core.windows.net/"),
	// 				File: to.Ptr("https://sto4445.file.core.windows.net/"),
	// 				Queue: to.Ptr("https://sto4445.queue.core.windows.net/"),
	// 				Table: to.Ptr("https://sto4445.table.core.windows.net/"),
	// 				Web: to.Ptr("https://sto4445.web.core.windows.net/"),
	// 			},
	// 			PrimaryLocation: to.Ptr("eastus"),
	// 			PrivateEndpointConnections: []*armstorage.PrivateEndpointConnection{
	// 			},
	// 			ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 			StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
	// 			EnableHTTPSTrafficOnly: to.Ptr(true),
	// 		},
	// 		SKU: &armstorage.SKU{
	// 			Name: to.Ptr(armstorage.SKUNameStandardLRS),
	// 			Tier: to.Ptr(armstorage.SKUTierStandard),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
