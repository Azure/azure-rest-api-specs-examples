const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Undo draft edit to last known published state identified by runbook name.
 *
 * @summary Undo draft edit to last known published state identified by runbook name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/undoDraftEditToLastKnownPublishedState.json
 */
async function undoDraftEditToLastKnownPublishedState() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "ContoseAutomationAccount";
  const runbookName = "Get-AzureVMTutorial";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.runbookDraftOperations.undoEdit(
    resourceGroupName,
    automationAccountName,
    runbookName
  );
  console.log(result);
}
