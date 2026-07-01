const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create the runbook identified by runbook name.
 *
 * @summary create the runbook identified by runbook name.
 * x-ms-original-file: 2024-10-23/runbook/createRunbookAsDraft.json
 */
async function createRunbookAsDraft() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee";
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.runbookOperations.createOrUpdate(
    "rg",
    "ContoseAutomationAccount",
    "Get-AzureVMTutorial",
    {
      name: "Get-AzureVMTutorial",
      location: "East US 2",
      description: "Description of the Runbook",
      draft: {},
      logProgress: false,
      logVerbose: false,
      runbookType: "PowerShell",
      runtimeEnvironment: "environmentName",
      tags: { tag01: "value01", tag02: "value02" },
    },
  );
  console.log(result);
}
