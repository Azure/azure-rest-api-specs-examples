const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a paginated list of evidences for a specified report.
 *
 * @summary Returns a paginated list of evidences for a specified report.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Evidence_ListByReport.json
 */
async function evidenceListByReport() {
  const reportName = "reportName";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const resArray = new Array();
  for await (let item of client.evidence.listByReport(reportName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
