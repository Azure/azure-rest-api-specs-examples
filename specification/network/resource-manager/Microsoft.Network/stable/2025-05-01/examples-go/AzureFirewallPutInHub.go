package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/AzureFirewallPutInHub.json
func ExampleAzureFirewallsClient_BeginCreateOrUpdate_createAzureFirewallInVirtualHub() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureFirewallsClient().BeginCreateOrUpdate(ctx, "rg1", "azurefirewall", armnetwork.AzureFirewall{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureFirewall = armnetwork.AzureFirewall{
	// 	Name: to.Ptr("azurefirewall"),
	// 	Type: to.Ptr("Microsoft.Network/azureFirewalls"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.AzureFirewallPropertiesFormat{
	// 		AdditionalProperties: map[string]*string{
	// 		},
	// 		FirewallPolicy: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/policy1"),
	// 		},
	// 		HubIPAddresses: &armnetwork.HubIPAddresses{
	// 			PrivateIPAddress: to.Ptr("10.0.0.0"),
	// 			PublicIPs: &armnetwork.HubPublicIPAddresses{
	// 				Addresses: []*armnetwork.AzureFirewallPublicIPAddress{
	// 					{
	// 						Address: to.Ptr("13.73.240.12"),
	// 				}},
	// 				Count: to.Ptr[int32](1),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		SKU: &armnetwork.AzureFirewallSKU{
	// 			Name: to.Ptr(armnetwork.AzureFirewallSKUNameAZFWHub),
	// 			Tier: to.Ptr(armnetwork.AzureFirewallSKUTierStandard),
	// 		},
	// 		ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
	// 		VirtualHub: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
	// 		},
	// 	},
	// 	Zones: []*string{
	// 	},
	// }
}
