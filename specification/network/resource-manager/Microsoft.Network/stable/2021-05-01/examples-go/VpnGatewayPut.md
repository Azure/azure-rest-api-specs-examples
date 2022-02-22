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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnGatewayPut.json
func ExampleVPNGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewVPNGatewaysClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<gateway-name>",
		armnetwork.VPNGateway{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
			},
			Properties: &armnetwork.VPNGatewayProperties{
				BgpSettings: &armnetwork.BgpSettings{
					Asn: to.Int64Ptr(65515),
					BgpPeeringAddresses: []*armnetwork.IPConfigurationBgpPeeringAddress{
						{
							CustomBgpIPAddresses: []*string{
								to.StringPtr("169.254.21.5")},
							IPConfigurationID: to.StringPtr("<ipconfiguration-id>"),
						},
						{
							CustomBgpIPAddresses: []*string{
								to.StringPtr("169.254.21.10")},
							IPConfigurationID: to.StringPtr("<ipconfiguration-id>"),
						}},
					PeerWeight: to.Int32Ptr(0),
				},
				Connections: []*armnetwork.VPNConnection{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.VPNConnectionProperties{
							RemoteVPNSite: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							VPNLinkConnections: []*armnetwork.VPNSiteLinkConnection{
								{
									Name: to.StringPtr("<name>"),
									Properties: &armnetwork.VPNSiteLinkConnectionProperties{
										ConnectionBandwidth: to.Int32Ptr(200),
										EgressNatRules: []*armnetwork.SubResource{
											{
												ID: to.StringPtr("<id>"),
											}},
										SharedKey:                 to.StringPtr("<shared-key>"),
										VPNConnectionProtocolType: armnetwork.VirtualNetworkGatewayConnectionProtocol("IKEv2").ToPtr(),
										VPNSiteLink: &armnetwork.SubResource{
											ID: to.StringPtr("<id>"),
										},
									},
								}},
						},
					}},
				EnableBgpRouteTranslationForNat: to.BoolPtr(false),
				IsRoutingPreferenceInternet:     to.BoolPtr(false),
				NatRules: []*armnetwork.VPNGatewayNatRule{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.VPNGatewayNatRuleProperties{
							Type: armnetwork.VPNNatRuleType("Static").ToPtr(),
							ExternalMappings: []*armnetwork.VPNNatRuleMapping{
								{
									AddressSpace: to.StringPtr("<address-space>"),
								}},
							InternalMappings: []*armnetwork.VPNNatRuleMapping{
								{
									AddressSpace: to.StringPtr("<address-space>"),
								}},
							IPConfigurationID: to.StringPtr("<ipconfiguration-id>"),
							Mode:              armnetwork.VPNNatRuleMode("EgressSnat").ToPtr(),
						},
					}},
				VirtualHub: &armnetwork.SubResource{
					ID: to.StringPtr("<id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VPNGatewaysClientCreateOrUpdateResult)
}
```
