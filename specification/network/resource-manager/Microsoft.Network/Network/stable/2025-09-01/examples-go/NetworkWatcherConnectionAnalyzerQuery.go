package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/NetworkWatcherConnectionAnalyzerQuery.json
func ExampleWatchersClient_BeginConnectionAnalyzersQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("7f4a1d92-3b6e-4c8f-9a25-e1b8c3d7f024", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginConnectionAnalyzersQuery(ctx, "connectionAnalyzerRG", "nw1", "ca1", nil)
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
	// res = armnetwork.WatchersClientConnectionAnalyzersQueryResponse{
	// 	ConnectionAnalyzerQueryStatusResult: armnetwork.ConnectionAnalyzerQueryStatusResult{
	// 		ID: to.Ptr("/subscriptions/7f4a1d92-3b6e-4c8f-9a25-e1b8c3d7f024/resourceGroups/connectionAnalyzerRG/providers/Microsoft.Network/networkWatchers/nw1/connectionAnalyzers/ca1"),
	// 		ConnectionAnalyzerStatus: to.Ptr(armnetwork.ConnectionAnalyzerStatusSucceeded),
	// 		OutputStoragePath: to.Ptr("https://sa1.blob.core.windows.net/network-watcher-logs/connectionanalyzer/ca1_2025_09_01_10_00_00.json"),
	// 		ExpiryInUTC: to.Ptr(time.Date(2025, time.October, 1, 10, 0, 0, 0, time.UTC)),
	// 		DiagnosticOperationResults: []*armnetwork.DiagnosticOperationResult{
	// 			{
	// 				DiagnosticOperation: to.Ptr(armnetwork.DiagnosticOperationConnectivityCheck),
	// 				Result: to.Ptr("{\"ConnectivityStatus\":\"Reachable\"}"),
	// 			},
	// 		},
	// 	},
	// }
}
