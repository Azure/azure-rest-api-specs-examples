const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Download compliance needs from snapshot, like: Compliance Report, Resource List.
 *
 * @summary Download compliance needs from snapshot, like: Compliance Report, Resource List.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Snapshot_Download_Snapshot_Download_Resource_List.json
 */
async function snapshotDownloadResourceList() {
  const reportName = "testReportName";
  const snapshotName = "testSnapshotName";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.snapshot.beginDownloadAndWait(reportName, snapshotName, {
    downloadType: "ResourceList",
    offerGuid: "00000000-0000-0000-0000-000000000001",
    reportCreatorTenantId: "00000000-0000-0000-0000-000000000000",
  });
  console.log(result);
}
