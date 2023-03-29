package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82ea406b73d671269217053d7ef336450d860345/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Snapshot_ComplianceDetailedPdfReport_Download.json
func ExampleSnapshotClient_BeginDownload_snapshotDownloadComplianceDetailedPdfReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSnapshotClient().BeginDownload(ctx, "testReportName", "testSnapshotName", armappcomplianceautomation.SnapshotDownloadRequest{
		DownloadType:          to.Ptr(armappcomplianceautomation.DownloadTypeComplianceDetailedPDFReport),
		OfferGUID:             to.Ptr("00000000-0000-0000-0000-000000000000"),
		ReportCreatorTenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
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
	// res.DownloadResponse = armappcomplianceautomation.DownloadResponse{
	// 	ComplianceDetailedPDFReport: &armappcomplianceautomation.DownloadResponseComplianceDetailedPDFReport{
	// 		SasURI: to.Ptr("this is a uri"),
	// 	},
	// }
}
