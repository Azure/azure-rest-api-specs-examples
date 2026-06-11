package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteCircuitGetCircuitLinkFailoverAllTestsDetails.json
func ExampleExpressRouteCircuitsClient_BeginGetCircuitLinkFailoverAllTestsDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCircuitsClient().BeginGetCircuitLinkFailoverAllTestsDetails(ctx, "rg1", "circuit1", &armnetwork.ExpressRouteCircuitsClientBeginGetCircuitLinkFailoverAllTestsDetailsOptions{
		FailoverTestType: to.Ptr("SingleSiteFailover"),
		FetchLatest:      to.Ptr(true)})
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
	// res = armnetwork.ExpressRouteCircuitsClientGetCircuitLinkFailoverAllTestsDetailsResponse{
	// 	ExpressRouteLinkFailoverAllTestsDetailsArray: []*armnetwork.ExpressRouteLinkFailoverAllTestsDetails{
	// 		{
	// 			Status: to.Ptr(armnetwork.FailoverTestStatusCompleted),
	// 			StartTime: to.Ptr("2025-01-01T00:00:00Z"),
	// 			EndTime: to.Ptr("2025-01-01T01:00:00Z"),
	// 			TestGUID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 			TestType: to.Ptr(armnetwork.FailoverTestTypeSingleSiteFailover),
	// 			Issues: []*string{
	// 			},
	// 			WasSimulationSuccessful: to.Ptr(true),
	// 			CircuitTestCategory: to.Ptr(armnetwork.MaintenanceTestCategoryBgpDisconnect),
	// 			LinkType: to.Ptr(armnetwork.ExpressRouteFailoverLinkTypePrimary),
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
