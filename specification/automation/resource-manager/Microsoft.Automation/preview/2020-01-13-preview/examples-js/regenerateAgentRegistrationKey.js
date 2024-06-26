const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate a primary or secondary agent registration key
 *
 * @summary Regenerate a primary or secondary agent registration key
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/regenerateAgentRegistrationKey.json
 */
async function regenerateRegistrationKey() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount18";
  const parameters = {
    keyName: "primary",
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.agentRegistrationInformation.regenerateKey(
    resourceGroupName,
    automationAccountName,
    parameters
  );
  console.log(result);
}
