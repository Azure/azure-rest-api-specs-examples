package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/RouteMapGet.json
func ExampleRouteMapsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRouteMapsClient().Get(ctx, "rg1", "virtualHub1", "routeMap1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.RouteMapsClientGetResponse{
	// 	RouteMap: armnetwork.RouteMap{
	// 		Name: to.Ptr("routeMap1"),
	// 		Type: to.Ptr("Microsoft.Network/virtualHubs/routeMaps"),
	// 		Etag: to.Ptr("W/\"e203e953-7ba7-4302-a246-aa2ec03f6edf\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
	// 		Properties: &armnetwork.RouteMapProperties{
	// 			AssociatedInboundConnections: []*string{
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGateway1/expressRouteConnections/exrConn1"),
	// 			},
	// 			AssociatedOutboundConnections: []*string{
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Rules: []*armnetwork.RouteMapRule{
	// 				{
	// 					Name: to.Ptr("rule1"),
	// 					Actions: []*armnetwork.Action{
	// 						{
	// 							Type: to.Ptr(armnetwork.RouteMapActionTypeAdd),
	// 							Parameters: []*armnetwork.Parameter{
	// 								{
	// 									AsPath: []*string{
	// 										to.Ptr("22334"),
	// 									},
	// 									Community: []*string{
	// 									},
	// 									RoutePrefix: []*string{
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 					MatchCriteria: []*armnetwork.Criterion{
	// 						{
	// 							AsPath: []*string{
	// 							},
	// 							Community: []*string{
	// 							},
	// 							MatchCondition: to.Ptr(armnetwork.RouteMapMatchConditionContains),
	// 							RoutePrefix: []*string{
	// 								to.Ptr("10.0.0.0/8"),
	// 							},
	// 						},
	// 					},
	// 					NextStepIfMatched: to.Ptr(armnetwork.NextStepContinue),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
