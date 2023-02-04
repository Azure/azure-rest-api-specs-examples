const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resume the watcher identified by watcher name.
 *
 * @summary Resume the watcher identified by watcher name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/stopWatcher.json
 */
async function startWatcher() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "MyTestAutomationAccount";
  const watcherName = "MyTestWatcher";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.watcherOperations.stop(
    resourceGroupName,
    automationAccountName,
    watcherName
  );
  console.log(result);
}
