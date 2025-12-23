package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VpnConnectionGet.json
func ExampleVPNConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVPNConnectionsClient().Get(ctx, "rg1", "gateway1", "vpnConnection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VPNConnection = armnetwork.VPNConnection{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1"),
	// 	Name: to.Ptr("vpnConnection1"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VPNConnectionProperties{
	// 		EgressBytesTransferred: to.Ptr[int64](0),
	// 		EnableInternetSecurity: to.Ptr(false),
	// 		IngressBytesTransferred: to.Ptr[int64](0),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RemoteVPNSite: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"),
	// 		},
	// 		RoutingConfiguration: &armnetwork.RoutingConfiguration{
	// 			AssociatedRouteTable: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
	// 			},
	// 			InboundRouteMap: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
	// 			},
	// 			OutboundRouteMap: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
	// 			},
	// 			PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
	// 				IDs: []*armnetwork.SubResource{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable2"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable3"),
	// 				}},
	// 				Labels: []*string{
	// 					to.Ptr("label1"),
	// 					to.Ptr("label2")},
	// 				},
	// 			},
	// 			TrafficSelectorPolicies: []*armnetwork.TrafficSelectorPolicy{
	// 			},
	// 			VPNLinkConnections: []*armnetwork.VPNSiteLinkConnection{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/VpnSiteLinkConnections/Connection-Link1"),
	// 					Name: to.Ptr("Connection-Link1"),
	// 					Type: to.Ptr("Microsoft.Network/vpnGateways/vpnConnections/VpnSiteLinkConnections"),
	// 					Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 					Properties: &armnetwork.VPNSiteLinkConnectionProperties{
	// 						ConnectionBandwidth: to.Ptr[int32](200),
	// 						EgressBytesTransferred: to.Ptr[int64](0),
	// 						EnableBgp: to.Ptr(false),
	// 						EnableRateLimiting: to.Ptr(false),
	// 						IngressBytesTransferred: to.Ptr[int64](0),
	// 						IPSecPolicies: []*armnetwork.IPSecPolicy{
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						RoutingWeight: to.Ptr[int32](0),
	// 						UseLocalAzureIPAddress: to.Ptr(false),
	// 						UsePolicyBasedTrafficSelectors: to.Ptr(false),
	// 						VPNConnectionProtocolType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
	// 						VPNLinkConnectionMode: to.Ptr(armnetwork.VPNLinkConnectionModeResponderOnly),
	// 						VPNSiteLink: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/siteLink1"),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/VpnSiteLinkConnections/Connection-Link2"),
	// 					Name: to.Ptr("Connection-Link2"),
	// 					Type: to.Ptr("Microsoft.Network/vpnGateways/vpnConnections/VpnSiteLinkConnections"),
	// 					Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 					Properties: &armnetwork.VPNSiteLinkConnectionProperties{
	// 						ConnectionBandwidth: to.Ptr[int32](200),
	// 						EgressBytesTransferred: to.Ptr[int64](0),
	// 						EnableBgp: to.Ptr(false),
	// 						EnableRateLimiting: to.Ptr(false),
	// 						IngressBytesTransferred: to.Ptr[int64](0),
	// 						IPSecPolicies: []*armnetwork.IPSecPolicy{
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						RoutingWeight: to.Ptr[int32](0),
	// 						UseLocalAzureIPAddress: to.Ptr(false),
	// 						UsePolicyBasedTrafficSelectors: to.Ptr(false),
	// 						VPNConnectionProtocolType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionProtocolIKEv2),
	// 						VPNLinkConnectionMode: to.Ptr(armnetwork.VPNLinkConnectionModeInitiatorOnly),
	// 						VPNSiteLink: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/siteLink2"),
	// 						},
	// 					},
	// 			}},
	// 		},
	// 	}
}
