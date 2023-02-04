const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a variable.
 *
 * @summary Create a variable.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateVariable.json
 */
async function createOrUpdateAVariable() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "sampleAccount9";
  const variableName = "sampleVariable";
  const parameters = {
    name: "sampleVariable",
    description: "my description",
    isEncrypted: false,
    value: '"ComputerName.domain.com"',
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.variableOperations.createOrUpdate(
    resourceGroupName,
    automationAccountName,
    variableName,
    parameters
  );
  console.log(result);
}
