Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv1.0.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VpnGatewayPut.json
func ExampleVPNGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewVPNGatewaysClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"gateway1",
		armnetwork.VPNGateway{
			Location: to.Ptr("westcentralus"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Properties: &armnetwork.VPNGatewayProperties{
				BgpSettings: &armnetwork.BgpSettings{
					Asn: to.Ptr[int64](65515),
					BgpPeeringAddresses: []*armnetwork.IPConfigurationBgpPeeringAddress{
						{
							CustomBgpIPAddresses: []*string{
								to.Ptr("169.254.21.5")},
							IPConfigurationID: to.Ptr("Instance0"),
						},
						{
							CustomBgpIPAddresses: []*string{
								to.Ptr("169.254.21.10")},
							IPConfigurationID: to.Ptr("Instance1"),
						}},
					PeerWeight: to.Ptr[int32](0),
				},
				Connections: []*armnetwork.VPNConnection{
					{
						Name: to.Ptr("vpnConnection1"),
						Properties: &armnetwork.VPNConnectionProperties{
							RemoteVPNSite: &armnetwork.SubResource{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"),
							},
							VPNLinkConnections: []*armnetwork.VPNSiteLinkConnection{
								{
									Name: to.Ptr("Connection-Link1"),
									Properties: &armnetwork.VPNSiteLinkConnectionProperties{
										ConnectionBandwidth: to.Ptr[int32](200),
										EgressNatRules: []*armnetwork.SubResource{
											{
												ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/natRules/nat03"),
											}},
										SharedKey:                 to.Ptr("key"),
										VPNConnectionProtocolType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
										VPNSiteLink: &armnetwork.SubResource{
											ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/siteLink1"),
										},
									},
								}},
						},
					}},
				EnableBgpRouteTranslationForNat: to.Ptr(false),
				IsRoutingPreferenceInternet:     to.Ptr(false),
				NatRules: []*armnetwork.VPNGatewayNatRule{
					{
						Name: to.Ptr("nat03"),
						Properties: &armnetwork.VPNGatewayNatRuleProperties{
							Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
							ExternalMappings: []*armnetwork.VPNNatRuleMapping{
								{
									AddressSpace: to.Ptr("192.168.0.0/26"),
								}},
							InternalMappings: []*armnetwork.VPNNatRuleMapping{
								{
									AddressSpace: to.Ptr("0.0.0.0/26"),
								}},
							IPConfigurationID: to.Ptr(""),
							Mode:              to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
						},
					}},
				VirtualHub: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
				},
			},
		},
		nil)
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
```
