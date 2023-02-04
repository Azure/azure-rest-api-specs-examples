const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve deleted automation account.
 *
 * @summary Retrieve deleted automation account.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-01-31/examples/getDeletedAutomationAccount.json
 */
async function getDeletedAutomationAccount() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.deletedAutomationAccounts.listBySubscription();
  console.log(result);
}
