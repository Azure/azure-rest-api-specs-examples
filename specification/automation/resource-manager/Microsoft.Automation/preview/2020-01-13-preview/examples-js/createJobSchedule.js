const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a job schedule.
 *
 * @summary Create a job schedule.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createJobSchedule.json
 */
async function createAJobSchedule() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "ContoseAutomationAccount";
  const jobScheduleId = "0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc";
  const parameters = {
    parameters: {
      jobscheduletag01: "jobschedulevalue01",
      jobscheduletag02: "jobschedulevalue02",
    },
    runbook: { name: "TestRunbook" },
    schedule: {
      name: "ScheduleNameGoesHere332204b5-debe-4348-a5c7-6357457189f2",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.jobScheduleOperations.create(
    resourceGroupName,
    automationAccountName,
    jobScheduleId,
    parameters
  );
  console.log(result);
}
