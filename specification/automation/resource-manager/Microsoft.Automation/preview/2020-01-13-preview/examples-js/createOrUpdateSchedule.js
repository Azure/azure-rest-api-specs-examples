const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a schedule.
 *
 * @summary Create a schedule.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateSchedule.json
 */
async function createOrUpdateASchedule() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const scheduleName = "mySchedule";
  const parameters = {
    name: "mySchedule",
    description: "my description of schedule goes here",
    advancedSchedule: {},
    expiryTime: new Date("2017-04-01T17:28:57.2494819Z"),
    frequency: "Hour",
    interval: 1,
    startTime: new Date("2017-03-27T17:28:57.2494819Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.scheduleOperations.createOrUpdate(
    resourceGroupName,
    automationAccountName,
    scheduleName,
    parameters
  );
  console.log(result);
}
