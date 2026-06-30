const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieve a list of configurations.
 *
 * @summary retrieve a list of configurations.
 * x-ms-original-file: 2024-10-23/getPagedlDscConfigurationsWithNameFilter.json
 */
async function listPagedDSCConfigurationsWithNameFilter() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee";
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.dscConfigurationOperations.listByAutomationAccount(
    "rg",
    "myAutomationAccount33",
    { filter: "contains(name,'server')", skip: 0, top: 2, inlinecount: "allpages" },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
