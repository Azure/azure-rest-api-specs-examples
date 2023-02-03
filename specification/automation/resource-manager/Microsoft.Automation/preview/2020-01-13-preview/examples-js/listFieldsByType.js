const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of fields of a given type across all accessible modules.
 *
 * @summary Retrieve a list of fields of a given type across all accessible modules.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listFieldsByType.json
 */
async function getAListOfFieldsOfAGivenTypeAcrossAllAccessibleModules() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "MyAutomationAccount";
  const typeName = "MyCustomType";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.objectDataTypes.listFieldsByType(
    resourceGroupName,
    automationAccountName,
    typeName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
