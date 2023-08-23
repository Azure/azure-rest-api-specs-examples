package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountList.json
func ExampleAccountsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListPager(nil)
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
		// page.AccountListResult = armstorage.AccountListResult{
		// 	Value: []*armstorage.Account{
		// 		{
		// 			Name: to.Ptr("sto1125"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res2627/providers/Microsoft.Storage/storageAccounts/sto1125"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Kind: to.Ptr(armstorage.KindStorage),
		// 			Properties: &armstorage.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T13:28:53.4540398Z"); return t}()),
		// 				Encryption: &armstorage.Encryption{
		// 					KeySource: to.Ptr(armstorage.KeySourceMicrosoftStorage),
		// 					Services: &armstorage.EncryptionServices{
		// 						Blob: &armstorage.EncryptionService{
		// 							Enabled: to.Ptr(true),
		// 							KeyType: to.Ptr(armstorage.KeyTypeAccount),
		// 							LastEnabledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-11T20:49:31.7036140Z"); return t}()),
		// 						},
		// 						File: &armstorage.EncryptionService{
		// 							Enabled: to.Ptr(true),
		// 							KeyType: to.Ptr(armstorage.KeyTypeAccount),
		// 							LastEnabledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-11T20:49:31.7036140Z"); return t}()),
		// 						},
		// 					},
		// 				},
		// 				IsHnsEnabled: to.Ptr(true),
		// 				PrimaryEndpoints: &armstorage.Endpoints{
		// 					Blob: to.Ptr("https://sto1125.blob.core.windows.net/"),
		// 					Dfs: to.Ptr("https://sto1125.dfs.core.windows.net/"),
		// 					File: to.Ptr("https://sto1125.file.core.windows.net/"),
		// 					InternetEndpoints: &armstorage.AccountInternetEndpoints{
		// 						Blob: to.Ptr("https://sto1125-internetrouting.blob.core.windows.net/"),
		// 						Dfs: to.Ptr("https://sto1125-internetrouting.dfs.core.windows.net/"),
		// 						File: to.Ptr("https://sto1125-internetrouting.file.core.windows.net/"),
		// 						Web: to.Ptr("https://sto1125-internetrouting.web.core.windows.net/"),
		// 					},
		// 					MicrosoftEndpoints: &armstorage.AccountMicrosoftEndpoints{
		// 						Blob: to.Ptr("https://sto1125-microsoftrouting.blob.core.windows.net/"),
		// 						Dfs: to.Ptr("https://sto1125-microsoftrouting.dfs.core.windows.net/"),
		// 						File: to.Ptr("https://sto1125-microsoftrouting.file.core.windows.net/"),
		// 						Queue: to.Ptr("https://sto1125-microsoftrouting.queue.core.windows.net/"),
		// 						Table: to.Ptr("https://sto1125-microsoftrouting.table.core.windows.net/"),
		// 						Web: to.Ptr("https://sto1125-microsoftrouting.web.core.windows.net/"),
		// 					},
		// 					Queue: to.Ptr("https://sto1125.queue.core.windows.net/"),
		// 					Table: to.Ptr("https://sto1125.table.core.windows.net/"),
		// 					Web: to.Ptr("https://sto1125.web.core.windows.net/"),
		// 				},
		// 				PrimaryLocation: to.Ptr("eastus"),
		// 				ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
		// 				RoutingPreference: &armstorage.RoutingPreference{
		// 					PublishInternetEndpoints: to.Ptr(true),
		// 					PublishMicrosoftEndpoints: to.Ptr(true),
		// 					RoutingChoice: to.Ptr(armstorage.RoutingChoiceMicrosoftRouting),
		// 				},
		// 				SecondaryLocation: to.Ptr("centraluseuap"),
		// 				StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				EnableHTTPSTrafficOnly: to.Ptr(false),
		// 			},
		// 			SKU: &armstorage.SKU{
		// 				Name: to.Ptr(armstorage.SKUNameStandardGRS),
		// 				Tier: to.Ptr(armstorage.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sto3699"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/testcmk3/providers/Microsoft.Storage/storageAccounts/sto3699"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Identity: &armstorage.Identity{
		// 				Type: to.Ptr(armstorage.IdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("356d057d-cba5-44dd-8a30-b2e547bc416b"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Kind: to.Ptr(armstorage.KindStorage),
		// 			Properties: &armstorage.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T10:06:30.6093014Z"); return t}()),
		// 				PrimaryEndpoints: &armstorage.Endpoints{
		// 					Blob: to.Ptr("https://sto3699.blob.core.windows.net/"),
		// 					File: to.Ptr("https://sto3699.file.core.windows.net/"),
		// 					Queue: to.Ptr("https://sto3699.queue.core.windows.net/"),
		// 					Table: to.Ptr("https://sto3699.table.core.windows.net/"),
		// 				},
		// 				PrimaryLocation: to.Ptr("eastus"),
		// 				ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
		// 				SecondaryLocation: to.Ptr("centraluseuap"),
		// 				StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				EnableHTTPSTrafficOnly: to.Ptr(false),
		// 			},
		// 			SKU: &armstorage.SKU{
		// 				Name: to.Ptr(armstorage.SKUNameStandardGRS),
		// 				Tier: to.Ptr(armstorage.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sto8596"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9407/providers/Microsoft.Storage/storageAccounts/sto8596"),
		// 			Location: to.Ptr("eastus2(stage)"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Identity: &armstorage.Identity{
		// 				Type: to.Ptr(armstorage.IdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("911871cc-ffd1-4fc4-ac11-7a316433ea66"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Kind: to.Ptr(armstorage.KindStorage),
		// 			Properties: &armstorage.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-01T02:42:41.7633306Z"); return t}()),
		// 				Encryption: &armstorage.Encryption{
		// 					KeySource: to.Ptr(armstorage.KeySourceMicrosoftKeyvault),
		// 					KeyVaultProperties: &armstorage.KeyVaultProperties{
		// 						CurrentVersionedKeyIdentifier: to.Ptr("https://myvault8569.vault.azure.net/keys/wrappingKey/0682afdd9c104f4285df20107e956cad"),
		// 						KeyName: to.Ptr("wrappingKey"),
		// 						KeyVaultURI: to.Ptr("https://myvault8569.vault.azure.net"),
		// 						KeyVersion: to.Ptr(""),
		// 						LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-13T20:36:23.7023290Z"); return t}()),
		// 					},
		// 					Services: &armstorage.EncryptionServices{
		// 						Blob: &armstorage.EncryptionService{
		// 							Enabled: to.Ptr(true),
		// 							KeyType: to.Ptr(armstorage.KeyTypeAccount),
		// 							LastEnabledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-11T20:49:31.7036140Z"); return t}()),
		// 						},
		// 						File: &armstorage.EncryptionService{
		// 							Enabled: to.Ptr(true),
		// 							KeyType: to.Ptr(armstorage.KeyTypeAccount),
		// 							LastEnabledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-11T20:49:31.7036140Z"); return t}()),
		// 						},
		// 					},
		// 				},
		// 				GeoReplicationStats: &armstorage.GeoReplicationStats{
		// 					CanFailover: to.Ptr(true),
		// 					LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-30T00:25:34Z"); return t}()),
		// 					Status: to.Ptr(armstorage.GeoReplicationStatusLive),
		// 				},
		// 				IsHnsEnabled: to.Ptr(true),
		// 				NetworkRuleSet: &armstorage.NetworkRuleSet{
		// 					Bypass: to.Ptr(armstorage.BypassAzureServices),
		// 					DefaultAction: to.Ptr(armstorage.DefaultActionAllow),
		// 					IPRules: []*armstorage.IPRule{
		// 					},
		// 					ResourceAccessRules: []*armstorage.ResourceAccessRule{
		// 						{
		// 							ResourceID: to.Ptr("/subscriptions/a7e99807-abbf-4642-bdec-2c809a96a8bc/resourceGroups/res9407/providers/Microsoft.Synapse/workspaces/testworkspace"),
		// 							TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 					}},
		// 					VirtualNetworkRules: []*armstorage.VirtualNetworkRule{
		// 					},
		// 				},
		// 				PrimaryEndpoints: &armstorage.Endpoints{
		// 					Blob: to.Ptr("https://sto8596.blob.core.windows.net/"),
		// 					Dfs: to.Ptr("https://sto8596.dfs.core.windows.net/"),
		// 					File: to.Ptr("https://sto8596.file.core.windows.net/"),
		// 					InternetEndpoints: &armstorage.AccountInternetEndpoints{
		// 						Blob: to.Ptr("https://sto8596-internetrouting.blob.core.windows.net/"),
		// 						Dfs: to.Ptr("https://sto8596-internetrouting.dfs.core.windows.net/"),
		// 						File: to.Ptr("https://sto8596-internetrouting.file.core.windows.net/"),
		// 						Web: to.Ptr("https://sto8596-internetrouting.web.core.windows.net/"),
		// 					},
		// 					MicrosoftEndpoints: &armstorage.AccountMicrosoftEndpoints{
		// 						Blob: to.Ptr("https://sto8596-microsoftrouting.blob.core.windows.net/"),
		// 						Dfs: to.Ptr("https://sto8596-microsoftrouting.dfs.core.windows.net/"),
		// 						File: to.Ptr("https://sto8596-microsoftrouting.file.core.windows.net/"),
		// 						Queue: to.Ptr("https://sto8596-microsoftrouting.queue.core.windows.net/"),
		// 						Table: to.Ptr("https://sto8596-microsoftrouting.table.core.windows.net/"),
		// 						Web: to.Ptr("https://sto8596-microsoftrouting.web.core.windows.net/"),
		// 					},
		// 					Queue: to.Ptr("https://sto8596.queue.core.windows.net/"),
		// 					Table: to.Ptr("https://sto8596.table.core.windows.net/"),
		// 					Web: to.Ptr("https://sto8596.web.core.windows.net/"),
		// 				},
		// 				PrimaryLocation: to.Ptr("eastus2(stage)"),
		// 				ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
		// 				RoutingPreference: &armstorage.RoutingPreference{
		// 					PublishInternetEndpoints: to.Ptr(true),
		// 					PublishMicrosoftEndpoints: to.Ptr(true),
		// 					RoutingChoice: to.Ptr(armstorage.RoutingChoiceMicrosoftRouting),
		// 				},
		// 				SecondaryLocation: to.Ptr("northcentralus(stage)"),
		// 				StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				EnableHTTPSTrafficOnly: to.Ptr(false),
		// 			},
		// 			SKU: &armstorage.SKU{
		// 				Name: to.Ptr(armstorage.SKUNameStandardGRS),
		// 				Tier: to.Ptr(armstorage.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sto6637"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/testcmk3/providers/Microsoft.Storage/storageAccounts/sto6637"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Identity: &armstorage.Identity{
		// 				Type: to.Ptr(armstorage.IdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("911871cc-ffd1-4fc4-ac11-7a316433ea66"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Kind: to.Ptr(armstorage.KindStorage),
		// 			Properties: &armstorage.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T10:09:39.5625175Z"); return t}()),
		// 				PrimaryEndpoints: &armstorage.Endpoints{
		// 					Blob: to.Ptr("https://sto6637.blob.core.windows.net/"),
		// 					File: to.Ptr("https://sto6637.file.core.windows.net/"),
		// 					Queue: to.Ptr("https://sto6637.queue.core.windows.net/"),
		// 					Table: to.Ptr("https://sto6637.table.core.windows.net/"),
		// 				},
		// 				PrimaryLocation: to.Ptr("eastus"),
		// 				ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
		// 				SecondaryLocation: to.Ptr("centraluseuap"),
		// 				StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				EnableHTTPSTrafficOnly: to.Ptr(false),
		// 			},
		// 			SKU: &armstorage.SKU{
		// 				Name: to.Ptr(armstorage.SKUNameStandardGRS),
		// 				Tier: to.Ptr(armstorage.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sto834"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res8186/providers/Microsoft.Storage/storageAccounts/sto834"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Kind: to.Ptr(armstorage.KindStorage),
		// 			Properties: &armstorage.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T13:28:20.8686541Z"); return t}()),
		// 				PrimaryEndpoints: &armstorage.Endpoints{
		// 					Blob: to.Ptr("https://sto834.blob.core.windows.net/"),
		// 					File: to.Ptr("https://sto834.file.core.windows.net/"),
		// 					Queue: to.Ptr("https://sto834.queue.core.windows.net/"),
		// 					Table: to.Ptr("https://sto834.table.core.windows.net/"),
		// 				},
		// 				PrimaryLocation: to.Ptr("eastus"),
		// 				ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
		// 				SecondaryLocation: to.Ptr("centraluseuap"),
		// 				StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				EnableHTTPSTrafficOnly: to.Ptr(false),
		// 			},
		// 			SKU: &armstorage.SKU{
		// 				Name: to.Ptr(armstorage.SKUNameStandardGRS),
		// 				Tier: to.Ptr(armstorage.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sto9174"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/testcmk3/providers/Microsoft.Storage/storageAccounts/sto9174"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Identity: &armstorage.Identity{
		// 				Type: to.Ptr(armstorage.IdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("933e3ddf-1802-4a51-9469-18a33b576f88"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Kind: to.Ptr(armstorage.KindStorage),
		// 			Properties: &armstorage.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T09:46:19.6556989Z"); return t}()),
		// 				PrimaryEndpoints: &armstorage.Endpoints{
		// 					Blob: to.Ptr("https://sto9174.blob.core.windows.net/"),
		// 					File: to.Ptr("https://sto9174.file.core.windows.net/"),
		// 					Queue: to.Ptr("https://sto9174.queue.core.windows.net/"),
		// 					Table: to.Ptr("https://sto9174.table.core.windows.net/"),
		// 				},
		// 				PrimaryLocation: to.Ptr("eastus"),
		// 				ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
		// 				SecondaryLocation: to.Ptr("centraluseuap"),
		// 				StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				EnableHTTPSTrafficOnly: to.Ptr(false),
		// 			},
		// 			SKU: &armstorage.SKU{
		// 				Name: to.Ptr(armstorage.SKUNameStandardGRS),
		// 				Tier: to.Ptr(armstorage.SKUTierStandard),
		// 			},
		// 	}},
		// }
	}
}
