package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/NetworkWatcherConnectionMonitorQuery.json
func ExampleConnectionMonitorsClient_BeginQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionMonitorsClient().BeginQuery(ctx, "rg1", "nw1", "cm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectionMonitorQueryResult = armnetwork.ConnectionMonitorQueryResult{
	// 	SourceStatus: to.Ptr(armnetwork.ConnectionMonitorSourceStatusActive),
	// 	States: []*armnetwork.ConnectionStateSnapshot{
	// 		{
	// 			ConnectionState: to.Ptr(armnetwork.ConnectionStateReachable),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-08T05:12:41.5265438Z"); return t}()),
	// 			EvaluationState: to.Ptr(armnetwork.EvaluationStateCompleted),
	// 			Hops: []*armnetwork.ConnectivityHop{
	// 				{
	// 					Type: to.Ptr("Source"),
	// 					Address: to.Ptr("10.1.1.4"),
	// 					ID: to.Ptr("7dbbe7aa-60ba-4650-831e-63d775d38e9e"),
	// 					Issues: []*armnetwork.ConnectivityIssue{
	// 					},
	// 					NextHopIDs: []*string{
	// 						to.Ptr("75c8d819-b208-4584-a311-1aa45ce753f9")},
	// 						ResourceID: to.Ptr("subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic0/ipConfigurations/ipconfig1"),
	// 					},
	// 					{
	// 						Type: to.Ptr("VirtualNetwork"),
	// 						Address: to.Ptr("192.168.100.4"),
	// 						ID: to.Ptr("75c8d819-b208-4584-a311-1aa45ce753f9"),
	// 						Issues: []*armnetwork.ConnectivityIssue{
	// 						},
	// 						NextHopIDs: []*string{
	// 						},
	// 						ResourceID: to.Ptr("subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic1/ipConfigurations/ipconfig1"),
	// 				}},
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-08T03:42:33.3387305Z"); return t}()),
	// 		}},
	// 	}
}
