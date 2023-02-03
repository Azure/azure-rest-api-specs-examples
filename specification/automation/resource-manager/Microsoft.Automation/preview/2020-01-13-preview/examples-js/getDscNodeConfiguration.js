const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the Dsc node configurations by node configuration.
 *
 * @summary Retrieve the Dsc node configurations by node configuration.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNodeConfiguration.json
 */
async function getADscNodeConfiguration() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const nodeConfigurationName = "SetupServer.localhost";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.dscNodeConfigurationOperations.get(
    resourceGroupName,
    automationAccountName,
    nodeConfigurationName
  );
  console.log(result);
}
