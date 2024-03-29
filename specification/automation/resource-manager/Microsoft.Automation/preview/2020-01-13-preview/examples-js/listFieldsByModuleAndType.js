const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of fields of a given type identified by module name.
 *
 * @summary Retrieve a list of fields of a given type identified by module name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listFieldsByModuleAndType.json
 */
async function getAListOfFieldsOfAGivenType() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "MyAutomationAccount";
  const moduleName = "MyModule";
  const typeName = "MyCustomType";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fields.listByType(
    resourceGroupName,
    automationAccountName,
    moduleName,
    typeName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
