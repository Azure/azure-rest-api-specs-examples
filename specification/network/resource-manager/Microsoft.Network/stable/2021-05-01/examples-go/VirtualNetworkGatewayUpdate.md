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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayUpdate.json
func ExampleVirtualNetworkGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewVirtualNetworkGatewaysClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-network-gateway-name>",
		armnetwork.VirtualNetworkGateway{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
				Active: to.BoolPtr(false),
				BgpSettings: &armnetwork.BgpSettings{
					Asn:               to.Int64Ptr(65515),
					BgpPeeringAddress: to.StringPtr("<bgp-peering-address>"),
					PeerWeight:        to.Int32Ptr(0),
				},
				CustomRoutes: &armnetwork.AddressSpace{
					AddressPrefixes: []*string{
						to.StringPtr("101.168.0.6/32")},
				},
				DisableIPSecReplayProtection:    to.BoolPtr(false),
				EnableBgp:                       to.BoolPtr(false),
				EnableBgpRouteTranslationForNat: to.BoolPtr(false),
				EnableDNSForwarding:             to.BoolPtr(true),
				GatewayType:                     armnetwork.VirtualNetworkGatewayType("Vpn").ToPtr(),
				IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.VirtualNetworkGatewayIPConfigurationPropertiesFormat{
							PrivateIPAllocationMethod: armnetwork.IPAllocationMethod("Dynamic").ToPtr(),
							PublicIPAddress: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							Subnet: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
						},
					}},
				NatRules: []*armnetwork.VirtualNetworkGatewayNatRule{
					{
						ID:   to.StringPtr("<id>"),
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
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
					},
					{
						ID:   to.StringPtr("<id>"),
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
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
							Mode:              armnetwork.VPNNatRuleMode("IngressSnat").ToPtr(),
						},
					}},
				SKU: &armnetwork.VirtualNetworkGatewaySKU{
					Name: armnetwork.VirtualNetworkGatewaySKUName("VpnGw1").ToPtr(),
					Tier: armnetwork.VirtualNetworkGatewaySKUTier("VpnGw1").ToPtr(),
				},
				VPNClientConfiguration: &armnetwork.VPNClientConfiguration{
					RadiusServers: []*armnetwork.RadiusServer{
						{
							RadiusServerAddress: to.StringPtr("<radius-server-address>"),
							RadiusServerScore:   to.Int64Ptr(20),
							RadiusServerSecret:  to.StringPtr("<radius-server-secret>"),
						}},
					VPNClientProtocols: []*armnetwork.VPNClientProtocol{
						armnetwork.VPNClientProtocol("OpenVPN").ToPtr()},
					VPNClientRevokedCertificates: []*armnetwork.VPNClientRevokedCertificate{},
					VPNClientRootCertificates:    []*armnetwork.VPNClientRootCertificate{},
				},
				VPNType: armnetwork.VPNType("RouteBased").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.VirtualNetworkGatewaysClientCreateOrUpdateResult)
}
```
