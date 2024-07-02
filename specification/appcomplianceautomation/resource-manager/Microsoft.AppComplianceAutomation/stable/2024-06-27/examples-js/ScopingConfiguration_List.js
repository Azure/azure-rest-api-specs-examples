const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list format of the singleton scopingConfiguration for a specified report.
 *
 * @summary Returns a list format of the singleton scopingConfiguration for a specified report.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/ScopingConfiguration_List.json
 */
async function scopingConfigurationList() {
  const reportName = "testReportName";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const resArray = new Array();
  for await (let item of client.scopingConfiguration.list(reportName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
