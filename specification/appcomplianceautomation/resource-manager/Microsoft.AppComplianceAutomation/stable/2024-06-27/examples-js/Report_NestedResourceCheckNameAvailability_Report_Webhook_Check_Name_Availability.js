const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the report's nested resource name availability, e.g: Webhooks, Evidences, Snapshots.
 *
 * @summary Checks the report's nested resource name availability, e.g: Webhooks, Evidences, Snapshots.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_NestedResourceCheckNameAvailability_Report_Webhook_Check_Name_Availability.json
 */
async function reportWebhookCheckNameAvailability() {
  const reportName = "reportABC";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.report.nestedResourceCheckNameAvailability(reportName, {
    name: "webhookABC",
    type: "Microsoft.AppComplianceAutomation/reports/webhooks",
  });
  console.log(result);
}
