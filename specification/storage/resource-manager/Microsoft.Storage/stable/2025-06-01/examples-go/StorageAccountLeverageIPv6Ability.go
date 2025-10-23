package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4e9df3afd38a1cfa00a5d49419dce51bd014601f/specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/StorageAccountLeverageIPv6Ability.json
func ExampleAccountsClient_Update_storageAccountUpdateEnableIpv6Features() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Update(ctx, "res9407", "sto8596", armstorage.AccountUpdateParameters{
		Properties: &armstorage.AccountPropertiesUpdateParameters{
			DualStackEndpointPreference: &armstorage.DualStackEndpointPreference{
				PublishIPv6Endpoint: to.Ptr(true),
			},
			NetworkRuleSet: &armstorage.NetworkRuleSet{
				DefaultAction: to.Ptr(armstorage.DefaultActionDeny),
				IPv6Rules: []*armstorage.IPRule{
					{
						Action:           to.Ptr("Allow"),
						IPAddressOrRange: to.Ptr("2001:0db8:85a3::/64"),
					}},
			},
		},
	}, nil)
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
	// 	Kind: to.Ptr(armstorage.KindStorage),
	// 	Properties: &armstorage.AccountProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-01T02:42:41.763Z"); return t}()),
	// 		DualStackEndpointPreference: &armstorage.DualStackEndpointPreference{
	// 			PublishIPv6Endpoint: to.Ptr(true),
	// 		},
	// 		KeyCreationTime: &armstorage.KeyCreationTime{
	// 			Key1: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-18T04:42:22.432Z"); return t}()),
	// 			Key2: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-18T04:42:22.432Z"); return t}()),
	// 		},
	// 		NetworkRuleSet: &armstorage.NetworkRuleSet{
	// 			Bypass: to.Ptr(armstorage.BypassAzureServices),
	// 			DefaultAction: to.Ptr(armstorage.DefaultActionDeny),
	// 			IPRules: []*armstorage.IPRule{
	// 			},
	// 			IPv6Rules: []*armstorage.IPRule{
	// 				{
	// 					Action: to.Ptr("Allow"),
	// 					IPAddressOrRange: to.Ptr("2001:0db8:85a3::/64"),
	// 			}},
	// 			VirtualNetworkRules: []*armstorage.VirtualNetworkRule{
	// 			},
	// 		},
	// 		PrimaryEndpoints: &armstorage.Endpoints{
	// 			Blob: to.Ptr("https://sto8596.blob.core.windows.net/"),
	// 			Dfs: to.Ptr("https://sto8596.dfs.core.windows.net/"),
	// 			File: to.Ptr("https://sto8596.file.core.windows.net/"),
	// 			IPv6Endpoints: &armstorage.AccountIPv6Endpoints{
	// 				Blob: to.Ptr("https://sto8596-ipv6.blob.core.windows.net/"),
	// 				Dfs: to.Ptr("https://sto8596-ipv6.dfs.core.windows.net/"),
	// 				File: to.Ptr("https://sto8596-ipv6.file.core.windows.net/"),
	// 				Queue: to.Ptr("https://sto8596-ipv6.queue.core.windows.net/"),
	// 				Table: to.Ptr("https://sto8596-ipv6.table.core.windows.net/"),
	// 				Web: to.Ptr("https://sto8596-ipv6.web.core.windows.net/"),
	// 			},
	// 			Queue: to.Ptr("https://sto8596.queue.core.windows.net/"),
	// 			Table: to.Ptr("https://sto8596.table.core.windows.net/"),
	// 			Web: to.Ptr("https://sto8596.web.core.windows.net/"),
	// 		},
	// 		PrimaryLocation: to.Ptr("eastus2(stage)"),
	// 		ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 		StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
	// 	},
	// 	SKU: &armstorage.SKU{
	// 		Name: to.Ptr(armstorage.SKUNameStandardLRS),
	// 		Tier: to.Ptr(armstorage.SKUTierStandard),
	// 	},
	// }
}
