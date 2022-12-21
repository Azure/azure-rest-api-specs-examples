package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/HubVirtualNetworkConnectionPut.json
func ExampleHubVirtualNetworkConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewHubVirtualNetworkConnectionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "virtualHub1", "connection1", armnetwork.HubVirtualNetworkConnection{
		Properties: &armnetwork.HubVirtualNetworkConnectionProperties{
			EnableInternetSecurity: to.Ptr(false),
			RemoteVirtualNetwork: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/SpokeVnet1"),
			},
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
						to.Ptr("label1"),
						to.Ptr("label2")},
				},
				VnetRoutes: &armnetwork.VnetRoute{
					StaticRoutes: []*armnetwork.StaticRoute{
						{
							Name: to.Ptr("route1"),
							AddressPrefixes: []*string{
								to.Ptr("10.1.0.0/16"),
								to.Ptr("10.2.0.0/16")},
							NextHopIPAddress: to.Ptr("10.0.0.68"),
						},
						{
							Name: to.Ptr("route2"),
							AddressPrefixes: []*string{
								to.Ptr("10.3.0.0/16"),
								to.Ptr("10.4.0.0/16")},
							NextHopIPAddress: to.Ptr("10.0.0.65"),
						}},
					StaticRoutesConfig: &armnetwork.StaticRoutesConfig{
						VnetLocalRouteOverrideCriteria: to.Ptr(armnetwork.VnetLocalRouteOverrideCriteriaEqual),
					},
				},
			},
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
