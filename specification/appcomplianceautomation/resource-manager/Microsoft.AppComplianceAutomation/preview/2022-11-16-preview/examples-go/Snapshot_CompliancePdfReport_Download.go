package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Snapshot_CompliancePdfReport_Download.json
func ExampleSnapshotClient_BeginDownload_snapshotDownloadCompliancePdfReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappcomplianceautomation.NewSnapshotClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDownload(ctx, "testReportName", "testSnapshotName", armappcomplianceautomation.SnapshotDownloadRequest{
		DownloadType:          to.Ptr(armappcomplianceautomation.DownloadTypeCompliancePDFReport),
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
	// TODO: use response item
	_ = res
}
