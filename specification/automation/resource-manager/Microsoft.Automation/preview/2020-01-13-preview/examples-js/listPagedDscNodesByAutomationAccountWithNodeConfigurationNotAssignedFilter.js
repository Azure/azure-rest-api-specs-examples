const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of dsc nodes.
 *
 * @summary Retrieve a list of dsc nodes.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodesByAutomationAccountWithNodeConfigurationNotAssignedFilter.json
 */
async function listPagedDscNodesByAutomationAccountWhereNodeConfigurationsAreNotAssignedFilter() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const filter = "properties/nodeConfiguration/name eq ''";
  const skip = 0;
  const top = 20;
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
  for await (let item of client.dscNodeOperations.listByAutomationAccount(
    resourceGroupName,
    automationAccountName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
