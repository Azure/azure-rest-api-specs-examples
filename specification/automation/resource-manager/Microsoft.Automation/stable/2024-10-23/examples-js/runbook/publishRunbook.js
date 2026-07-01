const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to publish runbook draft.
 *
 * @summary publish runbook draft.
 * x-ms-original-file: 2024-10-23/runbook/publishRunbook.json
 */
async function publishRunbookDraft() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee";
  const client = new AutomationClient(credential, subscriptionId);
  await client.runbook.publish("rg", "ContoseAutomationAccount", "Get-AzureVMTutorial");
}
