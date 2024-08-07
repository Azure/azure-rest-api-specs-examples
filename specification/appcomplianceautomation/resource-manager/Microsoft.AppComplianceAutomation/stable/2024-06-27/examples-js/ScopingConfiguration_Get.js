const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the AppComplianceAutomation scoping configuration of the specific report.
 *
 * @summary Get the AppComplianceAutomation scoping configuration of the specific report.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/ScopingConfiguration_Get.json
 */
async function scopingConfiguration() {
  const reportName = "testReportName";
  const scopingConfigurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.scopingConfiguration.get(reportName, scopingConfigurationName);
  console.log(result);
}
