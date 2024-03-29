const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of hybrid runbook workers.
 *
 * @summary Retrieve a list of hybrid runbook workers.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/listHybridRunbookWorker.json
 */
async function listHybridWorkersByHybridRunbookWorkerGroup() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "testaccount";
  const hybridRunbookWorkerGroupName = "TestHybridGroup";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hybridRunbookWorkers.listByHybridRunbookWorkerGroup(
    resourceGroupName,
    automationAccountName,
    hybridRunbookWorkerGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
