package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteCircuitGetCircuitLinkFailoverSingleTestDetails.json
func ExampleExpressRouteCircuitsClient_BeginGetCircuitLinkFailoverSingleTestDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCircuitsClient().BeginGetCircuitLinkFailoverSingleTestDetails(ctx, "rg1", "circuit1", "Primary", "BgpDisconnect", "00000000-0000-0000-0000-000000000001", nil)
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
	// res = armnetwork.ExpressRouteCircuitsClientGetCircuitLinkFailoverSingleTestDetailsResponse{
	// 	ExpressRouteLinkFailoverSingleTestDetailsArray: []*armnetwork.ExpressRouteLinkFailoverSingleTestDetails{
	// 		{
	// 			StartTimeUTC: to.Ptr("2025-01-01T00:00:00Z"),
	// 			EndTimeUTC: to.Ptr("2025-01-01T01:00:00Z"),
	// 			Status: to.Ptr(armnetwork.FailoverTestStatusCompleted),
	// 			WasSimulationSuccessful: to.Ptr(true),
	// 			LinkType: to.Ptr(armnetwork.ExpressRouteFailoverLinkTypePrimary),
	// 			CircuitTestCategory: to.Ptr(armnetwork.MaintenanceTestCategoryBgpDisconnect),
	// 			IsSimulationVerified: to.Ptr(true),
	// 			RedundantRoutes: &armnetwork.ExpressRouteLinkFailoverRouteList{
	// 				BeforeSimulation: []*armnetwork.ExpressRouteLinkFailoverRoute{
	// 					{
	// 						Route: to.Ptr("10.0.0.0/24"),
	// 						NextHop: to.Ptr("10.0.0.1"),
	// 						PrimaryASPath: to.Ptr("65000 65001"),
	// 						SecondaryASPath: to.Ptr("65000 65002"),
	// 					},
	// 				},
	// 				DuringSimulation: []*armnetwork.ExpressRouteLinkFailoverRoute{
	// 					{
	// 						Route: to.Ptr("10.0.0.0/24"),
	// 						NextHop: to.Ptr("10.0.0.2"),
	// 						PrimaryASPath: to.Ptr("65000 65001"),
	// 						SecondaryASPath: to.Ptr("65000 65002"),
	// 					},
	// 				},
	// 			},
	// 			NonRedundantRoutes: &armnetwork.ExpressRouteLinkFailoverRouteList{
	// 				BeforeSimulation: []*armnetwork.ExpressRouteLinkFailoverRoute{
	// 				},
	// 				DuringSimulation: []*armnetwork.ExpressRouteLinkFailoverRoute{
	// 				},
	// 			},
	// 			BgpStatus: []*armnetwork.ExpressRouteLinkFailoverTestBgpStatus{
	// 				{
	// 					Type: to.Ptr(armnetwork.ExpressRouteFailoverBgpStatusAddressFamilyIPv4),
	// 					Link: to.Ptr(armnetwork.ExpressRouteFailoverLinkTypePrimary),
	// 					Status: to.Ptr(armnetwork.ExpressRouteLinkFailoverBgpStatusConnected),
	// 					CheckTimeUTC: to.Ptr("2025-01-01T00:30:00Z"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
