const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a single software update configuration by name.
 *
 * @summary Get a single software update configuration by name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/getSoftwareUpdateConfigurationByName.json
 */
async function getSoftwareUpdateConfigurationByName() {
  const subscriptionId =
    process.env["AUTOMATION_SUBSCRIPTION_ID"] || "51766542-3ed7-4a72-a187-0c8ab644ddab";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "mygroup";
  const automationAccountName = "myaccount";
  const softwareUpdateConfigurationName = "mypatch";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.softwareUpdateConfigurations.getByName(
    resourceGroupName,
    automationAccountName,
    softwareUpdateConfigurationName
  );
  console.log(result);
}
