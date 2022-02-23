Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherAzureReachabilityReportGet.json
func ExampleWatchersClient_BeginGetAzureReachabilityReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewWatchersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginGetAzureReachabilityReport(ctx,
		"<resource-group-name>",
		"<network-watcher-name>",
		armnetwork.AzureReachabilityReportParameters{
			AzureLocations: []*string{
				to.StringPtr("West US")},
			EndTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-10T00:00:00Z"); return t }()),
			ProviderLocation: &armnetwork.AzureReachabilityReportLocation{
				Country: to.StringPtr("<country>"),
				State:   to.StringPtr("<state>"),
			},
			Providers: []*string{
				to.StringPtr("Frontier Communications of America, Inc. - ASN 5650")},
			StartTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-07T00:00:00Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WatchersClientGetAzureReachabilityReportResult)
}
```
