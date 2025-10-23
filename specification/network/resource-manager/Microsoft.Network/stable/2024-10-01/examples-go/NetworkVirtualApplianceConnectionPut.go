package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a9dbb28e788355a47dc5bad3ea5f8da212b4bf6/specification/network/resource-manager/Microsoft.Network/stable/2024-10-01/examples/NetworkVirtualApplianceConnectionPut.json
func ExampleVirtualApplianceConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualApplianceConnectionsClient().BeginCreateOrUpdate(ctx, "rg1", "nva1", "connection1", armnetwork.VirtualApplianceConnection{
		Properties: &armnetwork.VirtualApplianceConnectionProperties{
			Name: to.Ptr("connection1"),
			Asn:  to.Ptr[int64](64512),
			BgpPeerAddress: []*string{
				to.Ptr("169.254.16.13"),
				to.Ptr("169.254.16.14")},
			EnableInternetSecurity: to.Ptr(false),
			RoutingConfiguration: &armnetwork.RoutingConfiguration{
				AssociatedRouteTable: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
				},
				InboundRouteMap: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
				},
				OutboundRouteMap: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
				},
				PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
					IDs: []*armnetwork.SubResource{
						{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
						}},
					Labels: []*string{
						to.Ptr("label1")},
				},
			},
			TunnelIdentifier: to.Ptr[int64](0),
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
	// res.VirtualApplianceConnection = armnetwork.VirtualApplianceConnection{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva1/networkVirtualApplianceConnections/connection1"),
	// 	Name: to.Ptr("connection1"),
	// 	Properties: &armnetwork.VirtualApplianceConnectionProperties{
	// 		Name: to.Ptr("connection1"),
	// 		Asn: to.Ptr[int64](64512),
	// 		BgpPeerAddress: []*string{
	// 			to.Ptr("169.254.16.13"),
	// 			to.Ptr("169.254.16.14")},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			RoutingConfiguration: &armnetwork.RoutingConfiguration{
	// 				AssociatedRouteTable: &armnetwork.SubResource{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
	// 				},
	// 				InboundRouteMap: &armnetwork.SubResource{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
	// 				},
	// 				OutboundRouteMap: &armnetwork.SubResource{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
	// 				},
	// 				PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
	// 					IDs: []*armnetwork.SubResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
	// 					}},
	// 					Labels: []*string{
	// 						to.Ptr("label1")},
	// 					},
	// 				},
	// 				TunnelIdentifier: to.Ptr[int64](0),
	// 			},
	// 		}
}
