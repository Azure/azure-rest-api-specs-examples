const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the schedule identified by schedule name.
 *
 * @summary Update the schedule identified by schedule name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateSchedule.json
 */
async function updateASchedule() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const scheduleName = "mySchedule";
  const parameters = {
    name: "mySchedule",
    description: "my updated description of schedule goes here",
    isEnabled: false,
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.scheduleOperations.update(
    resourceGroupName,
    automationAccountName,
    scheduleName,
    parameters
  );
  console.log(result);
}
