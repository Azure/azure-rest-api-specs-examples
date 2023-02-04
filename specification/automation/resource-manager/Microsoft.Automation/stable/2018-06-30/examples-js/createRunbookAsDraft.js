const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create the runbook identified by runbook name.
 *
 * @summary Create the runbook identified by runbook name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/createRunbookAsDraft.json
 */
async function createRunbookAsDraft() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "ContoseAutomationAccount";
  const runbookName = "Get-AzureVMTutorial";
  const parameters = {
    name: "Get-AzureVMTutorial",
    description: "Description of the Runbook",
    draft: {},
    location: "East US 2",
    logProgress: false,
    logVerbose: false,
    runbookType: "PowerShellWorkflow",
    tags: { tag01: "value01", tag02: "value02" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.runbookOperations.createOrUpdate(
    resourceGroupName,
    automationAccountName,
    runbookName,
    parameters
  );
  console.log(result);
}
