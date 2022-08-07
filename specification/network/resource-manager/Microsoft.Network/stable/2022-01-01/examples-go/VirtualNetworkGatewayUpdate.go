package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkGatewayUpdate.json
func ExampleVirtualNetworkGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualNetworkGatewaysClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "vpngw", armnetwork.VirtualNetworkGateway{
		Location: to.Ptr("centralus"),
		Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
			Active: to.Ptr(false),
			BgpSettings: &armnetwork.BgpSettings{
				Asn:               to.Ptr[int64](65515),
				BgpPeeringAddress: to.Ptr("10.0.1.30"),
				PeerWeight:        to.Ptr[int32](0),
			},
			CustomRoutes: &armnetwork.AddressSpace{
				AddressPrefixes: []*string{
					to.Ptr("101.168.0.6/32")},
			},
			DisableIPSecReplayProtection:    to.Ptr(false),
			EnableBgp:                       to.Ptr(false),
			EnableBgpRouteTranslationForNat: to.Ptr(false),
			EnableDNSForwarding:             to.Ptr(true),
			GatewayType:                     to.Ptr(armnetwork.VirtualNetworkGatewayTypeVPN),
			IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
				{
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
			NatRules: []*armnetwork.VirtualNetworkGatewayNatRule{
				{
					ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1"),
					Name: to.Ptr("natRule1"),
					Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
						Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
						ExternalMappings: []*armnetwork.VPNNatRuleMapping{
							{
								AddressSpace: to.Ptr("50.0.0.0/24"),
							}},
						InternalMappings: []*armnetwork.VPNNatRuleMapping{
							{
								AddressSpace: to.Ptr("10.10.0.0/24"),
							}},
						IPConfigurationID: to.Ptr(""),
						Mode:              to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
					},
				},
				{
					ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2"),
					Name: to.Ptr("natRule2"),
					Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
						Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
						ExternalMappings: []*armnetwork.VPNNatRuleMapping{
							{
								AddressSpace: to.Ptr("30.0.0.0/24"),
							}},
						InternalMappings: []*armnetwork.VPNNatRuleMapping{
							{
								AddressSpace: to.Ptr("20.10.0.0/24"),
							}},
						IPConfigurationID: to.Ptr(""),
						Mode:              to.Ptr(armnetwork.VPNNatRuleModeIngressSnat),
					},
				}},
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
					}},
				VPNClientProtocols: []*armnetwork.VPNClientProtocol{
					to.Ptr(armnetwork.VPNClientProtocolOpenVPN)},
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
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
