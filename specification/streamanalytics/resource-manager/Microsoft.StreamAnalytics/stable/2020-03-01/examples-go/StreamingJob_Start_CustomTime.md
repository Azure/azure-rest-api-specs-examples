Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstreamanalytics%2Farmstreamanalytics%2Fv0.3.0/sdk/resourcemanager/streamanalytics/armstreamanalytics/README.md) on how to add the SDK to your project and authenticate.

```go
package armstreamanalytics_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Start_CustomTime.json
func ExampleStreamingJobsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstreamanalytics.NewStreamingJobsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStart(ctx,
		"<resource-group-name>",
		"<job-name>",
		&armstreamanalytics.StreamingJobsClientBeginStartOptions{StartJobParameters: &armstreamanalytics.StartStreamingJobParameters{
			OutputStartMode: armstreamanalytics.OutputStartMode("CustomTime").ToPtr(),
			OutputStartTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2012-12-12T12:12:12Z"); return t }()),
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
