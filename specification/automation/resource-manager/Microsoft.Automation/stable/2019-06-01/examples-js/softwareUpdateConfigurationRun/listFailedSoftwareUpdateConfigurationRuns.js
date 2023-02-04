const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return list of software update configuration runs
 *
 * @summary Return list of software update configuration runs
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfigurationRun/listFailedSoftwareUpdateConfigurationRuns.json
 */
async function listSoftwareUpdateConfigurationMachineRunWithStatusEqualToFailed() {
  const subscriptionId =
    process.env["AUTOMATION_SUBSCRIPTION_ID"] || "51766542-3ed7-4a72-a187-0c8ab644ddab";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "mygroup";
  const automationAccountName = "myaccount";
  const filter = "properties/status%20eq%20'Failed'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.softwareUpdateConfigurationRuns.list(
    resourceGroupName,
    automationAccountName,
    options
  );
  console.log(result);
}
