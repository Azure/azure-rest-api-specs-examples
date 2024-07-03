const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the AppComplianceAutomation snapshot list.
 *
 * @summary Get the AppComplianceAutomation snapshot list.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Snapshot_List.json
 */
async function snapshotList() {
  const skipToken = "1";
  const top = 100;
  const offerGuid = "00000000-0000-0000-0000-000000000001";
  const reportCreatorTenantId = "00000000-0000-0000-0000-000000000000";
  const reportName = "testReportName";
  const options = {
    skipToken,
    top,
    offerGuid,
    reportCreatorTenantId,
  };
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const resArray = new Array();
  for await (let item of client.snapshot.list(reportName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
