const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a hybrid runbook worker group.
 *
 * @summary Delete a hybrid runbook worker group.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-02-22/examples/deleteHybridRunbookWorkerGroup.json
 */
async function deleteAHybridWorkerGroup() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount20";
  const hybridRunbookWorkerGroupName = "myGroup";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.hybridRunbookWorkerGroupOperations.delete(
    resourceGroupName,
    automationAccountName,
    hybridRunbookWorkerGroupName
  );
  console.log(result);
}
