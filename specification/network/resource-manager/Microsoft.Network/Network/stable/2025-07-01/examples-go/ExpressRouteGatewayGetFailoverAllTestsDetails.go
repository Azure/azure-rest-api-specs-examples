package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteGatewayGetFailoverAllTestsDetails.json
func ExampleExpressRouteGatewaysClient_BeginGetFailoverAllTestsDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteGatewaysClient().BeginGetFailoverAllTestsDetails(ctx, "rg1", "ergw1", &armnetwork.ExpressRouteGatewaysClientBeginGetFailoverAllTestsDetailsOptions{
		Type:        to.Ptr("SingleSiteFailover"),
		FetchLatest: to.Ptr(true)})
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
	// res = armnetwork.ExpressRouteGatewaysClientGetFailoverAllTestsDetailsResponse{
	// 	ExpressRouteFailoverTestDetailsArray: []*armnetwork.ExpressRouteFailoverTestDetails{
	// 		{
	// 			PeeringLocation: to.Ptr("Hong Kong"),
	// 			Circuits: []*armnetwork.ExpressRouteFailoverCircuitResourceDetails{
	// 				{
	// 					NrpResourceURI: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuit1"),
	// 					Name: to.Ptr("circuit1"),
	// 					ConnectionName: to.Ptr("conn1"),
	// 				},
	// 			},
	// 			Status: to.Ptr(armnetwork.FailoverTestStatusCompleted),
	// 			StartTime: to.Ptr("2026-02-11T07:37:04Z"),
	// 			EndTime: to.Ptr("2026-02-11T08:41:50Z"),
	// 			Connections: []*armnetwork.ExpressRouteFailoverConnectionResourceDetails{
	// 				{
	// 					NrpResourceURI: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/ergw1/expressRouteConnections/conn1"),
	// 					Name: to.Ptr("conn1"),
	// 					Status: to.Ptr(armnetwork.FailoverConnectionStatusConnected),
	// 					LastUpdatedTime: to.Ptr("2026-02-11T08:41:49Z"),
	// 				},
	// 			},
	// 			TestGUID: to.Ptr("16546d1a-0548-4d6d-84ca-7149f3af4403"),
	// 			TestType: to.Ptr(armnetwork.FailoverTestTypeSingleSiteFailover),
	// 			Issues: []*string{
	// 			},
	// 		},
	// 		{
	// 			PeeringLocation: to.Ptr("Hong Kong2"),
	// 			Circuits: []*armnetwork.ExpressRouteFailoverCircuitResourceDetails{
	// 				{
	// 					NrpResourceURI: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuit2"),
	// 					Name: to.Ptr("circuit2"),
	// 					ConnectionName: to.Ptr("conn2"),
	// 				},
	// 			},
	// 			Status: to.Ptr(armnetwork.FailoverTestStatusNotStarted),
	// 			StartTime: to.Ptr(""),
	// 			EndTime: to.Ptr(""),
	// 			Connections: []*armnetwork.ExpressRouteFailoverConnectionResourceDetails{
	// 				{
	// 					NrpResourceURI: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/ergw1/expressRouteConnections/conn2"),
	// 					Name: to.Ptr("conn2"),
	// 					Status: to.Ptr(armnetwork.FailoverConnectionStatusConnected),
	// 				},
	// 			},
	// 			TestGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TestType: to.Ptr(armnetwork.FailoverTestTypeSingleSiteFailover),
	// 			Issues: []*string{
	// 			},
	// 		},
	// 	},
	// }
}
