Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/AzureFirewallPut.json
func ExampleAzureFirewallsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewAzureFirewallsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<azure-firewall-name>",
		armnetwork.AzureFirewall{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
			},
			Properties: &armnetwork.AzureFirewallPropertiesFormat{
				ApplicationRuleCollections: []*armnetwork.AzureFirewallApplicationRuleCollection{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.AzureFirewallApplicationRuleCollectionPropertiesFormat{
							Action: &armnetwork.AzureFirewallRCAction{
								Type: armnetwork.AzureFirewallRCActionType("Deny").ToPtr(),
							},
							Priority: to.Int32Ptr(110),
							Rules: []*armnetwork.AzureFirewallApplicationRule{
								{
									Name:        to.StringPtr("<name>"),
									Description: to.StringPtr("<description>"),
									Protocols: []*armnetwork.AzureFirewallApplicationRuleProtocol{
										{
											Port:         to.Int32Ptr(443),
											ProtocolType: armnetwork.AzureFirewallApplicationRuleProtocolType("Https").ToPtr(),
										}},
									SourceAddresses: []*string{
										to.StringPtr("216.58.216.164"),
										to.StringPtr("10.0.0.0/24")},
									TargetFqdns: []*string{
										to.StringPtr("www.test.com")},
								}},
						},
					}},
				IPConfigurations: []*armnetwork.AzureFirewallIPConfiguration{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.AzureFirewallIPConfigurationPropertiesFormat{
							PublicIPAddress: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							Subnet: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
						},
					}},
				NatRuleCollections: []*armnetwork.AzureFirewallNatRuleCollection{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.AzureFirewallNatRuleCollectionProperties{
							Action: &armnetwork.AzureFirewallNatRCAction{
								Type: armnetwork.AzureFirewallNatRCActionType("Dnat").ToPtr(),
							},
							Priority: to.Int32Ptr(112),
							Rules: []*armnetwork.AzureFirewallNatRule{
								{
									Name:        to.StringPtr("<name>"),
									Description: to.StringPtr("<description>"),
									DestinationAddresses: []*string{
										to.StringPtr("1.2.3.4")},
									DestinationPorts: []*string{
										to.StringPtr("443")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										armnetwork.AzureFirewallNetworkRuleProtocol("TCP").ToPtr()},
									SourceAddresses: []*string{
										to.StringPtr("*")},
									TranslatedAddress: to.StringPtr("<translated-address>"),
									TranslatedPort:    to.StringPtr("<translated-port>"),
								},
								{
									Name:        to.StringPtr("<name>"),
									Description: to.StringPtr("<description>"),
									DestinationAddresses: []*string{
										to.StringPtr("1.2.3.4")},
									DestinationPorts: []*string{
										to.StringPtr("80")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										armnetwork.AzureFirewallNetworkRuleProtocol("TCP").ToPtr()},
									SourceAddresses: []*string{
										to.StringPtr("*")},
									TranslatedFqdn: to.StringPtr("<translated-fqdn>"),
									TranslatedPort: to.StringPtr("<translated-port>"),
								}},
						},
					}},
				NetworkRuleCollections: []*armnetwork.AzureFirewallNetworkRuleCollection{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.AzureFirewallNetworkRuleCollectionPropertiesFormat{
							Action: &armnetwork.AzureFirewallRCAction{
								Type: armnetwork.AzureFirewallRCActionType("Deny").ToPtr(),
							},
							Priority: to.Int32Ptr(112),
							Rules: []*armnetwork.AzureFirewallNetworkRule{
								{
									Name:        to.StringPtr("<name>"),
									Description: to.StringPtr("<description>"),
									DestinationAddresses: []*string{
										to.StringPtr("*")},
									DestinationPorts: []*string{
										to.StringPtr("443-444"),
										to.StringPtr("8443")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										armnetwork.AzureFirewallNetworkRuleProtocol("TCP").ToPtr()},
									SourceAddresses: []*string{
										to.StringPtr("192.168.1.1-192.168.1.12"),
										to.StringPtr("10.1.4.12-10.1.4.255")},
								},
								{
									Name:        to.StringPtr("<name>"),
									Description: to.StringPtr("<description>"),
									DestinationFqdns: []*string{
										to.StringPtr("www.amazon.com")},
									DestinationPorts: []*string{
										to.StringPtr("443-444"),
										to.StringPtr("8443")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										armnetwork.AzureFirewallNetworkRuleProtocol("TCP").ToPtr()},
									SourceAddresses: []*string{
										to.StringPtr("10.2.4.12-10.2.4.255")},
								}},
						},
					}},
				SKU: &armnetwork.AzureFirewallSKU{
					Name: armnetwork.AzureFirewallSKUName("AZFW_VNet").ToPtr(),
					Tier: armnetwork.AzureFirewallSKUTier("Standard").ToPtr(),
				},
				ThreatIntelMode: armnetwork.AzureFirewallThreatIntelMode("Alert").ToPtr(),
			},
			Zones: []*string{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AzureFirewallsClientCreateOrUpdateResult)
}
```
