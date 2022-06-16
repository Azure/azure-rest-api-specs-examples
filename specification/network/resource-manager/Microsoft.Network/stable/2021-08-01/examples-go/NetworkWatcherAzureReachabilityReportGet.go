package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkWatcherAzureReachabilityReportGet.json
func ExampleWatchersClient_BeginGetAzureReachabilityReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginGetAzureReachabilityReport(ctx,
		"rg1",
		"nw1",
		armnetwork.AzureReachabilityReportParameters{
			AzureLocations: []*string{
				to.Ptr("West US")},
			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-10T00:00:00Z"); return t }()),
			ProviderLocation: &armnetwork.AzureReachabilityReportLocation{
				Country: to.Ptr("United States"),
				State:   to.Ptr("washington"),
			},
			Providers: []*string{
				to.Ptr("Frontier Communications of America, Inc. - ASN 5650")},
			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-07T00:00:00Z"); return t }()),
		},
		nil)
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
