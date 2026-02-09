package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/RouteMapList.json
func ExampleRouteMapsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRouteMapsClient().NewListPager("rg1", "virtualHub1", nil)
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
		// page.ListRouteMapsResult = armnetwork.ListRouteMapsResult{
		// 	Value: []*armnetwork.RouteMap{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
		// 			Name: to.Ptr("routeMap1"),
		// 			Type: to.Ptr("Microsoft.Network/virtualHubs/routeMaps"),
		// 			Etag: to.Ptr("W/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.RouteMapProperties{
		// 				AssociatedInboundConnections: []*string{
		// 					to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGateway1/expressRouteConnections/exrConn1")},
		// 					AssociatedOutboundConnections: []*string{
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					Rules: []*armnetwork.RouteMapRule{
		// 						{
		// 							Name: to.Ptr("rule1"),
		// 							Actions: []*armnetwork.Action{
		// 								{
		// 									Type: to.Ptr(armnetwork.RouteMapActionTypeAdd),
		// 									Parameters: []*armnetwork.Parameter{
		// 										{
		// 											AsPath: []*string{
		// 												to.Ptr("22334")},
		// 												Community: []*string{
		// 												},
		// 												RoutePrefix: []*string{
		// 												},
		// 										}},
		// 								}},
		// 								MatchCriteria: []*armnetwork.Criterion{
		// 									{
		// 										AsPath: []*string{
		// 										},
		// 										Community: []*string{
		// 										},
		// 										MatchCondition: to.Ptr(armnetwork.RouteMapMatchConditionContains),
		// 										RoutePrefix: []*string{
		// 											to.Ptr("10.0.0.0/8")},
		// 									}},
		// 									NextStepIfMatched: to.Ptr(armnetwork.NextStepContinue),
		// 							}},
		// 						},
		// 				}},
		// 			}
	}
}
