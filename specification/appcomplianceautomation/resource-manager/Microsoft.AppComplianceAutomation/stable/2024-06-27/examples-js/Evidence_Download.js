const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Download evidence file.
 *
 * @summary Download evidence file.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Evidence_Download.json
 */
async function evidenceDownload() {
  const reportName = "testReportName";
  const evidenceName = "evidence1";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.evidence.download(reportName, evidenceName, {
    offerGuid: "00000000-0000-0000-0000-000000000000",
    reportCreatorTenantId: "00000000-0000-0000-0000-000000000000",
  });
  console.log(result);
}
