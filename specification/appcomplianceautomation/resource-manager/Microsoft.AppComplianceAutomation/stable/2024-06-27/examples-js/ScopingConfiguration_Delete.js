const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Clean the AppComplianceAutomation scoping configuration of the specific report.
 *
 * @summary Clean the AppComplianceAutomation scoping configuration of the specific report.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/ScopingConfiguration_Delete.json
 */
async function scopingConfigurationDelete() {
  const reportName = "testReportName";
  const scopingConfigurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.scopingConfiguration.delete(reportName, scopingConfigurationName);
  console.log(result);
}
