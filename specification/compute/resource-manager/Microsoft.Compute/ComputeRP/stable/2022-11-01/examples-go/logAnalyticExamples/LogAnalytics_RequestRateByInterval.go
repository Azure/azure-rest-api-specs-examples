package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/logAnalyticExamples/LogAnalytics_RequestRateByInterval.json
func ExampleLogAnalyticsClient_BeginExportRequestRateByInterval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewLogAnalyticsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginExportRequestRateByInterval(ctx, "westus", armcompute.RequestRateByIntervalInput{
		BlobContainerSasURI: to.Ptr("https://somesasuri"),
		FromTime:            to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-21T01:54:06.862601Z"); return t }()),
		GroupByResourceName: to.Ptr(true),
		ToTime:              to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-23T01:54:06.862601Z"); return t }()),
		IntervalLength:      to.Ptr(armcompute.IntervalInMinsFiveMins),
	}, nil)
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
