package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VpnGatewayReset.json
func ExampleVPNGatewaysClient_BeginReset() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVPNGatewaysClient().BeginReset(ctx, "rg1", "vpngw", &armnetwork.VPNGatewaysClientBeginResetOptions{IPConfigurationID: nil})
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
	// res.VPNGateway = armnetwork.VPNGateway{
	// 	Name: to.Ptr("vpngw"),
	// 	Type: to.Ptr("Microsoft.Network/vpnGateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/vpngw"),
	// 	Location: to.Ptr("West US"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VPNGatewayProperties{
	// 		BgpSettings: &armnetwork.BgpSettings{
	// 			Asn: to.Ptr[int64](65514),
	// 			BgpPeeringAddress: to.Ptr("10.0.1.30"),
	// 			BgpPeeringAddresses: []*armnetwork.IPConfigurationBgpPeeringAddress{
	// 				{
	// 					CustomBgpIPAddresses: []*string{
	// 						to.Ptr("169.254.21.5")},
	// 						DefaultBgpIPAddresses: []*string{
	// 							to.Ptr("10.30.0.4")},
	// 							IPConfigurationID: to.Ptr("Instance0"),
	// 							TunnelIPAddresses: []*string{
	// 								to.Ptr("104.208.48.178")},
	// 							},
	// 							{
	// 								CustomBgpIPAddresses: []*string{
	// 									to.Ptr("169.254.21.10")},
	// 									DefaultBgpIPAddresses: []*string{
	// 										to.Ptr("10.30.0.5")},
	// 										IPConfigurationID: to.Ptr("Instance1"),
	// 										TunnelIPAddresses: []*string{
	// 											to.Ptr("104.208.48.179")},
	// 									}},
	// 									PeerWeight: to.Ptr[int32](0),
	// 								},
	// 								Connections: []*armnetwork.VPNConnection{
	// 									{
	// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/vpngw/vpnConnections/vpnConnection1"),
	// 										Name: to.Ptr("vpnConnection1"),
	// 										Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 										Properties: &armnetwork.VPNConnectionProperties{
	// 											ConnectionBandwidth: to.Ptr[int32](100),
	// 											ConnectionStatus: to.Ptr(armnetwork.VPNConnectionStatusConnected),
	// 											EgressBytesTransferred: to.Ptr[int64](0),
	// 											EnableBgp: to.Ptr(false),
	// 											IngressBytesTransferred: to.Ptr[int64](0),
	// 											IPSecPolicies: []*armnetwork.IPSecPolicy{
	// 											},
	// 											ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 											RemoteVPNSite: &armnetwork.SubResource{
	// 												ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"),
	// 											},
	// 											RoutingConfiguration: &armnetwork.RoutingConfiguration{
	// 												AssociatedRouteTable: &armnetwork.SubResource{
	// 													ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
	// 												},
	// 												PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
	// 													IDs: []*armnetwork.SubResource{
	// 														{
	// 															ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
	// 														},
	// 														{
	// 															ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable2"),
	// 														},
	// 														{
	// 															ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable3"),
	// 													}},
	// 													Labels: []*string{
	// 														to.Ptr("label1"),
	// 														to.Ptr("label2")},
	// 													},
	// 													VnetRoutes: &armnetwork.VnetRoute{
	// 														StaticRoutes: []*armnetwork.StaticRoute{
	// 														},
	// 													},
	// 												},
	// 												RoutingWeight: to.Ptr[int32](0),
	// 												UseLocalAzureIPAddress: to.Ptr(false),
	// 											},
	// 									}},
	// 									EnableBgpRouteTranslationForNat: to.Ptr(false),
	// 									IsRoutingPreferenceInternet: to.Ptr(false),
	// 									NatRules: []*armnetwork.VPNGatewayNatRule{
	// 									},
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 									VirtualHub: &armnetwork.SubResource{
	// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
	// 									},
	// 								},
	// 							}
}
