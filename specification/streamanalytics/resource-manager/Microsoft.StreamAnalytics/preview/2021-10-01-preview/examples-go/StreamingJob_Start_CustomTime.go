package armstreamanalytics_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/StreamingJob_Start_CustomTime.json
func ExampleStreamingJobsClient_BeginStart_startAStreamingJobWithCustomTimeOutputStartMode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStreamingJobsClient().BeginStart(ctx, "sjrg6936", "sj59", &armstreamanalytics.StreamingJobsClientBeginStartOptions{StartJobParameters: &armstreamanalytics.StartStreamingJobParameters{
		OutputStartMode: to.Ptr(armstreamanalytics.OutputStartModeCustomTime),
		OutputStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2012-12-12T12:12:12.000Z"); return t }()),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
