const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the job schedule identified by job schedule name.
 *
 * @summary Retrieve the job schedule identified by job schedule name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getJobSchedule.json
 */
async function getAJobSchedule() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "ContoseAutomationAccount";
  const jobScheduleId = "0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.jobScheduleOperations.get(
    resourceGroupName,
    automationAccountName,
    jobScheduleId
  );
  console.log(result);
}
