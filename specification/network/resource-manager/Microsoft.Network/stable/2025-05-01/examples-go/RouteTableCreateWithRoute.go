package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/RouteTableCreateWithRoute.json
func ExampleRouteTablesClient_BeginCreateOrUpdate_createRouteTableWithRoute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRouteTablesClient().BeginCreateOrUpdate(ctx, "rg1", "testrt", armnetwork.RouteTable{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.RouteTablePropertiesFormat{
			DisableBgpRoutePropagation: to.Ptr(true),
			Routes: []*armnetwork.Route{
				{
					Name: to.Ptr("route1"),
					Properties: &armnetwork.RoutePropertiesFormat{
						AddressPrefix: to.Ptr("10.0.3.0/24"),
						NextHopType:   to.Ptr(armnetwork.RouteNextHopTypeVirtualNetworkGateway),
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RouteTable = armnetwork.RouteTable{
	// 	Name: to.Ptr("testrt"),
	// 	Type: to.Ptr("Microsoft.Network/routeTables"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/routeTables/testrt"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.RouteTablePropertiesFormat{
	// 		DisableBgpRoutePropagation: to.Ptr(true),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		Routes: []*armnetwork.Route{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/routeTables/testrt/routes/route1"),
	// 				Name: to.Ptr("route1"),
	// 				Properties: &armnetwork.RoutePropertiesFormat{
	// 					AddressPrefix: to.Ptr("10.0.3.0/24"),
	// 					NextHopType: to.Ptr(armnetwork.RouteNextHopTypeVirtualNetworkGateway),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 	},
	// }
}
