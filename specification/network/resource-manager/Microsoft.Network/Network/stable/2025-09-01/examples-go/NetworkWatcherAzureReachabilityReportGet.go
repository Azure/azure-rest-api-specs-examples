package armnetwork_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/NetworkWatcherAzureReachabilityReportGet.json
func ExampleWatchersClient_BeginGetAzureReachabilityReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginGetAzureReachabilityReport(ctx, "rg1", "nw1", armnetwork.AzureReachabilityReportParameters{
		AzureLocations: []*string{
			to.Ptr("West US"),
		},
		EndTime: to.Ptr(time.Date(2017, time.September, 10, 0, 0, 0, 0, time.UTC)),
		ProviderLocation: &armnetwork.AzureReachabilityReportLocation{
			Country: to.Ptr("United States"),
			State:   to.Ptr("washington"),
		},
		Providers: []*string{
			to.Ptr("Frontier Communications of America, Inc. - ASN 5650"),
		},
		StartTime: to.Ptr(time.Date(2017, time.September, 7, 0, 0, 0, 0, time.UTC)),
	}, nil)
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
	// res = armnetwork.WatchersClientGetAzureReachabilityReportResponse{
	// 	AzureReachabilityReport: armnetwork.AzureReachabilityReport{
	// 		AggregationLevel: to.Ptr("State"),
	// 		ProviderLocation: &armnetwork.AzureReachabilityReportLocation{
	// 			Country: to.Ptr("United States"),
	// 			State: to.Ptr("washington"),
	// 		},
	// 		ReachabilityReport: []*armnetwork.AzureReachabilityReportItem{
	// 			{
	// 				AzureLocation: to.Ptr("West US"),
	// 				Latencies: []*armnetwork.AzureReachabilityReportLatencyInfo{
	// 					{
	// 						Score: to.Ptr[int32](94),
	// 						TimeStamp: to.Ptr(time.Date(2017, time.September, 7, 0, 0, 0, 0, time.UTC)),
	// 					},
	// 					{
	// 						Score: to.Ptr[int32](94),
	// 						TimeStamp: to.Ptr(time.Date(2017, time.September, 8, 0, 0, 0, 0, time.UTC)),
	// 					},
	// 					{
	// 						Score: to.Ptr[int32](94),
	// 						TimeStamp: to.Ptr(time.Date(2017, time.September, 9, 0, 0, 0, 0, time.UTC)),
	// 					},
	// 				},
	// 				Provider: to.Ptr("Frontier Communications of America, Inc. - ASN 5650"),
	// 			},
	// 		},
	// 	},
	// }
}
