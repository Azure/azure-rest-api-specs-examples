const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of source control sync jobs.
 *
 * @summary Retrieve a list of source control sync jobs.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControlSyncJob/getAllSourceControlSyncJobs.json
 */
async function getAListOfSourceControlSyncJobs() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const sourceControlName = "MySourceControl";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sourceControlSyncJobOperations.listByAutomationAccount(
    resourceGroupName,
    automationAccountName,
    sourceControlName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
