package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/HubRouteTableList.json
func ExampleHubRouteTablesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHubRouteTablesClient().NewListPager("rg1", "virtualHub1", nil)
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
		// page = armnetwork.HubRouteTablesClientListResponse{
		// 	ListHubRouteTablesResult: armnetwork.ListHubRouteTablesResult{
		// 		Value: []*armnetwork.HubRouteTable{
		// 			{
		// 				Name: to.Ptr("hubRouteTable1"),
		// 				Type: to.Ptr("Microsoft.Network/virtualHubs/hubRouteTables"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
		// 				Properties: &armnetwork.HubRouteTableProperties{
		// 					AssociatedConnections: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubVirtualNetworkConnections/vnetConnn1"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubVirtualNetworkConnections/vnetConnn2"),
		// 					},
		// 					Labels: []*string{
		// 						to.Ptr("label1"),
		// 						to.Ptr("label2"),
		// 					},
		// 					PropagatingConnections: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/erg1/expressRouteConnections/erConn1"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/vpngw1/vpnConnections/vpnConn2"),
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					Routes: []*armnetwork.HubRoute{
		// 						{
		// 							Name: to.Ptr("route1a"),
		// 							DestinationType: to.Ptr("CIDR"),
		// 							Destinations: []*string{
		// 								to.Ptr("10.0.0.0/8"),
		// 								to.Ptr("20.0.0.0/8"),
		// 								to.Ptr("30.0.0.0/8"),
		// 							},
		// 							NextHop: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azureFirewall1"),
		// 							NextHopType: to.Ptr("ResourceId"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("hubRouteTable2"),
		// 				Type: to.Ptr("Microsoft.Network/virtualHubs/hubRouteTables"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable2"),
		// 				Properties: &armnetwork.HubRouteTableProperties{
		// 					AssociatedConnections: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubVirtualNetworkConnections/vnetConnn3"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubVirtualNetworkConnections/vnetConnn4"),
		// 					},
		// 					Labels: []*string{
		// 						to.Ptr("label3"),
		// 					},
		// 					PropagatingConnections: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/erg1/expressRouteConnections/erConn2"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/vpngw1/vpnConnections/vpnConn1"),
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					Routes: []*armnetwork.HubRoute{
		// 						{
		// 							Name: to.Ptr("route2a"),
		// 							DestinationType: to.Ptr("CIDR"),
		// 							Destinations: []*string{
		// 								to.Ptr("40.0.0.0/8"),
		// 								to.Ptr("50.0.0.0/8"),
		// 								to.Ptr("60.0.0.0/8"),
		// 							},
		// 							NextHop: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubVirtualNetworkConnections/vn1"),
		// 							NextHopType: to.Ptr("ResourceId"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
