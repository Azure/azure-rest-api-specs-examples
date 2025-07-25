package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/logAnalyticExamples/LogAnalytics_RequestRateByInterval.json
func ExampleLogAnalyticsClient_BeginExportRequestRateByInterval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLogAnalyticsClient().BeginExportRequestRateByInterval(ctx, "westus", armcompute.RequestRateByIntervalInput{
		BlobContainerSasURI: to.Ptr("https://somesasuri"),
		FromTime:            to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-21T01:54:06.862Z"); return t }()),
		GroupByResourceName: to.Ptr(true),
		ToTime:              to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-23T01:54:06.862Z"); return t }()),
		IntervalLength:      to.Ptr(armcompute.IntervalInMinsFiveMins),
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
	// res.LogAnalyticsOperationResult = armcompute.LogAnalyticsOperationResult{
	// 	Properties: &armcompute.LogAnalyticsOutput{
	// 		Output: to.Ptr("https://crptestar4227.blob.core.windows.net:443/sascontainer/RequestRateByInterval_20180121-0154_20180123-0154.csv"),
	// 	},
	// }
}
