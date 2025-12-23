package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRouteGatewayListBySubscription.json
func ExampleExpressRouteGatewaysClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteGatewaysClient().ListBySubscription(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRouteGatewayList = armnetwork.ExpressRouteGatewayList{
	// 	Value: []*armnetwork.ExpressRouteGateway{
	// 		{
	// 			Name: to.Ptr("expressRouteGatewayName"),
	// 			Type: to.Ptr("Microsoft.Network/expressRouteGateways"),
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName"),
	// 			Location: to.Ptr("westus"),
	// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 			Properties: &armnetwork.ExpressRouteGatewayProperties{
	// 				AllowNonVirtualWanTraffic: to.Ptr(false),
	// 				AutoScaleConfiguration: &armnetwork.ExpressRouteGatewayPropertiesAutoScaleConfiguration{
	// 					Bounds: &armnetwork.ExpressRouteGatewayPropertiesAutoScaleConfigurationBounds{
	// 						Min: to.Ptr[int32](2),
	// 					},
	// 				},
	// 				ExpressRouteConnections: []*armnetwork.ExpressRouteConnection{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName/expressRouteConnections/connectionName"),
	// 						Name: to.Ptr("connectionName"),
	// 						Properties: &armnetwork.ExpressRouteConnectionProperties{
	// 							AuthorizationKey: to.Ptr("f28e9c99-78d8-4248-a855-c54cf6beb99d"),
	// 							EnableInternetSecurity: to.Ptr(false),
	// 							ExpressRouteCircuitPeering: &armnetwork.ExpressRouteCircuitPeeringID{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering"),
	// 							},
	// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 							RoutingConfiguration: &armnetwork.RoutingConfiguration{
	// 								AssociatedRouteTable: &armnetwork.SubResource{
	// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/virtualHubName/hubRouteTables/hubRouteTable1"),
	// 								},
	// 								PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
	// 									IDs: []*armnetwork.SubResource{
	// 										{
	// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/virtualHubName/hubRouteTables/hubRouteTable1"),
	// 										},
	// 										{
	// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/virtualHubName/hubRouteTables/hubRouteTable2"),
	// 										},
	// 										{
	// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hvirtualHubNameub1/hubRouteTables/hubRouteTable3"),
	// 									}},
	// 									Labels: []*string{
	// 										to.Ptr("label1"),
	// 										to.Ptr("label2")},
	// 									},
	// 									VnetRoutes: &armnetwork.VnetRoute{
	// 										StaticRoutes: []*armnetwork.StaticRoute{
	// 										},
	// 									},
	// 								},
	// 								RoutingWeight: to.Ptr[int32](1),
	// 							},
	// 					}},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					VirtualHub: &armnetwork.VirtualHubID{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/virtualHubName"),
	// 					},
	// 				},
	// 		}},
	// 	}
}
