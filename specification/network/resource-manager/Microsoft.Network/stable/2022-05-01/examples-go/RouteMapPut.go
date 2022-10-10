package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/RouteMapPut.json
func ExampleRouteMapsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewRouteMapsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "virtualHub1", "routeMap1", armnetwork.RouteMap{
		Properties: &armnetwork.RouteMapProperties{
			AssociatedInboundConnections: []*string{
				to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGateway1/expressRouteConnections/exrConn1")},
			AssociatedOutboundConnections: []*string{},
			Rules: []*armnetwork.RouteMapRule{
				{
					Name: to.Ptr("rule1"),
					Actions: []*armnetwork.Action{
						{
							Type: to.Ptr(armnetwork.RouteMapActionTypeAdd),
							Parameters: []*armnetwork.Parameter{
								{
									AsPath: []*string{
										to.Ptr("22334")},
									Community:   []*string{},
									RoutePrefix: []*string{},
								}},
						}},
					MatchCriteria: []*armnetwork.Criterion{
						{
							AsPath:         []*string{},
							Community:      []*string{},
							MatchCondition: to.Ptr(armnetwork.RouteMapMatchConditionContains),
							RoutePrefix: []*string{
								to.Ptr("10.0.0.0/8")},
						}},
					NextStepIfMatched: to.Ptr(armnetwork.NextStepContinue),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
