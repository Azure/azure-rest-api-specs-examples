package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/NfsV3AccountCreate.json
func ExampleAccountsClient_BeginCreate_nfsV3AccountCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewAccountsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "res9101", "sto4445", armstorage.AccountCreateParameters{
		Kind:     to.Ptr(armstorage.KindBlockBlobStorage),
		Location: to.Ptr("eastus"),
		Properties: &armstorage.AccountPropertiesCreateParameters{
			IsHnsEnabled: to.Ptr(true),
			EnableNfsV3:  to.Ptr(true),
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
	// TODO: use response item
	_ = res
}
