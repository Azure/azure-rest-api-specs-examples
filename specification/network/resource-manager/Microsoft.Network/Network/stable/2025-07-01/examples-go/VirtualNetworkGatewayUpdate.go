package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualNetworkGatewayUpdate.json
func ExampleVirtualNetworkGatewaysClient_BeginCreateOrUpdate_updateVirtualNetworkGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkGatewaysClient().BeginCreateOrUpdate(ctx, "rg1", "vpngw", armnetwork.VirtualNetworkGateway{
		Identity: &armnetwork.ManagedServiceIdentity{
			Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armnetwork.ManagedServiceIdentityUserAssignedIdentities{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
			},
		},
		Location: to.Ptr("centralus"),
		Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
			Active:                 to.Ptr(false),
			AllowRemoteVnetTraffic: to.Ptr(false),
			AllowVirtualWanTraffic: to.Ptr(false),
			BgpSettings: &armnetwork.BgpSettings{
				Asn:               to.Ptr[int64](65515),
				BgpPeeringAddress: to.Ptr("10.0.1.30"),
				PeerWeight:        to.Ptr[int32](0),
			},
			CustomRoutes: &armnetwork.AddressSpace{
				AddressPrefixes: []*string{
					to.Ptr("101.168.0.6/32"),
				},
			},
			DisableIPSecReplayProtection:    to.Ptr(false),
			EnableBgp:                       to.Ptr(false),
			EnableBgpRouteTranslationForNat: to.Ptr(false),
			EnableDNSForwarding:             to.Ptr(true),
			EnableHighBandwidthVPNGateway:   to.Ptr(false),
			GatewayType:                     to.Ptr(armnetwork.VirtualNetworkGatewayTypeVPN),
			IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
				{
					Name: to.Ptr("gwipconfig1"),
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
			NatRules: []*armnetwork.VirtualNetworkGatewayNatRule{
				{
					Name: to.Ptr("natRule1"),
					ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1"),
					Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
						Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
						ExternalMappings: []*armnetwork.VPNNatRuleMapping{
							{
								AddressSpace: to.Ptr("50.0.0.0/24"),
							},
						},
						InternalMappings: []*armnetwork.VPNNatRuleMapping{
							{
								AddressSpace: to.Ptr("10.10.0.0/24"),
							},
						},
						IPConfigurationID: to.Ptr(""),
						Mode:              to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
					},
				},
				{
					Name: to.Ptr("natRule2"),
					ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2"),
					Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
						Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
						ExternalMappings: []*armnetwork.VPNNatRuleMapping{
							{
								AddressSpace: to.Ptr("30.0.0.0/24"),
							},
						},
						InternalMappings: []*armnetwork.VPNNatRuleMapping{
							{
								AddressSpace: to.Ptr("20.10.0.0/24"),
							},
						},
						IPConfigurationID: to.Ptr(""),
						Mode:              to.Ptr(armnetwork.VPNNatRuleModeIngressSnat),
					},
				},
			},
			SKU: &armnetwork.VirtualNetworkGatewaySKU{
				Name: to.Ptr(armnetwork.VirtualNetworkGatewaySKUNameVPNGw1),
				Tier: to.Ptr(armnetwork.VirtualNetworkGatewaySKUTierVPNGw1),
			},
			VPNClientConfiguration: &armnetwork.VPNClientConfiguration{
				RadiusServers: []*armnetwork.RadiusServer{
					{
						RadiusServerAddress: to.Ptr("10.2.0.0"),
						RadiusServerScore:   to.Ptr[int64](20),
						RadiusServerSecret:  to.Ptr("radiusServerSecret"),
					},
				},
				VPNClientProtocols: []*armnetwork.VPNClientProtocol{
					to.Ptr(armnetwork.VPNClientProtocolOpenVPN),
				},
				VPNClientRevokedCertificates: []*armnetwork.VPNClientRevokedCertificate{},
				VPNClientRootCertificates:    []*armnetwork.VPNClientRootCertificate{},
			},
			VPNType: to.Ptr(armnetwork.VPNTypeRouteBased),
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
	// res = armnetwork.VirtualNetworkGatewaysClientCreateOrUpdateResponse{
	// 	VirtualNetworkGateway: armnetwork.VirtualNetworkGateway{
	// 		Name: to.Ptr("vpngw"),
	// 		Type: to.Ptr("Microsoft.Network/virtualNetworkGateways"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw"),
	// 		Identity: &armnetwork.ManagedServiceIdentity{
	// 			Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armnetwork.ManagedServiceIdentityUserAssignedIdentities{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armnetwork.ManagedServiceIdentityUserAssignedIdentities{
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("centralus"),
	// 		Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
	// 			Active: to.Ptr(false),
	// 			AllowRemoteVnetTraffic: to.Ptr(false),
	// 			AllowVirtualWanTraffic: to.Ptr(false),
	// 			BgpSettings: &armnetwork.BgpSettings{
	// 				Asn: to.Ptr[int64](65515),
	// 				BgpPeeringAddress: to.Ptr("10.0.1.30"),
	// 				BgpPeeringAddresses: []*armnetwork.IPConfigurationBgpPeeringAddress{
	// 					{
	// 						CustomBgpIPAddresses: []*string{
	// 							to.Ptr("169.254.21.10"),
	// 						},
	// 						DefaultBgpIPAddresses: []*string{
	// 							to.Ptr("10.3.1.254"),
	// 						},
	// 						IPConfigurationID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/gwipconfig1"),
	// 						TunnelIPAddresses: []*string{
	// 							to.Ptr("52.161.10.135"),
	// 						},
	// 					},
	// 				},
	// 				PeerWeight: to.Ptr[int32](0),
	// 			},
	// 			CustomRoutes: &armnetwork.AddressSpace{
	// 				AddressPrefixes: []*string{
	// 					to.Ptr("101.168.0.6/32"),
	// 				},
	// 			},
	// 			DisableIPSecReplayProtection: to.Ptr(false),
	// 			EnableBgp: to.Ptr(false),
	// 			EnableBgpRouteTranslationForNat: to.Ptr(false),
	// 			EnableDNSForwarding: to.Ptr(true),
	// 			EnableHighBandwidthVPNGateway: to.Ptr(false),
	// 			GatewayType: to.Ptr(armnetwork.VirtualNetworkGatewayTypeVPN),
	// 			InboundDNSForwardingEndpoint: to.Ptr("10.0.1.14"),
	// 			IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
	// 				{
	// 					Name: to.Ptr("gwipconfig1"),
	// 					Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/gwipconfig1"),
	// 					Properties: &armnetwork.VirtualNetworkGatewayIPConfigurationPropertiesFormat{
	// 						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						PublicIPAddress: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"),
	// 						},
	// 						Subnet: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			NatRules: []*armnetwork.VirtualNetworkGatewayNatRule{
	// 				{
	// 					Name: to.Ptr("natRule1"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1"),
	// 					Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
	// 						Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
	// 						ExternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 							{
	// 								AddressSpace: to.Ptr("50.0.0.0/24"),
	// 							},
	// 						},
	// 						InternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 							{
	// 								AddressSpace: to.Ptr("10.10.0.0/24"),
	// 							},
	// 						},
	// 						IPConfigurationID: to.Ptr(""),
	// 						Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("natRule2"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2"),
	// 					Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
	// 						Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
	// 						ExternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 							{
	// 								AddressSpace: to.Ptr("30.0.0.0/24"),
	// 							},
	// 						},
	// 						InternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 							{
	// 								AddressSpace: to.Ptr("20.10.0.0/24"),
	// 							},
	// 						},
	// 						IPConfigurationID: to.Ptr(""),
	// 						Mode: to.Ptr(armnetwork.VPNNatRuleModeIngressSnat),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			SKU: &armnetwork.VirtualNetworkGatewaySKU{
	// 				Name: to.Ptr(armnetwork.VirtualNetworkGatewaySKUNameVPNGw1),
	// 				Capacity: to.Ptr[int32](0),
	// 				Tier: to.Ptr(armnetwork.VirtualNetworkGatewaySKUTierVPNGw1),
	// 			},
	// 			VPNClientConfiguration: &armnetwork.VPNClientConfiguration{
	// 				RadiusServers: []*armnetwork.RadiusServer{
	// 					{
	// 						RadiusServerAddress: to.Ptr("10.2.0.0"),
	// 						RadiusServerScore: to.Ptr[int64](20),
	// 					},
	// 				},
	// 				VPNClientProtocols: []*armnetwork.VPNClientProtocol{
	// 					to.Ptr(armnetwork.VPNClientProtocolOpenVPN),
	// 				},
	// 				VPNClientRevokedCertificates: []*armnetwork.VPNClientRevokedCertificate{
	// 				},
	// 				VPNClientRootCertificates: []*armnetwork.VPNClientRootCertificate{
	// 				},
	// 			},
	// 			VPNGatewayGeneration: to.Ptr(armnetwork.VPNGatewayGenerationNone),
	// 			VPNType: to.Ptr(armnetwork.VPNTypeRouteBased),
	// 		},
	// 	},
	// }
}
