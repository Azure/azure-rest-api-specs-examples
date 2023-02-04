const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Publish runbook draft.
 *
 * @summary Publish runbook draft.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/publishRunbook.json
 */
async function publishRunbookDraft() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "ContoseAutomationAccount";
  const runbookName = "Get-AzureVMTutorial";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.runbookOperations.beginPublishAndWait(
    resourceGroupName,
    automationAccountName,
    runbookName
  );
  console.log(result);
}
