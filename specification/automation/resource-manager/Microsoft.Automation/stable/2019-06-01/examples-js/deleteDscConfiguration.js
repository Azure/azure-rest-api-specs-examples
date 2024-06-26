const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the dsc configuration identified by configuration name.
 *
 * @summary Delete the dsc configuration identified by configuration name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/deleteDscConfiguration.json
 */
async function deleteDscConfiguration() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const configurationName = "TemplateBasic";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.dscConfigurationOperations.delete(
    resourceGroupName,
    automationAccountName,
    configurationName
  );
  console.log(result);
}
