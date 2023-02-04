const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resume the test job.
 *
 * @summary Resume the test job.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/resumeTestJob.json
 */
async function resumeTestJob() {
  const subscriptionId =
    process.env["AUTOMATION_SUBSCRIPTION_ID"] || "51766542-3ed7-4a72-a187-0c8ab644ddab";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "mygroup";
  const automationAccountName = "ContoseAutomationAccount";
  const runbookName = "Get-AzureVMTutorial";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.testJobOperations.resume(
    resourceGroupName,
    automationAccountName,
    runbookName
  );
  console.log(result);
}
