const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a single software update configuration Run by Id.
 *
 * @summary Get a single software update configuration Run by Id.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfigurationRun/getSoftwareUpdateConfigurationRunById.json
 */
async function getSoftwareUpdateConfigurationRunsById() {
  const subscriptionId =
    process.env["AUTOMATION_SUBSCRIPTION_ID"] || "51766542-3ed7-4a72-a187-0c8ab644ddab";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "mygroup";
  const automationAccountName = "myaccount";
  const softwareUpdateConfigurationRunId = "2bd77cfa-2e9c-41b4-a45b-684a77cfeca9";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.softwareUpdateConfigurationRuns.getById(
    resourceGroupName,
    automationAccountName,
    softwareUpdateConfigurationRunId
  );
  console.log(result);
}
