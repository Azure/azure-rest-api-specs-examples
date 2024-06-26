const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a hybrid runbook worker.
 *
 * @summary Delete a hybrid runbook worker.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/deleteHybridRunbookWorker.json
 */
async function deleteAV2HybridRunbookWorker() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount20";
  const hybridRunbookWorkerGroupName = "myGroup";
  const hybridRunbookWorkerId = "c010ad12-ef14-4a2a-aa9e-ef22c4745ddd";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.hybridRunbookWorkers.delete(
    resourceGroupName,
    automationAccountName,
    hybridRunbookWorkerGroupName,
    hybridRunbookWorkerId
  );
  console.log(result);
}
