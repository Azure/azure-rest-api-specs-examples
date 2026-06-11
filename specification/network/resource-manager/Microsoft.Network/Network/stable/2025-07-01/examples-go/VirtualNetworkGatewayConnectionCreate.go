package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualNetworkGatewayConnectionCreate.json
func ExampleVirtualNetworkGatewayConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkGatewayConnectionsClient().BeginCreateOrUpdate(ctx, "rg1", "connS2S", armnetwork.VirtualNetworkGatewayConnection{
		Location: to.Ptr("centralus"),
		Properties: &armnetwork.VirtualNetworkGatewayConnectionPropertiesFormat{
			AuthenticationType: to.Ptr(armnetwork.ConnectionAuthenticationTypeCertificate),
			CertificateAuthentication: &armnetwork.CertificateAuthentication{
				InboundAuthCertificateChain: []*string{
					to.Ptr("MIIC+TCCAeGgAwIBAgIQFOJUqDaxV5xJcKpTKO..."),
					to.Ptr("MIIC+TCCAeGgAwIBAgIQPJerInitNblK7yBgkqh..."),
				},
				InboundAuthCertificateSubjectName: to.Ptr("CN=rootCert.com"),
				OutboundAuthCertificate:           to.Ptr("https://customerKv.vault.azure.net/Certificates/outBoundcert/Version"),
			},
			ConnectionMode:     to.Ptr(armnetwork.VirtualNetworkGatewayConnectionModeDefault),
			ConnectionProtocol: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
			ConnectionType:     to.Ptr(armnetwork.VirtualNetworkGatewayConnectionTypeIPsec),
			DpdTimeoutSeconds:  to.Ptr[int32](30),
			EgressNatRules: []*armnetwork.SubResource{
				{
					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2"),
				},
			},
			EnableBgp: to.Ptr(false),
			GatewayCustomBgpIPAddresses: []*armnetwork.GatewayCustomBgpIPAddressIPConfiguration{
				{
					CustomBgpIPAddress: to.Ptr("169.254.21.1"),
					IPConfigurationID:  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/default"),
				},
				{
					CustomBgpIPAddress: to.Ptr("169.254.21.3"),
					IPConfigurationID:  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/ActiveActive"),
				},
			},
			IngressNatRules: []*armnetwork.SubResource{
				{
					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1"),
				},
			},
			IPSecPolicies: []*armnetwork.IPSecPolicy{},
			LocalNetworkGateway2: &armnetwork.LocalNetworkGateway{
				ID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw"),
				Location: to.Ptr("centralus"),
				Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
					GatewayIPAddress: to.Ptr("x.x.x.x"),
					LocalNetworkAddressSpace: &armnetwork.AddressSpace{
						AddressPrefixes: []*string{
							to.Ptr("10.1.0.0/16"),
						},
					},
				},
				Tags: map[string]*string{},
			},
			RoutingWeight: to.Ptr[int32](0),
			SharedKey:     to.Ptr("Abc123"),
			RoutingConfiguration: &armnetwork.RoutingConfiguration{
				InboundRouteMap: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
				},
				OutboundRouteMap: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
				},
			},
			TrafficSelectorPolicies: []*armnetwork.TrafficSelectorPolicy{},
			TunnelProperties: []*armnetwork.VirtualNetworkGatewayConnectionTunnelProperties{
				{
					BgpPeeringAddress: to.Ptr("10.78.1.17"),
					TunnelIPAddress:   to.Ptr("10.78.1.5"),
				},
				{
					BgpPeeringAddress: to.Ptr("10.78.1.20"),
					TunnelIPAddress:   to.Ptr("10.78.1.7"),
				},
			},
			UsePolicyBasedTrafficSelectors: to.Ptr(false),
			VirtualNetworkGateway1: &armnetwork.VirtualNetworkGateway{
				ID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw"),
				Location: to.Ptr("centralus"),
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
							Name: to.Ptr("gwipconfig1"),
							ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/gwipconfig1"),
							Properties: &armnetwork.VirtualNetworkGatewayIPConfigurationPropertiesFormat{
								PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
								PublicIPAddress: &armnetwork.SubResource{
									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"),
								},
								Subnet: &armnetwork.SubResource{
									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"),
								},
							},
						},
					},
					SKU: &armnetwork.VirtualNetworkGatewaySKU{
						Name: to.Ptr(armnetwork.VirtualNetworkGatewaySKUNameVPNGw1),
						Tier: to.Ptr(armnetwork.VirtualNetworkGatewaySKUTierVPNGw1),
					},
					VPNType: to.Ptr(armnetwork.VPNTypeRouteBased),
				},
				Tags: map[string]*string{},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualNetworkGatewayConnectionsClientCreateOrUpdateResponse{
	// 	VirtualNetworkGatewayConnection: armnetwork.VirtualNetworkGatewayConnection{
	// 		Name: to.Ptr("connS2S"),
	// 		Type: to.Ptr("Microsoft.Network/connections"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/connections/connS2S"),
	// 		Location: to.Ptr("centralus"),
	// 		Properties: &armnetwork.VirtualNetworkGatewayConnectionPropertiesFormat{
	// 			AuthenticationType: to.Ptr(armnetwork.ConnectionAuthenticationTypeCertificate),
	// 			CertificateAuthentication: &armnetwork.CertificateAuthentication{
	// 				InboundAuthCertificateChain: []*string{
	// 					to.Ptr("MIIC+TCCAeGgAwIBAgIQFOJUqDaxV5xJcKpTKO..."),
	// 					to.Ptr("MIIC+TCCAeGgAwIBAgIQPJerInitNblK7yBgkqh..."),
	// 				},
	// 				InboundAuthCertificateSubjectName: to.Ptr("CN=rootCert.com"),
	// 				OutboundAuthCertificate: to.Ptr("https://customerKv.vault.azure.net/Certificates/outBoundcert/Version"),
	// 			},
	// 			ConnectionMode: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionModeDefault),
	// 			ConnectionProtocol: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
	// 			ConnectionType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionTypeIPsec),
	// 			DpdTimeoutSeconds: to.Ptr[int32](30),
	// 			EgressBytesTransferred: to.Ptr[int64](0),
	// 			EgressNatRules: []*armnetwork.SubResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2"),
	// 				},
	// 			},
	// 			EnableBgp: to.Ptr(false),
	// 			GatewayCustomBgpIPAddresses: []*armnetwork.GatewayCustomBgpIPAddressIPConfiguration{
	// 				{
	// 					CustomBgpIPAddress: to.Ptr("169.254.21.1"),
	// 					IPConfigurationID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/default"),
	// 				},
	// 				{
	// 					CustomBgpIPAddress: to.Ptr("169.254.21.3"),
	// 					IPConfigurationID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/ActiveActive"),
	// 				},
	// 			},
	// 			IngressBytesTransferred: to.Ptr[int64](0),
	// 			IngressNatRules: []*armnetwork.SubResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1"),
	// 				},
	// 			},
	// 			IPSecPolicies: []*armnetwork.IPSecPolicy{
	// 			},
	// 			LocalNetworkGateway2: &armnetwork.LocalNetworkGateway{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw"),
	// 				Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			RoutingConfiguration: &armnetwork.RoutingConfiguration{
	// 				InboundRouteMap: &armnetwork.SubResource{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
	// 				},
	// 				OutboundRouteMap: &armnetwork.SubResource{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
	// 				},
	// 			},
	// 			RoutingWeight: to.Ptr[int32](0),
	// 			TunnelProperties: []*armnetwork.VirtualNetworkGatewayConnectionTunnelProperties{
	// 				{
	// 					BgpPeeringAddress: to.Ptr("10.78.1.17"),
	// 					TunnelIPAddress: to.Ptr("10.78.1.5"),
	// 				},
	// 				{
	// 					BgpPeeringAddress: to.Ptr("10.78.1.20"),
	// 					TunnelIPAddress: to.Ptr("10.78.1.7"),
	// 				},
	// 			},
	// 			UseLocalAzureIPAddress: to.Ptr(false),
	// 			UsePolicyBasedTrafficSelectors: to.Ptr(false),
	// 			VirtualNetworkGateway1: &armnetwork.VirtualNetworkGateway{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw"),
	// 				Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
