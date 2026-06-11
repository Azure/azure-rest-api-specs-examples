package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteGatewayGetFailoverSingleTestDetails.json
func ExampleExpressRouteGatewaysClient_BeginGetFailoverSingleTestDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteGatewaysClient().BeginGetFailoverSingleTestDetails(ctx, "rg1", "ergw1", "Vancouver", "00000000-0000-0000-0000-000000000001", nil)
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
	// res = armnetwork.ExpressRouteGatewaysClientGetFailoverSingleTestDetailsResponse{
	// 	ExpressRouteFailoverSingleTestDetailsArray: []*armnetwork.ExpressRouteFailoverSingleTestDetails{
	// 		{
	// 			PeeringLocation: to.Ptr("Hong Kong"),
	// 			StartTimeUTC: to.Ptr("2026-02-11T07:37:04Z"),
	// 			EndTimeUTC: to.Ptr("2026-02-11T08:41:50Z"),
	// 			Status: to.Ptr(armnetwork.FailoverTestStatusForSingleTestCompleted),
	// 			WasSimulationSuccessful: to.Ptr(true),
	// 			RedundantRoutes: []*armnetwork.ExpressRouteFailoverRedundantRoute{
	// 				{
	// 					PeeringLocations: []*string{
	// 						to.Ptr("Atlanta"),
	// 						to.Ptr("Silicon Valley"),
	// 					},
	// 					Routes: []*string{
	// 						to.Ptr("10.0.0.0/24"),
	// 					},
	// 				},
	// 			},
	// 			NonRedundantRoutes: []*string{
	// 				to.Ptr("10.10.0.0/16"),
	// 				to.Ptr("10.11.0.0/24"),
	// 			},
	// 			FailoverConnectionDetails: []*armnetwork.FailoverConnectionDetails{
	// 				{
	// 					FailoverConnectionName: to.Ptr("conn1"),
	// 					FailoverLocation: to.Ptr("Hong Kong2"),
	// 					IsVerified: to.Ptr(true),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
