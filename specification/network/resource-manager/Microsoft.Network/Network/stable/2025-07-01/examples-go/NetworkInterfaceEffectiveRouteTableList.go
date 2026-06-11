package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkInterfaceEffectiveRouteTableList.json
func ExampleInterfacesClient_BeginGetEffectiveRouteTable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInterfacesClient().BeginGetEffectiveRouteTable(ctx, "rg1", "nic1", nil)
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
	// res = armnetwork.InterfacesClientGetEffectiveRouteTableResponse{
	// 	EffectiveRouteListResult: armnetwork.EffectiveRouteListResult{
	// 		Value: []*armnetwork.EffectiveRoute{
	// 			{
	// 				AddressPrefix: []*string{
	// 					to.Ptr("172.20.2.0/24"),
	// 				},
	// 				NextHopIPAddress: []*string{
	// 				},
	// 				NextHopType: to.Ptr(armnetwork.RouteNextHopTypeVnetLocal),
	// 				Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 				State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 			},
	// 			{
	// 				AddressPrefix: []*string{
	// 					to.Ptr("0.0.0.0/0"),
	// 				},
	// 				NextHopIPAddress: []*string{
	// 				},
	// 				NextHopType: to.Ptr(armnetwork.RouteNextHopTypeInternet),
	// 				Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 				State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 			},
	// 			{
	// 				AddressPrefix: []*string{
	// 					to.Ptr("10.0.0.0/8"),
	// 				},
	// 				NextHopIPAddress: []*string{
	// 				},
	// 				NextHopType: to.Ptr(armnetwork.RouteNextHopTypeNone),
	// 				Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 				State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 			},
	// 			{
	// 				AddressPrefix: []*string{
	// 					to.Ptr("100.64.0.0/10"),
	// 				},
	// 				NextHopIPAddress: []*string{
	// 				},
	// 				NextHopType: to.Ptr(armnetwork.RouteNextHopTypeNone),
	// 				Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 				State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 			},
	// 			{
	// 				AddressPrefix: []*string{
	// 					to.Ptr("172.16.0.0/12"),
	// 				},
	// 				NextHopIPAddress: []*string{
	// 				},
	// 				NextHopType: to.Ptr(armnetwork.RouteNextHopTypeNone),
	// 				Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 				State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 			},
	// 			{
	// 				AddressPrefix: []*string{
	// 					to.Ptr("192.168.0.0/16"),
	// 				},
	// 				NextHopIPAddress: []*string{
	// 				},
	// 				NextHopType: to.Ptr(armnetwork.RouteNextHopTypeNone),
	// 				Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 				State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 			},
	// 		},
	// 	},
	// }
}
