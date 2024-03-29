const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the webhook identified by webhook name.
 *
 * @summary Update the webhook identified by webhook name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/updateWebhook.json
 */
async function updateWebhook() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const webhookName = "TestWebhook";
  const parameters = {
    name: "TestWebhook",
    description: "updated webhook",
    isEnabled: false,
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.webhookOperations.update(
    resourceGroupName,
    automationAccountName,
    webhookName,
    parameters
  );
  console.log(result);
}
