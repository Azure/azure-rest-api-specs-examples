const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve all the job streams for the compilation Job.
 *
 * @summary Retrieve all the job streams for the compilation Job.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/compilationJobStreamList.json
 */
async function listDscCompilationJobStreams() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const jobId = "836d4e06-2d88-46b4-8500-7febd4906838";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.dscCompilationJobStream.listByJob(
    resourceGroupName,
    automationAccountName,
    jobId
  );
  console.log(result);
}
