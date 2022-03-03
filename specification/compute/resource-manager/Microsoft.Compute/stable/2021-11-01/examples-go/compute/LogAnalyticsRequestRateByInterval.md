Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.4.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/LogAnalyticsRequestRateByInterval.json
func ExampleLogAnalyticsClient_BeginExportRequestRateByInterval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewLogAnalyticsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginExportRequestRateByInterval(ctx,
		"<location>",
		armcompute.RequestRateByIntervalInput{
			BlobContainerSasURI: to.StringPtr("<blob-container-sas-uri>"),
			FromTime:            to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-21T01:54:06.862601Z"); return t }()),
			GroupByResourceName: to.BoolPtr(true),
			ToTime:              to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-23T01:54:06.862601Z"); return t }()),
			IntervalLength:      armcompute.IntervalInMinsFiveMins.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LogAnalyticsClientExportRequestRateByIntervalResult)
}
```
