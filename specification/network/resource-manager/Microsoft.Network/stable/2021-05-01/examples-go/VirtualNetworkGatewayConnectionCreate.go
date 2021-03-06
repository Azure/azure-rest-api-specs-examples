package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayConnectionCreate.json
func ExampleVirtualNetworkGatewayConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualNetworkGatewayConnectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-network-gateway-connection-name>",
		armnetwork.VirtualNetworkGatewayConnection{
			Location: to.Ptr("<location>"),
			Properties: &armnetwork.VirtualNetworkGatewayConnectionPropertiesFormat{
				ConnectionMode:     to.Ptr(armnetwork.VirtualNetworkGatewayConnectionModeDefault),
				ConnectionProtocol: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
				ConnectionType:     to.Ptr(armnetwork.VirtualNetworkGatewayConnectionTypeIPsec),
				DpdTimeoutSeconds:  to.Ptr[int32](30),
				EgressNatRules: []*armnetwork.SubResource{
					{
						ID: to.Ptr("<id>"),
					}},
				EnableBgp: to.Ptr(false),
				IngressNatRules: []*armnetwork.SubResource{
					{
						ID: to.Ptr("<id>"),
					}},
				IPSecPolicies: []*armnetwork.IPSecPolicy{},
				LocalNetworkGateway2: &armnetwork.LocalNetworkGateway{
					ID:       to.Ptr("<id>"),
					Location: to.Ptr("<location>"),
					Tags:     map[string]*string{},
					Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
						GatewayIPAddress: to.Ptr("<gateway-ipaddress>"),
						LocalNetworkAddressSpace: &armnetwork.AddressSpace{
							AddressPrefixes: []*string{
								to.Ptr("10.1.0.0/16")},
						},
					},
				},
				RoutingWeight:                  to.Ptr[int32](0),
				SharedKey:                      to.Ptr("<shared-key>"),
				TrafficSelectorPolicies:        []*armnetwork.TrafficSelectorPolicy{},
				UsePolicyBasedTrafficSelectors: to.Ptr(false),
				VirtualNetworkGateway1: &armnetwork.VirtualNetworkGateway{
					ID:       to.Ptr("<id>"),
					Location: to.Ptr("<location>"),
					Tags:     map[string]*string{},
					Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
						Active: to.Ptr(false),
						BgpSettings: &armnetwork.BgpSettings{
							Asn:               to.Ptr[int64](65514),
							BgpPeeringAddress: to.Ptr("<bgp-peering-address>"),
							PeerWeight:        to.Ptr[int32](0),
						},
						EnableBgp:   to.Ptr(false),
						GatewayType: to.Ptr(armnetwork.VirtualNetworkGatewayTypeVPN),
						IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
							{
								ID:   to.Ptr("<id>"),
								Name: to.Ptr("<name>"),
								Properties: &armnetwork.VirtualNetworkGatewayIPConfigurationPropertiesFormat{
									PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
									PublicIPAddress: &armnetwork.SubResource{
										ID: to.Ptr("<id>"),
									},
									Subnet: &armnetwork.SubResource{
										ID: to.Ptr("<id>"),
									},
								},
							}},
						SKU: &armnetwork.VirtualNetworkGatewaySKU{
							Name: to.Ptr(armnetwork.VirtualNetworkGatewaySKUNameVPNGw1),
							Tier: to.Ptr(armnetwork.VirtualNetworkGatewaySKUTierVPNGw1),
						},
						VPNType: to.Ptr(armnetwork.VPNTypeRouteBased),
					},
				},
			},
		},
		&armnetwork.VirtualNetworkGatewayConnectionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
