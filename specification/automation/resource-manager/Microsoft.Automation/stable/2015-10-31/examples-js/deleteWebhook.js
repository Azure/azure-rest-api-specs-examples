const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the webhook by name.
 *
 * @summary Delete the webhook by name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/deleteWebhook.json
 */
async function deleteWebhook() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const webhookName = "TestWebhook";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.webhookOperations.delete(
    resourceGroupName,
    automationAccountName,
    webhookName
  );
  console.log(result);
}
