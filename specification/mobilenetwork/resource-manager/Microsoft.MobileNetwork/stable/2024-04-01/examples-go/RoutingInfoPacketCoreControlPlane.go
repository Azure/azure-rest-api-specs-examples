package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/RoutingInfoPacketCoreControlPlane.json
func ExampleRoutingInfoClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoutingInfoClient().Get(ctx, "rg1", "TestPacketCoreCP", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoutingInfoModel = armmobilenetwork.RoutingInfoModel{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlanes/routingInfo"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/routingInfo/default"),
	// 	Properties: &armmobilenetwork.RoutingInfoPropertiesFormat{
	// 		ControlPlaneAccessRoutes: []*armmobilenetwork.IPv4Route{
	// 			{
	// 				Destination: to.Ptr("198.51.100.1/32"),
	// 				NextHops: []*armmobilenetwork.IPv4RouteNextHop{
	// 					{
	// 						Address: to.Ptr("198.0.2.1"),
	// 						Priority: to.Ptr[int32](100),
	// 					},
	// 					{
	// 						Address: to.Ptr("198.0.2.128"),
	// 						Priority: to.Ptr[int32](100),
	// 				}},
	// 		}},
	// 		UserPlaneAccessRoutes: []*armmobilenetwork.IPv4Route{
	// 			{
	// 				Destination: to.Ptr("198.51.100.2/32"),
	// 				NextHops: []*armmobilenetwork.IPv4RouteNextHop{
	// 					{
	// 						Address: to.Ptr("198.0.2.2"),
	// 						Priority: to.Ptr[int32](100),
	// 					},
	// 					{
	// 						Address: to.Ptr("198.0.2.129"),
	// 						Priority: to.Ptr[int32](200),
	// 				}},
	// 		}},
	// 		UserPlaneDataRoutes: []*armmobilenetwork.UserPlaneDataRoutesItem{
	// 			{
	// 				AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
	// 				},
	// 				Routes: []*armmobilenetwork.IPv4Route{
	// 					{
	// 						Destination: to.Ptr("2.2.0.0/16"),
	// 						NextHops: []*armmobilenetwork.IPv4RouteNextHop{
	// 							{
	// 								Address: to.Ptr("203.0.113.1"),
	// 								Priority: to.Ptr[int32](100),
	// 							},
	// 							{
	// 								Address: to.Ptr("203.0.113.128"),
	// 								Priority: to.Ptr[int32](200),
	// 						}},
	// 					},
	// 					{
	// 						Destination: to.Ptr("2.4.0.0/16"),
	// 						NextHops: []*armmobilenetwork.IPv4RouteNextHop{
	// 							{
	// 								Address: to.Ptr("203.0.113.1"),
	// 								Priority: to.Ptr[int32](100),
	// 							},
	// 							{
	// 								Address: to.Ptr("203.0.113.128"),
	// 								Priority: to.Ptr[int32](200),
	// 						}},
	// 				}},
	// 		}},
	// 	},
	// }
}
