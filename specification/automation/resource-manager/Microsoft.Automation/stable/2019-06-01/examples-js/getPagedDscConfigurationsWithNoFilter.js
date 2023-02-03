const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of configurations.
 *
 * @summary Retrieve a list of configurations.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/getPagedDscConfigurationsWithNoFilter.json
 */
async function listPagedDscConfigurationsWithNoFilter() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const skip = 0;
  const top = 3;
  const inlinecount = "allpages";
  const options = {
    skip,
    top,
    inlinecount,
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dscConfigurationOperations.listByAutomationAccount(
    resourceGroupName,
    automationAccountName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
