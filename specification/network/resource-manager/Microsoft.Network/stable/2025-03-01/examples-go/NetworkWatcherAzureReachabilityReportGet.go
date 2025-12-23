package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkWatcherAzureReachabilityReportGet.json
func ExampleWatchersClient_BeginGetAzureReachabilityReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginGetAzureReachabilityReport(ctx, "rg1", "nw1", armnetwork.AzureReachabilityReportParameters{
		AzureLocations: []*string{
			to.Ptr("West US")},
		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-10T00:00:00.000Z"); return t }()),
		ProviderLocation: &armnetwork.AzureReachabilityReportLocation{
			Country: to.Ptr("United States"),
			State:   to.Ptr("washington"),
		},
		Providers: []*string{
			to.Ptr("Frontier Communications of America, Inc. - ASN 5650")},
		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-07T00:00:00.000Z"); return t }()),
	}, nil)
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
	// res.AzureReachabilityReport = armnetwork.AzureReachabilityReport{
	// 	AggregationLevel: to.Ptr("State"),
	// 	ProviderLocation: &armnetwork.AzureReachabilityReportLocation{
	// 		Country: to.Ptr("United States"),
	// 		State: to.Ptr("washington"),
	// 	},
	// 	ReachabilityReport: []*armnetwork.AzureReachabilityReportItem{
	// 		{
	// 			AzureLocation: to.Ptr("West US"),
	// 			Latencies: []*armnetwork.AzureReachabilityReportLatencyInfo{
	// 				{
	// 					Score: to.Ptr[int32](94),
	// 					TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-07T00:00:00.000Z"); return t}()),
	// 				},
	// 				{
	// 					Score: to.Ptr[int32](94),
	// 					TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-08T00:00:00.000Z"); return t}()),
	// 				},
	// 				{
	// 					Score: to.Ptr[int32](94),
	// 					TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-09T00:00:00.000Z"); return t}()),
	// 			}},
	// 			Provider: to.Ptr("Frontier Communications of America, Inc. - ASN 5650"),
	// 	}},
	// }
}
