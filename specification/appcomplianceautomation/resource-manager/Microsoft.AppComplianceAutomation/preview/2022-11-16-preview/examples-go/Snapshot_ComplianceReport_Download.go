package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82ea406b73d671269217053d7ef336450d860345/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Snapshot_ComplianceReport_Download.json
func ExampleSnapshotClient_BeginDownload_snapshotDownloadComplianceReport() {
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
		DownloadType:          to.Ptr(armappcomplianceautomation.DownloadTypeComplianceReport),
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
	// 	ComplianceReport: []*armappcomplianceautomation.ComplianceReportItem{
	// 		{
	// 			CategoryName: to.Ptr("Data Security & Privacy"),
	// 			ComplianceState: to.Ptr(armappcomplianceautomation.ComplianceStateHealthy),
	// 			ControlID: to.Ptr("1"),
	// 			ControlName: to.Ptr("Validate that TLS Configuration meets or exceeds the TLS Profile Configuration Requirements"),
	// 			ControlType: to.Ptr(armappcomplianceautomation.ControlTypeFullyAutomated),
	// 			PolicyDescription: to.Ptr("policy description"),
	// 			PolicyDisplayName: to.Ptr("policy name"),
	// 			PolicyID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ResourceGroup: to.Ptr("testGroup"),
	// 			ResourceID: to.Ptr("testResourceId"),
	// 			ResourceType: to.Ptr("storageaccounts"),
	// 			StatusChangeDate: to.Ptr("2021-09-01T17:26:57.4971616Z"),
	// 			SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	}},
	// }
}
