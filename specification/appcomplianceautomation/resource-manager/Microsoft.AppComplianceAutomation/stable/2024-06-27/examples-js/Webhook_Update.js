const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an exiting AppComplianceAutomation webhook.
 *
 * @summary Update an exiting AppComplianceAutomation webhook.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Webhook_Update.json
 */
async function webhookUpdate() {
  const reportName = "testReportName";
  const webhookName = "testWebhookName";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.webhook.update(reportName, webhookName, {
    properties: {
      contentType: "application/json",
      enableSslVerification: "true",
      events: ["generate_snapshot_failed"],
      payloadUrl: "https://example.com",
      sendAllEvents: "false",
      status: "Enabled",
      updateWebhookKey: "true",
      webhookKey: "00000000-0000-0000-0000-000000000000",
    },
  });
  console.log(result);
}
