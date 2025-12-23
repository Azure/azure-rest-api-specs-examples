package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkInterfaceEffectiveRouteTableList.json
func ExampleInterfacesClient_BeginGetEffectiveRouteTable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInterfacesClient().BeginGetEffectiveRouteTable(ctx, "rg1", "nic1", nil)
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
	// res.EffectiveRouteListResult = armnetwork.EffectiveRouteListResult{
	// 	Value: []*armnetwork.EffectiveRoute{
	// 		{
	// 			AddressPrefix: []*string{
	// 				to.Ptr("172.20.2.0/24")},
	// 				NextHopIPAddress: []*string{
	// 				},
	// 				NextHopType: to.Ptr(armnetwork.RouteNextHopTypeVnetLocal),
	// 				Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 				State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 			},
	// 			{
	// 				AddressPrefix: []*string{
	// 					to.Ptr("0.0.0.0/0")},
	// 					NextHopIPAddress: []*string{
	// 					},
	// 					NextHopType: to.Ptr(armnetwork.RouteNextHopTypeInternet),
	// 					Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 					State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 				},
	// 				{
	// 					AddressPrefix: []*string{
	// 						to.Ptr("10.0.0.0/8")},
	// 						NextHopIPAddress: []*string{
	// 						},
	// 						NextHopType: to.Ptr(armnetwork.RouteNextHopTypeNone),
	// 						Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 						State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 					},
	// 					{
	// 						AddressPrefix: []*string{
	// 							to.Ptr("100.64.0.0/10")},
	// 							NextHopIPAddress: []*string{
	// 							},
	// 							NextHopType: to.Ptr(armnetwork.RouteNextHopTypeNone),
	// 							Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 							State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 						},
	// 						{
	// 							AddressPrefix: []*string{
	// 								to.Ptr("172.16.0.0/12")},
	// 								NextHopIPAddress: []*string{
	// 								},
	// 								NextHopType: to.Ptr(armnetwork.RouteNextHopTypeNone),
	// 								Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 								State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 							},
	// 							{
	// 								AddressPrefix: []*string{
	// 									to.Ptr("192.168.0.0/16")},
	// 									NextHopIPAddress: []*string{
	// 									},
	// 									NextHopType: to.Ptr(armnetwork.RouteNextHopTypeNone),
	// 									Source: to.Ptr(armnetwork.EffectiveRouteSourceDefault),
	// 									State: to.Ptr(armnetwork.EffectiveRouteStateActive),
	// 							}},
	// 						}
}
