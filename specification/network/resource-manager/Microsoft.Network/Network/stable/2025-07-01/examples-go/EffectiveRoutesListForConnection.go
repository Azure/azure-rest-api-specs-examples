package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/EffectiveRoutesListForConnection.json
func ExampleVirtualHubsClient_BeginGetEffectiveVirtualHubRoutes_effectiveRoutesForAConnectionResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubsClient().BeginGetEffectiveVirtualHubRoutes(ctx, "rg1", "virtualHub1", &armnetwork.VirtualHubsClientBeginGetEffectiveVirtualHubRoutesOptions{
		EffectiveRoutesParameters: &armnetwork.EffectiveRoutesParameters{
			ResourceID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName/expressRouteConnections/connectionName"),
			VirtualWanResourceType: to.Ptr("ExpressRouteConnection"),
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualHubsClientGetEffectiveVirtualHubRoutesResponse{
	// 	VirtualHubEffectiveRouteList: armnetwork.VirtualHubEffectiveRouteList{
	// 		Value: []*armnetwork.VirtualHubEffectiveRoute{
	// 			{
	// 				AddressPrefixes: []*string{
	// 					to.Ptr("10.147.128.0/17"),
	// 				},
	// 				AsPath: to.Ptr("65520-65520"),
	// 				NextHopType: to.Ptr("Remote Hub"),
	// 				NextHops: []*string{
	// 					to.Ptr("/subscriptions/testSub/resourceGroups/testRg/providers/Microsoft.Network/virtualHubs/hub0"),
	// 				},
	// 				RouteOrigin: to.Ptr("/subscriptions/testSub/resourceGroups/testRg/providers/Microsoft.Network/virtualHubs/hub0"),
	// 			},
	// 			{
	// 				AddressPrefixes: []*string{
	// 					to.Ptr("10.0.0.0/16"),
	// 				},
	// 				AsPath: to.Ptr("12076-12076"),
	// 				NextHopType: to.Ptr("ExpressRouteGateway"),
	// 				NextHops: []*string{
	// 					to.Ptr("/subscriptions/testSub/resourceGroups/testRg/providers/Microsoft.Network/expressRouteGateways/ErGw1"),
	// 				},
	// 				RouteOrigin: to.Ptr("/subscriptions/testSub/resourceGroups/testRg/providers/Microsoft.Network/expressRouteGateways/ErGw1"),
	// 			},
	// 		},
	// 	},
	// }
}
