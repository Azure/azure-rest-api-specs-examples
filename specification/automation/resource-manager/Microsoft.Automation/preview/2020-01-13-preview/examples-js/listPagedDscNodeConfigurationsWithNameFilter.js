const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of dsc node configurations.
 *
 * @summary Retrieve a list of dsc node configurations.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodeConfigurationsWithNameFilter.json
 */
async function listPagedDscNodeConfigurationsByAutomationAccountWithNameFilter() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const filter = "contains('.localhost',name)";
  const skip = 0;
  const top = 2;
  const inlinecount = "allpages";
  const options = {
    filter,
    skip,
    top,
    inlinecount,
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dscNodeConfigurationOperations.listByAutomationAccount(
    resourceGroupName,
    automationAccountName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
