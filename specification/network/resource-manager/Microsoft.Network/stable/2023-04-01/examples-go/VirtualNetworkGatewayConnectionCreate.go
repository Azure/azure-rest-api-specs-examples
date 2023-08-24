package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkGatewayConnectionCreate.json
func ExampleVirtualNetworkGatewayConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkGatewayConnectionsClient().BeginCreateOrUpdate(ctx, "rg1", "connS2S", armnetwork.VirtualNetworkGatewayConnection{
		Location: to.Ptr("centralus"),
		Properties: &armnetwork.VirtualNetworkGatewayConnectionPropertiesFormat{
			ConnectionMode:     to.Ptr(armnetwork.VirtualNetworkGatewayConnectionModeDefault),
			ConnectionProtocol: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
			ConnectionType:     to.Ptr(armnetwork.VirtualNetworkGatewayConnectionTypeIPsec),
			DpdTimeoutSeconds:  to.Ptr[int32](30),
			EgressNatRules: []*armnetwork.SubResource{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2"),
				}},
			EnableBgp: to.Ptr(false),
			GatewayCustomBgpIPAddresses: []*armnetwork.GatewayCustomBgpIPAddressIPConfiguration{
				{
					CustomBgpIPAddress: to.Ptr("169.254.21.1"),
					IPConfigurationID:  to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/default"),
				},
				{
					CustomBgpIPAddress: to.Ptr("169.254.21.3"),
					IPConfigurationID:  to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/ActiveActive"),
				}},
			IngressNatRules: []*armnetwork.SubResource{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1"),
				}},
			IPSecPolicies: []*armnetwork.IPSecPolicy{},
			LocalNetworkGateway2: &armnetwork.LocalNetworkGateway{
				ID:       to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw"),
				Location: to.Ptr("centralus"),
				Tags:     map[string]*string{},
				Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
					GatewayIPAddress: to.Ptr("x.x.x.x"),
					LocalNetworkAddressSpace: &armnetwork.AddressSpace{
						AddressPrefixes: []*string{
							to.Ptr("10.1.0.0/16")},
					},
				},
			},
			RoutingWeight:                  to.Ptr[int32](0),
			SharedKey:                      to.Ptr("Abc123"),
			TrafficSelectorPolicies:        []*armnetwork.TrafficSelectorPolicy{},
			UsePolicyBasedTrafficSelectors: to.Ptr(false),
			VirtualNetworkGateway1: &armnetwork.VirtualNetworkGateway{
				ID:       to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw"),
				Location: to.Ptr("centralus"),
				Tags:     map[string]*string{},
				Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
					Active: to.Ptr(false),
					BgpSettings: &armnetwork.BgpSettings{
						Asn:               to.Ptr[int64](65514),
						BgpPeeringAddress: to.Ptr("10.0.1.30"),
						PeerWeight:        to.Ptr[int32](0),
					},
					EnableBgp:   to.Ptr(false),
					GatewayType: to.Ptr(armnetwork.VirtualNetworkGatewayTypeVPN),
					IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
						{
							ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/gwipconfig1"),
							Name: to.Ptr("gwipconfig1"),
							Properties: &armnetwork.VirtualNetworkGatewayIPConfigurationPropertiesFormat{
								PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
								PublicIPAddress: &armnetwork.SubResource{
									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"),
								},
								Subnet: &armnetwork.SubResource{
									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"),
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
	// res.VirtualNetworkGatewayConnection = armnetwork.VirtualNetworkGatewayConnection{
	// 	Name: to.Ptr("connS2S"),
	// 	Type: to.Ptr("Microsoft.Network/connections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/connections/connS2S"),
	// 	Location: to.Ptr("centralus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armnetwork.VirtualNetworkGatewayConnectionPropertiesFormat{
	// 		ConnectionMode: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionModeDefault),
	// 		ConnectionProtocol: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
	// 		ConnectionType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionTypeIPsec),
	// 		DpdTimeoutSeconds: to.Ptr[int32](30),
	// 		EgressBytesTransferred: to.Ptr[int64](0),
	// 		EgressNatRules: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2"),
	// 		}},
	// 		EnableBgp: to.Ptr(false),
	// 		GatewayCustomBgpIPAddresses: []*armnetwork.GatewayCustomBgpIPAddressIPConfiguration{
	// 			{
	// 				CustomBgpIPAddress: to.Ptr("169.254.21.1"),
	// 				IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/default"),
	// 			},
	// 			{
	// 				CustomBgpIPAddress: to.Ptr("169.254.21.3"),
	// 				IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/ActiveActive"),
	// 		}},
	// 		IngressBytesTransferred: to.Ptr[int64](0),
	// 		IngressNatRules: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1"),
	// 		}},
	// 		IPSecPolicies: []*armnetwork.IPSecPolicy{
	// 		},
	// 		LocalNetworkGateway2: &armnetwork.LocalNetworkGateway{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw"),
	// 			Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		RoutingWeight: to.Ptr[int32](0),
	// 		SharedKey: to.Ptr("Abc123"),
	// 		UseLocalAzureIPAddress: to.Ptr(false),
	// 		UsePolicyBasedTrafficSelectors: to.Ptr(false),
	// 		VirtualNetworkGateway1: &armnetwork.VirtualNetworkGateway{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw"),
	// 			Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
	// 			},
	// 		},
	// 	},
	// }
}
