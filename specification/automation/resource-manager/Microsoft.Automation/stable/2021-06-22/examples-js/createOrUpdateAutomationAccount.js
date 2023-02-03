const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update automation account.
 *
 * @summary Create or update automation account.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/createOrUpdateAutomationAccount.json
 */
async function createOrUpdateAutomationAccount() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount9";
  const parameters = {
    name: "myAutomationAccount9",
    location: "East US 2",
    sku: { name: "Free" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.automationAccountOperations.createOrUpdate(
    resourceGroupName,
    automationAccountName,
    parameters
  );
  console.log(result);
}
