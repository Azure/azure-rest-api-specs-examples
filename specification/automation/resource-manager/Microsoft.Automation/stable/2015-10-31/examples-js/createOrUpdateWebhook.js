const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create the webhook identified by webhook name.
 *
 * @summary Create the webhook identified by webhook name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/createOrUpdateWebhook.json
 */
async function createOrUpdateWebhook() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const webhookName = "TestWebhook";
  const parameters = {
    name: "TestWebhook",
    expiryTime: new Date("2018-03-29T22:18:13.7002872Z"),
    isEnabled: true,
    runbook: { name: "TestRunbook" },
    uri: "<uri>",
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.webhookOperations.createOrUpdate(
    resourceGroupName,
    automationAccountName,
    webhookName,
    parameters
  );
  console.log(result);
}
