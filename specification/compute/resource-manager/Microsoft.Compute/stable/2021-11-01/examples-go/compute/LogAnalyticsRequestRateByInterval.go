package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/LogAnalyticsRequestRateByInterval.json
func ExampleLogAnalyticsClient_BeginExportRequestRateByInterval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewLogAnalyticsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginExportRequestRateByInterval(ctx,
		"<location>",
		armcompute.RequestRateByIntervalInput{
			BlobContainerSasURI: to.Ptr("<blob-container-sas-uri>"),
			FromTime:            to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-21T01:54:06.862601Z"); return t }()),
			GroupByResourceName: to.Ptr(true),
			ToTime:              to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-23T01:54:06.862601Z"); return t }()),
			IntervalLength:      to.Ptr(armcompute.IntervalInMinsFiveMins),
		},
		&armcompute.LogAnalyticsClientBeginExportRequestRateByIntervalOptions{ResumeToken: ""})
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
