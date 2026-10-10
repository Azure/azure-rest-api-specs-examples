package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-10-01/DrillRuns_ListReportDownloadUrl_MaximumSet_Gen.json
func ExampleDrillRunsClient_BeginListReportDownloadURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDrillRunsClient().BeginListReportDownloadURL(ctx, "sampleServiceGroupName", "3f2b1c4d-5e6f-4a7b-8c9d-0e1f2a3b4c5d", "drill1", "ca92602e-53bf-43d2-ae62-d3fc940474b3", armresiliencemanagement.ListReportDownloadURLRequest{
		Format: to.Ptr(armresiliencemanagement.DrillReportFormatHTML),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresiliencemanagement.DrillRunsClientListReportDownloadURLResponse{
	// 	ListReportDownloadURLResponse: armresiliencemanagement.ListReportDownloadURLResponse{
	// 		Format: to.Ptr(armresiliencemanagement.DrillReportFormatHTML),
	// 		DownloadURL: to.Ptr("https://contoso.blob.core.windows.net/drill-reports/drill1/ca92602e-53bf-43d2-ae62-d3fc940474b3/report.html?sv=2025-01-05&sr=b&sig=redacted&se=2026-08-31T10%3A45%3A00Z&sp=r"),
	// 		ExpiryTimestamp: to.Ptr(time.Date(2026, time.August, 31, 10, 45, 0, 0, time.UTC)),
	// 	},
	// }
}
