const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of activities in the module identified by module name.
 *
 * @summary Retrieve a list of activities in the module identified by module name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listActivitiesByModule.json
 */
async function listActivitiesByAModule() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const moduleName = "OmsCompositeResources";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.activityOperations.listByModule(
    resourceGroupName,
    automationAccountName,
    moduleName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
