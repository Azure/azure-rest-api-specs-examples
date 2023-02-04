const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the runbook identified by runbook name.
 *
 * @summary Update the runbook identified by runbook name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/updateRunbook.json
 */
async function updateRunbook() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "ContoseAutomationAccount";
  const runbookName = "Get-AzureVMTutorial";
  const parameters = {
    description: "Updated Description of the Runbook",
    logActivityTrace: 1,
    logProgress: true,
    logVerbose: false,
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.runbookOperations.update(
    resourceGroupName,
    automationAccountName,
    runbookName,
    parameters
  );
  console.log(result);
}
