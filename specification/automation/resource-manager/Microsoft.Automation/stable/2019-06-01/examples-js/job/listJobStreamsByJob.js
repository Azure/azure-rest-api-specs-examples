const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of jobs streams identified by job name.
 *
 * @summary Retrieve a list of jobs streams identified by job name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/listJobStreamsByJob.json
 */
async function listJobStreamsByJobName() {
  const subscriptionId =
    process.env["AUTOMATION_SUBSCRIPTION_ID"] || "51766542-3ed7-4a72-a187-0c8ab644ddab";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "mygroup";
  const automationAccountName = "ContoseAutomationAccount";
  const jobName = "foo";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobStreamOperations.listByJob(
    resourceGroupName,
    automationAccountName,
    jobName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
