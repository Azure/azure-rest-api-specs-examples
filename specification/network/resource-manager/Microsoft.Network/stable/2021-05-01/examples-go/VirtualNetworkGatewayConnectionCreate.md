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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayConnectionCreate.json
func ExampleVirtualNetworkGatewayConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewVirtualNetworkGatewayConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-network-gateway-connection-name>",
		armnetwork.VirtualNetworkGatewayConnection{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.VirtualNetworkGatewayConnectionPropertiesFormat{
				ConnectionMode:     armnetwork.VirtualNetworkGatewayConnectionMode("Default").ToPtr(),
				ConnectionProtocol: armnetwork.VirtualNetworkGatewayConnectionProtocol("IKEv2").ToPtr(),
				ConnectionType:     armnetwork.VirtualNetworkGatewayConnectionType("IPsec").ToPtr(),
				DpdTimeoutSeconds:  to.Int32Ptr(30),
				EgressNatRules: []*armnetwork.SubResource{
					{
						ID: to.StringPtr("<id>"),
					}},
				EnableBgp: to.BoolPtr(false),
				IngressNatRules: []*armnetwork.SubResource{
					{
						ID: to.StringPtr("<id>"),
					}},
				IPSecPolicies: []*armnetwork.IPSecPolicy{},
				LocalNetworkGateway2: &armnetwork.LocalNetworkGateway{
					ID:       to.StringPtr("<id>"),
					Location: to.StringPtr("<location>"),
					Tags:     map[string]*string{},
					Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
						GatewayIPAddress: to.StringPtr("<gateway-ipaddress>"),
						LocalNetworkAddressSpace: &armnetwork.AddressSpace{
							AddressPrefixes: []*string{
								to.StringPtr("10.1.0.0/16")},
						},
					},
				},
				RoutingWeight:                  to.Int32Ptr(0),
				SharedKey:                      to.StringPtr("<shared-key>"),
				TrafficSelectorPolicies:        []*armnetwork.TrafficSelectorPolicy{},
				UsePolicyBasedTrafficSelectors: to.BoolPtr(false),
				VirtualNetworkGateway1: &armnetwork.VirtualNetworkGateway{
					ID:       to.StringPtr("<id>"),
					Location: to.StringPtr("<location>"),
					Tags:     map[string]*string{},
					Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
						Active: to.BoolPtr(false),
						BgpSettings: &armnetwork.BgpSettings{
							Asn:               to.Int64Ptr(65514),
							BgpPeeringAddress: to.StringPtr("<bgp-peering-address>"),
							PeerWeight:        to.Int32Ptr(0),
						},
						EnableBgp:   to.BoolPtr(false),
						GatewayType: armnetwork.VirtualNetworkGatewayType("Vpn").ToPtr(),
						IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
							{
								ID:   to.StringPtr("<id>"),
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
						SKU: &armnetwork.VirtualNetworkGatewaySKU{
							Name: armnetwork.VirtualNetworkGatewaySKUName("VpnGw1").ToPtr(),
							Tier: armnetwork.VirtualNetworkGatewaySKUTier("VpnGw1").ToPtr(),
						},
						VPNType: armnetwork.VPNType("RouteBased").ToPtr(),
					},
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
	log.Printf("Response result: %#v\n", res.VirtualNetworkGatewayConnectionsClientCreateOrUpdateResult)
}
```
