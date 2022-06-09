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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/AzureFirewallPut.json
func ExampleAzureFirewallsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewAzureFirewallsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<azure-firewall-name>",
		armnetwork.AzureFirewall{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Properties: &armnetwork.AzureFirewallPropertiesFormat{
				ApplicationRuleCollections: []*armnetwork.AzureFirewallApplicationRuleCollection{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.AzureFirewallApplicationRuleCollectionPropertiesFormat{
							Action: &armnetwork.AzureFirewallRCAction{
								Type: to.Ptr(armnetwork.AzureFirewallRCActionTypeDeny),
							},
							Priority: to.Ptr[int32](110),
							Rules: []*armnetwork.AzureFirewallApplicationRule{
								{
									Name:        to.Ptr("<name>"),
									Description: to.Ptr("<description>"),
									Protocols: []*armnetwork.AzureFirewallApplicationRuleProtocol{
										{
											Port:         to.Ptr[int32](443),
											ProtocolType: to.Ptr(armnetwork.AzureFirewallApplicationRuleProtocolTypeHTTPS),
										}},
									SourceAddresses: []*string{
										to.Ptr("216.58.216.164"),
										to.Ptr("10.0.0.0/24")},
									TargetFqdns: []*string{
										to.Ptr("www.test.com")},
								}},
						},
					}},
				IPConfigurations: []*armnetwork.AzureFirewallIPConfiguration{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.AzureFirewallIPConfigurationPropertiesFormat{
							PublicIPAddress: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							Subnet: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
						},
					}},
				NatRuleCollections: []*armnetwork.AzureFirewallNatRuleCollection{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.AzureFirewallNatRuleCollectionProperties{
							Action: &armnetwork.AzureFirewallNatRCAction{
								Type: to.Ptr(armnetwork.AzureFirewallNatRCActionTypeDnat),
							},
							Priority: to.Ptr[int32](112),
							Rules: []*armnetwork.AzureFirewallNatRule{
								{
									Name:        to.Ptr("<name>"),
									Description: to.Ptr("<description>"),
									DestinationAddresses: []*string{
										to.Ptr("1.2.3.4")},
									DestinationPorts: []*string{
										to.Ptr("443")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
									SourceAddresses: []*string{
										to.Ptr("*")},
									TranslatedAddress: to.Ptr("<translated-address>"),
									TranslatedPort:    to.Ptr("<translated-port>"),
								},
								{
									Name:        to.Ptr("<name>"),
									Description: to.Ptr("<description>"),
									DestinationAddresses: []*string{
										to.Ptr("1.2.3.4")},
									DestinationPorts: []*string{
										to.Ptr("80")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
									SourceAddresses: []*string{
										to.Ptr("*")},
									TranslatedFqdn: to.Ptr("<translated-fqdn>"),
									TranslatedPort: to.Ptr("<translated-port>"),
								}},
						},
					}},
				NetworkRuleCollections: []*armnetwork.AzureFirewallNetworkRuleCollection{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.AzureFirewallNetworkRuleCollectionPropertiesFormat{
							Action: &armnetwork.AzureFirewallRCAction{
								Type: to.Ptr(armnetwork.AzureFirewallRCActionTypeDeny),
							},
							Priority: to.Ptr[int32](112),
							Rules: []*armnetwork.AzureFirewallNetworkRule{
								{
									Name:        to.Ptr("<name>"),
									Description: to.Ptr("<description>"),
									DestinationAddresses: []*string{
										to.Ptr("*")},
									DestinationPorts: []*string{
										to.Ptr("443-444"),
										to.Ptr("8443")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
									SourceAddresses: []*string{
										to.Ptr("192.168.1.1-192.168.1.12"),
										to.Ptr("10.1.4.12-10.1.4.255")},
								},
								{
									Name:        to.Ptr("<name>"),
									Description: to.Ptr("<description>"),
									DestinationFqdns: []*string{
										to.Ptr("www.amazon.com")},
									DestinationPorts: []*string{
										to.Ptr("443-444"),
										to.Ptr("8443")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
									SourceAddresses: []*string{
										to.Ptr("10.2.4.12-10.2.4.255")},
								}},
						},
					}},
				SKU: &armnetwork.AzureFirewallSKU{
					Name: to.Ptr(armnetwork.AzureFirewallSKUNameAZFWVnet),
					Tier: to.Ptr(armnetwork.AzureFirewallSKUTierStandard),
				},
				ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
			},
			Zones: []*string{},
		},
		&armnetwork.AzureFirewallsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.
