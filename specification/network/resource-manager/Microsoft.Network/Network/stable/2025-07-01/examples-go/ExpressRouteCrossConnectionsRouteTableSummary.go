package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteCrossConnectionsRouteTableSummary.json
func ExampleExpressRouteCrossConnectionsClient_BeginListRoutesTableSummary() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCrossConnectionsClient().BeginListRoutesTableSummary(ctx, "CrossConnection-SiliconValley", "<circuitServiceKey>", "AzurePrivatePeering", "primary", nil)
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
	// res = armnetwork.ExpressRouteCrossConnectionsClientListRoutesTableSummaryResponse{
	// 	ExpressRouteCrossConnectionsRoutesTableSummaryListResult: armnetwork.ExpressRouteCrossConnectionsRoutesTableSummaryListResult{
	// 		Value: []*armnetwork.ExpressRouteCrossConnectionRoutesTableSummary{
	// 			{
	// 				Asn: to.Ptr[int32](65514),
	// 				Neighbor: to.Ptr("10.6.1.112"),
	// 				StateOrPrefixesReceived: to.Ptr("Active"),
	// 				UpDown: to.Ptr("1d14h"),
	// 			},
	// 			{
	// 				Asn: to.Ptr[int32](65514),
	// 				Neighbor: to.Ptr("10.6.1.113"),
	// 				StateOrPrefixesReceived: to.Ptr("1"),
	// 				UpDown: to.Ptr("1d14h"),
	// 			},
	// 		},
	// 	},
	// }
}
