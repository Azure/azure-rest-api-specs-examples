const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of sync job streams identified by sync job id.
 *
 * @summary Retrieve a list of sync job streams identified by sync job id.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControlSyncJobStreams/getSourceControlSyncJobStreams.json
 */
async function getAListOfSyncJobStreamsIdentifiedBySyncJobId() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const sourceControlName = "MySourceControl";
  const sourceControlSyncJobId = "ce6fe3e3-9db3-4096-a6b4-82bfb4c10a2b";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sourceControlSyncJobStreams.listBySyncJob(
    resourceGroupName,
    automationAccountName,
    sourceControlName,
    sourceControlSyncJobId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
