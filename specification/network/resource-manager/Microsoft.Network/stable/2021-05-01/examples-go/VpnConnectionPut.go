package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnConnectionPut.json
func ExampleVPNConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewVPNConnectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<gateway-name>",
		"<connection-name>",
		armnetwork.VPNConnection{
			Properties: &armnetwork.VPNConnectionProperties{
				RemoteVPNSite: &armnetwork.SubResource{
					ID: to.Ptr("<id>"),
				},
				TrafficSelectorPolicies: []*armnetwork.TrafficSelectorPolicy{},
				VPNLinkConnections: []*armnetwork.VPNSiteLinkConnection{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.VPNSiteLinkConnectionProperties{
							ConnectionBandwidth:            to.Ptr[int32](200),
							SharedKey:                      to.Ptr("<shared-key>"),
							UsePolicyBasedTrafficSelectors: to.Ptr(false),
							VPNConnectionProtocolType:      to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
							VPNLinkConnectionMode:          to.Ptr(armnetwork.VPNLinkConnectionModeDefault),
							VPNSiteLink: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
						},
					}},
			},
		},
		&armnetwork.VPNConnectionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
