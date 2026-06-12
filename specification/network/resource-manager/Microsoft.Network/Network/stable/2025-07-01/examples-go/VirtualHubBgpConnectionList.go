package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualHubBgpConnectionList.json
func ExampleVirtualHubBgpConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualHubBgpConnectionsClient().NewListPager("rg1", "hub1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armnetwork.VirtualHubBgpConnectionsClientListResponse{
		// 	ListVirtualHubBgpConnectionResults: armnetwork.ListVirtualHubBgpConnectionResults{
		// 		Value: []*armnetwork.BgpConnection{
		// 			{
		// 				Name: to.Ptr("conn1"),
		// 				Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/bgpConnections/conn1"),
		// 				Properties: &armnetwork.BgpConnectionProperties{
		// 					ConnectionState: to.Ptr(armnetwork.HubBgpConnectionStatusConnected),
		// 					HubVirtualNetworkConnection: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubVirtualNetworkConnections/hubVnetConn1"),
		// 					},
		// 					PeerAsn: to.Ptr[int64](20000),
		// 					PeerIP: to.Ptr("192.168.1.5"),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RoutingConfiguration: &armnetwork.RoutingConfiguration{
		// 						InboundRouteMap: &armnetwork.SubResource{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/routeMaps/routeMap1"),
		// 						},
		// 						OutboundRouteMap: &armnetwork.SubResource{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/routeMaps/routeMap2"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
