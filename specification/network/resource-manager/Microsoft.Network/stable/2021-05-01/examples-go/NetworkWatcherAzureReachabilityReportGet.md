Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherAzureReachabilityReportGet.json
func ExampleWatchersClient_BeginGetAzureReachabilityReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginGetAzureReachabilityReport(ctx,
		"<resource-group-name>",
		"<network-watcher-name>",
		armnetwork.AzureReachabilityReportParameters{
			AzureLocations: []*string{
				to.Ptr("West US")},
			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-10T00:00:00Z"); return t }()),
			ProviderLocation: &armnetwork.AzureReachabilityReportLocation{
				Country: to.Ptr("<country>"),
				State:   to.Ptr("<state>"),
			},
			Providers: []*string{
				to.Ptr("Frontier Communications of America, Inc. - ASN 5650")},
			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-07T00:00:00Z"); return t }()),
		},
		&armnetwork.WatchersClientBeginGetAzureReachabilityReportOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
