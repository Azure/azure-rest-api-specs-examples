package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/RouteTableGet.json
func ExampleRouteTablesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRouteTablesClient().Get(ctx, "rg1", "testrt", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.RouteTablesClientGetResponse{
	// 	RouteTable: armnetwork.RouteTable{
	// 		Name: to.Ptr("testrt"),
	// 		Type: to.Ptr("Microsoft.Network/routeTables"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/routeTables/testrt"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetwork.RouteTablePropertiesFormat{
	// 			DisableBgpRoutePropagation: to.Ptr(false),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Routes: []*armnetwork.Route{
	// 				{
	// 					Name: to.Ptr("route1"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/routeTables/testrt/routes/route1"),
	// 					Properties: &armnetwork.RoutePropertiesFormat{
	// 						AddressPrefix: to.Ptr("10.0.3.0/24"),
	// 						NextHopType: to.Ptr(armnetwork.RouteNextHopTypeVirtualNetworkGateway),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
