package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/HubVirtualNetworkConnectionList.json
func ExampleHubVirtualNetworkConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHubVirtualNetworkConnectionsClient().NewListPager("rg1", "virtualHub1", nil)
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
		// page = armnetwork.HubVirtualNetworkConnectionsClientListResponse{
		// 	ListHubVirtualNetworkConnectionsResult: armnetwork.ListHubVirtualNetworkConnectionsResult{
		// 		Value: []*armnetwork.HubVirtualNetworkConnection{
		// 			{
		// 				Name: to.Ptr("connection1"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/virtualHubVnetConnections/connection1"),
		// 				Properties: &armnetwork.HubVirtualNetworkConnectionProperties{
		// 					EnableInternetSecurity: to.Ptr(false),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RemoteVirtualNetwork: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1"),
		// 					},
		// 					RoutingConfiguration: &armnetwork.RoutingConfiguration{
		// 						AssociatedRouteTable: &armnetwork.SubResource{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
		// 						},
		// 						PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
		// 							IDs: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
		// 								},
		// 								{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable2"),
		// 								},
		// 								{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable3"),
		// 								},
		// 							},
		// 							Labels: []*string{
		// 								to.Ptr("label1"),
		// 								to.Ptr("label2"),
		// 							},
		// 						},
		// 						VnetRoutes: &armnetwork.VnetRoute{
		// 							BgpConnections: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/bgpConnections/bgpConn1"),
		// 								},
		// 							},
		// 							StaticRoutes: []*armnetwork.StaticRoute{
		// 								{
		// 									Name: to.Ptr("route1"),
		// 									AddressPrefixes: []*string{
		// 										to.Ptr("10.1.0.0/16"),
		// 										to.Ptr("10.2.0.0/16"),
		// 									},
		// 									NextHopIPAddress: to.Ptr("10.0.0.68"),
		// 								},
		// 								{
		// 									Name: to.Ptr("route2"),
		// 									AddressPrefixes: []*string{
		// 										to.Ptr("10.3.0.0/16"),
		// 										to.Ptr("10.4.0.0/16"),
		// 									},
		// 									NextHopIPAddress: to.Ptr("10.0.0.65"),
		// 								},
		// 							},
		// 							StaticRoutesConfig: &armnetwork.StaticRoutesConfig{
		// 								PropagateStaticRoutes: to.Ptr(true),
		// 								VnetLocalRouteOverrideCriteria: to.Ptr(armnetwork.VnetLocalRouteOverrideCriteriaEqual),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("connection2"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/virtualHubVnetConnections/connection2"),
		// 				Properties: &armnetwork.HubVirtualNetworkConnectionProperties{
		// 					EnableInternetSecurity: to.Ptr(false),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RemoteVirtualNetwork: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2"),
		// 					},
		// 					RoutingConfiguration: &armnetwork.RoutingConfiguration{
		// 						AssociatedRouteTable: &armnetwork.SubResource{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
		// 						},
		// 						PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
		// 							IDs: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
		// 								},
		// 								{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable2"),
		// 								},
		// 								{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable3"),
		// 								},
		// 							},
		// 							Labels: []*string{
		// 								to.Ptr("label1"),
		// 								to.Ptr("label2"),
		// 							},
		// 						},
		// 						VnetRoutes: &armnetwork.VnetRoute{
		// 							BgpConnections: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/bgpConnections/bgpConn2"),
		// 								},
		// 							},
		// 							StaticRoutes: []*armnetwork.StaticRoute{
		// 								{
		// 									Name: to.Ptr("route1"),
		// 									AddressPrefixes: []*string{
		// 										to.Ptr("10.1.0.0/16"),
		// 										to.Ptr("10.2.0.0/16"),
		// 									},
		// 									NextHopIPAddress: to.Ptr("10.0.0.68"),
		// 								},
		// 								{
		// 									Name: to.Ptr("route2"),
		// 									AddressPrefixes: []*string{
		// 										to.Ptr("10.3.0.0/16"),
		// 										to.Ptr("10.4.0.0/16"),
		// 									},
		// 									NextHopIPAddress: to.Ptr("10.0.0.65"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
