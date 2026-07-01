const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to return list of software update configuration runs
 *
 * @summary return list of software update configuration runs
 * x-ms-original-file: 2024-10-23/softwareUpdateConfigurationRun/listSoftwareUpdateConfigurationRuns.json
 */
async function listSoftwareUpdateConfigurationMachineRuns() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "51766542-3ed7-4a72-a187-0c8ab644ddab";
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.softwareUpdateConfigurationRuns.list("mygroup", "myaccount");
  console.log(result);
}
