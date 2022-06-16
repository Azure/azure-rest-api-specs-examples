package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnGatewayPut.json
func ExampleVPNGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewVPNGatewaysClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<gateway-name>",
		armnetwork.VPNGateway{
			Location: to.Ptr("<location>"),
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
							IPConfigurationID: to.Ptr("<ipconfiguration-id>"),
						},
						{
							CustomBgpIPAddresses: []*string{
								to.Ptr("169.254.21.10")},
							IPConfigurationID: to.Ptr("<ipconfiguration-id>"),
						}},
					PeerWeight: to.Ptr[int32](0),
				},
				Connections: []*armnetwork.VPNConnection{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.VPNConnectionProperties{
							RemoteVPNSite: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							VPNLinkConnections: []*armnetwork.VPNSiteLinkConnection{
								{
									Name: to.Ptr("<name>"),
									Properties: &armnetwork.VPNSiteLinkConnectionProperties{
										ConnectionBandwidth: to.Ptr[int32](200),
										EgressNatRules: []*armnetwork.SubResource{
											{
												ID: to.Ptr("<id>"),
											}},
										SharedKey:                 to.Ptr("<shared-key>"),
										VPNConnectionProtocolType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
										VPNSiteLink: &armnetwork.SubResource{
											ID: to.Ptr("<id>"),
										},
									},
								}},
						},
					}},
				EnableBgpRouteTranslationForNat: to.Ptr(false),
				IsRoutingPreferenceInternet:     to.Ptr(false),
				NatRules: []*armnetwork.VPNGatewayNatRule{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.VPNGatewayNatRuleProperties{
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
					}},
				VirtualHub: &armnetwork.SubResource{
					ID: to.Ptr("<id>"),
				},
			},
		},
		&armnetwork.VPNGatewaysClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
