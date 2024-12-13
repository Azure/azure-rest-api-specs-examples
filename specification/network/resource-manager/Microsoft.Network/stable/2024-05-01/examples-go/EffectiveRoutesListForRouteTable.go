package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/EffectiveRoutesListForRouteTable.json
func ExampleVirtualHubsClient_BeginGetEffectiveVirtualHubRoutes_effectiveRoutesForARouteTableResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubsClient().BeginGetEffectiveVirtualHubRoutes(ctx, "rg1", "virtualHub1", &armnetwork.VirtualHubsClientBeginGetEffectiveVirtualHubRoutesOptions{EffectiveRoutesParameters: &armnetwork.EffectiveRoutesParameters{
		ResourceID:             to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
		VirtualWanResourceType: to.Ptr("RouteTable"),
	},
	})
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
	// res.VirtualHubEffectiveRouteList = armnetwork.VirtualHubEffectiveRouteList{
	// 	Value: []*armnetwork.VirtualHubEffectiveRoute{
	// 		{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("10.147.128.0/17")},
	// 				AsPath: to.Ptr("65520-65520"),
	// 				NextHopType: to.Ptr("Remote Hub"),
	// 				NextHops: []*string{
	// 					to.Ptr("/subscriptions/testSub/resourceGroups/testRg/providers/Microsoft.Network/virtualHubs/hub0")},
	// 					RouteOrigin: to.Ptr("/subscriptions/testSub/resourceGroups/testRg/providers/Microsoft.Network/virtualHubs/hub0"),
	// 				},
	// 				{
	// 					AddressPrefixes: []*string{
	// 						to.Ptr("10.0.0.0/16")},
	// 						AsPath: to.Ptr("12076-12076"),
	// 						NextHopType: to.Ptr("ExpressRouteGateway"),
	// 						NextHops: []*string{
	// 							to.Ptr("/subscriptions/testSub/resourceGroups/testRg/providers/Microsoft.Network/expressRouteGateways/ErGw1")},
	// 							RouteOrigin: to.Ptr("/subscriptions/testSub/resourceGroups/testRg/providers/Microsoft.Network/expressRouteGateways/ErGw1"),
	// 					}},
	// 				}
}
