const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the usage for the account id.
 *
 * @summary Retrieve the usage for the account id.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/getUsagesOfAutomationAccount.json
 */
async function getUsagesOfAnAutomationAccount() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount11";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usages.listByAutomationAccount(
    resourceGroupName,
    automationAccountName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
