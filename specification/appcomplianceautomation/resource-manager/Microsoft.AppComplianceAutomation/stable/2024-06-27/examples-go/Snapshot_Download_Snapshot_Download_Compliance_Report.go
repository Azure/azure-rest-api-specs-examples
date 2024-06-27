package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Snapshot_Download_Snapshot_Download_Compliance_Report.json
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
	poller, err := clientFactory.NewSnapshotClient().BeginDownload(ctx, "testReportName", "testSnapshotName", armappcomplianceautomation.SnapshotDownloadRequest{}, nil)
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
	// 			ControlFamilyName: to.Ptr("Incident Response"),
	// 			ControlID: to.Ptr("Operational_Security_75"),
	// 			ControlName: to.Ptr("Provide demonstrable evidence that all member of the incident response team have completed annual training or a table top exercise"),
	// 			ControlStatus: to.Ptr(armappcomplianceautomation.ControlStatusPassed),
	// 			ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
	// 			ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAzure),
	// 			ResourceStatus: to.Ptr(armappcomplianceautomation.ResourceStatusHealthy),
	// 			ResourceStatusChangeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-12T16:17:20.150Z"); return t}()),
	// 			ResourceType: to.Ptr("Microsoft.SignalRService/SignalR"),
	// 			ResponsibilityDescription: to.Ptr("Restrict access to the Kubernetes Service Management API by granting API access only to IP addresses in specific ranges. It is recommended to limit access to authorized IP ranges to ensure that only applications from allowed networks can access the cluster."),
	// 			ResponsibilityTitle: to.Ptr("Authorized IP ranges should be defined on Kubernetes Services"),
	// 	}},
	// }
}
