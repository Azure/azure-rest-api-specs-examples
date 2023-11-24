package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountGetPropertiesCMKVersionExpirationTime.json
func ExampleAccountsClient_GetProperties_storageAccountGetPropertiesCmkVersionExpirationTime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().GetProperties(ctx, "res9407", "sto8596", &armstorage.AccountsClientGetPropertiesOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armstorage.Account{
	// 	Name: to.Ptr("sto8596"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9407/providers/Microsoft.Storage/storageAccounts/sto8596"),
	// 	Location: to.Ptr("eastus2(stage)"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Identity: &armstorage.Identity{
	// 		Type: to.Ptr(armstorage.IdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("911871cc-ffd1-4fc4-ac11-7a316433ea66"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Kind: to.Ptr(armstorage.KindStorage),
	// 	Properties: &armstorage.AccountProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-01T02:42:41.763Z"); return t}()),
	// 		Encryption: &armstorage.Encryption{
	// 			KeySource: to.Ptr(armstorage.KeySourceMicrosoftKeyvault),
	// 			KeyVaultProperties: &armstorage.KeyVaultProperties{
	// 				CurrentVersionedKeyExpirationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-13T20:36:23.702Z"); return t}()),
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
	// 					LastEnabledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-11T20:49:31.703Z"); return t}()),
	// 				},
	// 				File: &armstorage.EncryptionService{
	// 					Enabled: to.Ptr(true),
	// 					KeyType: to.Ptr(armstorage.KeyTypeAccount),
	// 					LastEnabledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-11T20:49:31.703Z"); return t}()),
	// 				},
	// 			},
	// 		},
	// 		GeoReplicationStats: &armstorage.GeoReplicationStats{
	// 			CanFailover: to.Ptr(true),
	// 			LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-30T00:25:34.000Z"); return t}()),
	// 			Status: to.Ptr(armstorage.GeoReplicationStatusLive),
	// 		},
	// 		IsHnsEnabled: to.Ptr(true),
	// 		NetworkRuleSet: &armstorage.NetworkRuleSet{
	// 			Bypass: to.Ptr(armstorage.BypassAzureServices),
	// 			DefaultAction: to.Ptr(armstorage.DefaultActionAllow),
	// 			IPRules: []*armstorage.IPRule{
	// 			},
	// 			ResourceAccessRules: []*armstorage.ResourceAccessRule{
	// 			},
	// 			VirtualNetworkRules: []*armstorage.VirtualNetworkRule{
	// 			},
	// 		},
	// 		PrimaryEndpoints: &armstorage.Endpoints{
	// 			Blob: to.Ptr("https://sto8596.blob.core.windows.net/"),
	// 			Dfs: to.Ptr("https://sto8596.dfs.core.windows.net/"),
	// 			File: to.Ptr("https://sto8596.file.core.windows.net/"),
	// 			InternetEndpoints: &armstorage.AccountInternetEndpoints{
	// 				Blob: to.Ptr("https://sto8596-internetrouting.blob.core.windows.net/"),
	// 				Dfs: to.Ptr("https://sto8596-internetrouting.dfs.core.windows.net/"),
	// 				File: to.Ptr("https://sto8596-internetrouting.file.core.windows.net/"),
	// 				Web: to.Ptr("https://sto8596-internetrouting.web.core.windows.net/"),
	// 			},
	// 			MicrosoftEndpoints: &armstorage.AccountMicrosoftEndpoints{
	// 				Blob: to.Ptr("https://sto8596-microsoftrouting.blob.core.windows.net/"),
	// 				Dfs: to.Ptr("https://sto8596-microsoftrouting.dfs.core.windows.net/"),
	// 				File: to.Ptr("https://sto8596-microsoftrouting.file.core.windows.net/"),
	// 				Queue: to.Ptr("https://sto8596-microsoftrouting.queue.core.windows.net/"),
	// 				Table: to.Ptr("https://sto8596-microsoftrouting.table.core.windows.net/"),
	// 				Web: to.Ptr("https://sto8596-microsoftrouting.web.core.windows.net/"),
	// 			},
	// 			Queue: to.Ptr("https://sto8596.queue.core.windows.net/"),
	// 			Table: to.Ptr("https://sto8596.table.core.windows.net/"),
	// 			Web: to.Ptr("https://sto8596.web.core.windows.net/"),
	// 		},
	// 		PrimaryLocation: to.Ptr("eastus2(stage)"),
	// 		ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 		RoutingPreference: &armstorage.RoutingPreference{
	// 			PublishInternetEndpoints: to.Ptr(true),
	// 			PublishMicrosoftEndpoints: to.Ptr(true),
	// 			RoutingChoice: to.Ptr(armstorage.RoutingChoiceMicrosoftRouting),
	// 		},
	// 		SecondaryLocation: to.Ptr("northcentralus(stage)"),
	// 		StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
	// 		StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
	// 		EnableHTTPSTrafficOnly: to.Ptr(false),
	// 	},
	// 	SKU: &armstorage.SKU{
	// 		Name: to.Ptr(armstorage.SKUNameStandardGRS),
	// 		Tier: to.Ptr(armstorage.SKUTierStandard),
	// 	},
	// }
}
