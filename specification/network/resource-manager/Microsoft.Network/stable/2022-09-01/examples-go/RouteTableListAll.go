package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/RouteTableListAll.json
func ExampleRouteTablesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRouteTablesClient().NewListAllPager(nil)
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
		// page.RouteTableListResult = armnetwork.RouteTableListResult{
		// 	Value: []*armnetwork.RouteTable{
		// 		{
		// 			Name: to.Ptr("testrt"),
		// 			Type: to.Ptr("Microsoft.Network/routeTables"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/routeTables/testrt"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.RouteTablePropertiesFormat{
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				Routes: []*armnetwork.Route{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/routeTables/testrt/routes/route1"),
		// 						Name: to.Ptr("route1"),
		// 						Properties: &armnetwork.RoutePropertiesFormat{
		// 							AddressPrefix: to.Ptr("10.0.3.0/24"),
		// 							NextHopType: to.Ptr(armnetwork.RouteNextHopTypeVirtualNetworkGateway),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						},
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testrt3"),
		// 			Type: to.Ptr("Microsoft.Network/routeTables"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/routeTables/testrt3"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.RouteTablePropertiesFormat{
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				Routes: []*armnetwork.Route{
		// 				},
		// 			},
		// 	}},
		// }
	}
}
