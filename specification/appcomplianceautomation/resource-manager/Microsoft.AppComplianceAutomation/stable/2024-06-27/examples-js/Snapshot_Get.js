const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the AppComplianceAutomation snapshot and its properties.
 *
 * @summary Get the AppComplianceAutomation snapshot and its properties.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Snapshot_Get.json
 */
async function snapshotGet() {
  const reportName = "testReportName";
  const snapshotName = "testSnapshot";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.snapshot.get(reportName, snapshotName);
  console.log(result);
}
