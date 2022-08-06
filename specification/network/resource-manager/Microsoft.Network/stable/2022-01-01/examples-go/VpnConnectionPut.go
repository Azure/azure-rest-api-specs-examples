package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnConnectionPut.json
func ExampleVPNConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewVPNConnectionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "gateway1", "vpnConnection1", armnetwork.VPNConnection{
		Properties: &armnetwork.VPNConnectionProperties{
			RemoteVPNSite: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"),
			},
			TrafficSelectorPolicies: []*armnetwork.TrafficSelectorPolicy{},
			VPNLinkConnections: []*armnetwork.VPNSiteLinkConnection{
				{
					Name: to.Ptr("Connection-Link1"),
					Properties: &armnetwork.VPNSiteLinkConnectionProperties{
						ConnectionBandwidth:            to.Ptr[int32](200),
						SharedKey:                      to.Ptr("key"),
						UsePolicyBasedTrafficSelectors: to.Ptr(false),
						VPNConnectionProtocolType:      to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
						VPNLinkConnectionMode:          to.Ptr(armnetwork.VPNLinkConnectionModeDefault),
						VPNSiteLink: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/siteLink1"),
						},
					},
				}},
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
