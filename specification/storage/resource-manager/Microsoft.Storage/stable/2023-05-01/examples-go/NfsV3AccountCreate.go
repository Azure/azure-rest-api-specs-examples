package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/NfsV3AccountCreate.json
func ExampleAccountsClient_BeginCreate_nfsV3AccountCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "res9101", "sto4445", armstorage.AccountCreateParameters{
		Kind:     to.Ptr(armstorage.KindBlockBlobStorage),
		Location: to.Ptr("eastus"),
		Properties: &armstorage.AccountPropertiesCreateParameters{
			EnableExtendedGroups: to.Ptr(true),
			IsHnsEnabled:         to.Ptr(true),
			EnableNfsV3:          to.Ptr(true),
			NetworkRuleSet: &armstorage.NetworkRuleSet{
				Bypass:        to.Ptr(armstorage.BypassAzureServices),
				DefaultAction: to.Ptr(armstorage.DefaultActionAllow),
				IPRules:       []*armstorage.IPRule{},
				VirtualNetworkRules: []*armstorage.VirtualNetworkRule{
					{
						VirtualNetworkResourceID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Network/virtualNetworks/net123/subnets/subnet12"),
					}},
			},
			EnableHTTPSTrafficOnly: to.Ptr(false),
		},
		SKU: &armstorage.SKU{
			Name: to.Ptr(armstorage.SKUNamePremiumLRS),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armstorage.Account{
	// 	Name: to.Ptr("sto4445"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Storage/storageAccounts/sto4445"),
	// 	Location: to.Ptr("eastus"),
	// 	Kind: to.Ptr(armstorage.KindBlockBlobStorage),
	// 	Properties: &armstorage.AccountProperties{
	// 		EnableExtendedGroups: to.Ptr(true),
	// 		IsHnsEnabled: to.Ptr(true),
	// 		EnableNfsV3: to.Ptr(true),
	// 		NetworkRuleSet: &armstorage.NetworkRuleSet{
	// 			Bypass: to.Ptr(armstorage.BypassAzureServices),
	// 			DefaultAction: to.Ptr(armstorage.DefaultActionAllow),
	// 			IPRules: []*armstorage.IPRule{
	// 			},
	// 			VirtualNetworkRules: []*armstorage.VirtualNetworkRule{
	// 				{
	// 					VirtualNetworkResourceID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Network/virtualNetworks/net123/subnets/subnet12"),
	// 			}},
	// 		},
	// 		EnableHTTPSTrafficOnly: to.Ptr(false),
	// 	},
	// 	SKU: &armstorage.SKU{
	// 		Name: to.Ptr(armstorage.SKUNamePremiumLRS),
	// 		Tier: to.Ptr(armstorage.SKUTierPremium),
	// 	},
	// }
}
