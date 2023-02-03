const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a job of the runbook.
 *
 * @summary Create a job of the runbook.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/createJob.json
 */
async function createJob() {
  const subscriptionId =
    process.env["AUTOMATION_SUBSCRIPTION_ID"] || "51766542-3ed7-4a72-a187-0c8ab644ddab";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "mygroup";
  const automationAccountName = "ContoseAutomationAccount";
  const jobName = "foo";
  const parameters = {
    parameters: { key01: "value01", key02: "value02" },
    runOn: "",
    runbook: { name: "TestRunbook" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.jobOperations.create(
    resourceGroupName,
    automationAccountName,
    jobName,
    parameters
  );
  console.log(result);
}
