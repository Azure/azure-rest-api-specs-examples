package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/AzureFirewallPutInHub.json
func ExampleAzureFirewallsClient_BeginCreateOrUpdate_createAzureFirewallInVirtualHub() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewAzureFirewallsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "azurefirewall", armnetwork.AzureFirewall{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armnetwork.AzureFirewallPropertiesFormat{
			FirewallPolicy: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/policy1"),
			},
			HubIPAddresses: &armnetwork.HubIPAddresses{
				PublicIPs: &armnetwork.HubPublicIPAddresses{
					Addresses: []*armnetwork.AzureFirewallPublicIPAddress{},
					Count:     to.Ptr[int32](1),
				},
			},
			SKU: &armnetwork.AzureFirewallSKU{
				Name: to.Ptr(armnetwork.AzureFirewallSKUNameAZFWHub),
				Tier: to.Ptr(armnetwork.AzureFirewallSKUTierStandard),
			},
			ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
			VirtualHub: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
			},
		},
		Zones: []*string{},
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
