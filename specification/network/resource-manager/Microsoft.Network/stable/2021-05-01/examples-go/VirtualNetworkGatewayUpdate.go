package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayUpdate.json
func ExampleVirtualNetworkGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualNetworkGatewaysClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-network-gateway-name>",
		armnetwork.VirtualNetworkGateway{
			Location: to.Ptr("<location>"),
			Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
				Active: to.Ptr(false),
				BgpSettings: &armnetwork.BgpSettings{
					Asn:               to.Ptr[int64](65515),
					BgpPeeringAddress: to.Ptr("<bgp-peering-address>"),
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
				NatRules: []*armnetwork.VirtualNetworkGatewayNatRule{
					{
						ID:   to.Ptr("<id>"),
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
							Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
							ExternalMappings: []*armnetwork.VPNNatRuleMapping{
								{
									AddressSpace: to.Ptr("<address-space>"),
								}},
							InternalMappings: []*armnetwork.VPNNatRuleMapping{
								{
									AddressSpace: to.Ptr("<address-space>"),
								}},
							IPConfigurationID: to.Ptr("<ipconfiguration-id>"),
							Mode:              to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
						},
					},
					{
						ID:   to.Ptr("<id>"),
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
							Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
							ExternalMappings: []*armnetwork.VPNNatRuleMapping{
								{
									AddressSpace: to.Ptr("<address-space>"),
								}},
							InternalMappings: []*armnetwork.VPNNatRuleMapping{
								{
									AddressSpace: to.Ptr("<address-space>"),
								}},
							IPConfigurationID: to.Ptr("<ipconfiguration-id>"),
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
							RadiusServerAddress: to.Ptr("<radius-server-address>"),
							RadiusServerScore:   to.Ptr[int64](20),
							RadiusServerSecret:  to.Ptr("<radius-server-secret>"),
						}},
					VPNClientProtocols: []*armnetwork.VPNClientProtocol{
						to.Ptr(armnetwork.VPNClientProtocolOpenVPN)},
					VPNClientRevokedCertificates: []*armnetwork.VPNClientRevokedCertificate{},
					VPNClientRootCertificates:    []*armnetwork.VPNClientRootCertificate{},
				},
				VPNType: to.Ptr(armnetwork.VPNTypeRouteBased),
			},
		},
		&armnetwork.VirtualNetworkGatewaysClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
