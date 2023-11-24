package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountUpdateUserAssignedIdentityWithFederatedIdentityClientId.json
func ExampleAccountsClient_Update_storageAccountUpdateUserAssignedIdentityWithFederatedIdentityClientId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Update(ctx, "res131918", "sto131918", armstorage.AccountUpdateParameters{
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
					EncryptionFederatedIdentityClientID: to.Ptr("3109d1c4-a5de-4d84-8832-feabb916a4b6"),
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armstorage.Account{
	// 	Name: to.Ptr("sto4445"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Storage/storageAccounts/sto4445"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armstorage.Identity{
	// 		Type: to.Ptr(armstorage.IdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armstorage.UserAssignedIdentity{
	// 			"/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}": &armstorage.UserAssignedIdentity{
	// 				ClientID: to.Ptr("fbaa6278-1ecc-415c-819f-6e2058d3acb5"),
	// 				PrincipalID: to.Ptr("8d823284-1060-42a5-9ec4-ed3d831e24d7"),
	// 			},
	// 		},
	// 	},
	// 	Kind: to.Ptr(armstorage.KindStorageV2),
	// 	Properties: &armstorage.AccountProperties{
	// 		AccessTier: to.Ptr(armstorage.AccessTierHot),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-15T00:43:14.083Z"); return t}()),
	// 		Encryption: &armstorage.Encryption{
	// 			EncryptionIdentity: &armstorage.EncryptionIdentity{
	// 				EncryptionFederatedIdentityClientID: to.Ptr("3109d1c4-a5de-4d84-8832-feabb916a4b6"),
	// 				EncryptionUserAssignedIdentity: to.Ptr("/subscriptions/{subscription-id}/resourcegroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}"),
	// 			},
	// 			KeySource: to.Ptr(armstorage.KeySourceMicrosoftKeyvault),
	// 			KeyVaultProperties: &armstorage.KeyVaultProperties{
	// 				CurrentVersionedKeyIdentifier: to.Ptr("https://myvault8569.vault.azure.net/keys/wrappingKey/0682afdd9c104f4285df20107e956cad"),
	// 				KeyName: to.Ptr("wrappingKey"),
	// 				KeyVaultURI: to.Ptr("https://myvault8569.vault.azure.net"),
	// 				KeyVersion: to.Ptr(""),
	// 				LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-13T20:36:23.702Z"); return t}()),
	// 			},
	// 			Services: &armstorage.EncryptionServices{
	// 				Blob: &armstorage.EncryptionService{
	// 					Enabled: to.Ptr(true),
	// 					KeyType: to.Ptr(armstorage.KeyTypeAccount),
	// 					LastEnabledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-15T00:43:14.173Z"); return t}()),
	// 				},
	// 				File: &armstorage.EncryptionService{
	// 					Enabled: to.Ptr(true),
	// 					KeyType: to.Ptr(armstorage.KeyTypeAccount),
	// 					LastEnabledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-15T00:43:14.173Z"); return t}()),
	// 				},
	// 			},
	// 		},
	// 		NetworkRuleSet: &armstorage.NetworkRuleSet{
	// 			Bypass: to.Ptr(armstorage.BypassAzureServices),
	// 			DefaultAction: to.Ptr(armstorage.DefaultActionAllow),
	// 			IPRules: []*armstorage.IPRule{
	// 			},
	// 			VirtualNetworkRules: []*armstorage.VirtualNetworkRule{
	// 			},
	// 		},
	// 		PrimaryEndpoints: &armstorage.Endpoints{
	// 			Blob: to.Ptr("https://sto4445.blob.core.windows.net/"),
	// 			Dfs: to.Ptr("https://sto4445.dfs.core.windows.net/"),
	// 			File: to.Ptr("https://sto4445.file.core.windows.net/"),
	// 			Queue: to.Ptr("https://sto4445.queue.core.windows.net/"),
	// 			Table: to.Ptr("https://sto4445.table.core.windows.net/"),
	// 			Web: to.Ptr("https://sto4445.web.core.windows.net/"),
	// 		},
	// 		PrimaryLocation: to.Ptr("eastus"),
	// 		PrivateEndpointConnections: []*armstorage.PrivateEndpointConnection{
	// 		},
	// 		ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 		StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
	// 		EnableHTTPSTrafficOnly: to.Ptr(true),
	// 	},
	// 	SKU: &armstorage.SKU{
	// 		Name: to.Ptr(armstorage.SKUNameStandardLRS),
	// 		Tier: to.Ptr(armstorage.SKUTierStandard),
	// 	},
	// }
}
