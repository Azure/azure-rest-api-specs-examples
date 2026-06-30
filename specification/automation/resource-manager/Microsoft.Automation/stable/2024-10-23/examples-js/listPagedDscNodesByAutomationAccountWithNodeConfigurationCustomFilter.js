const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieve a list of dsc nodes.
 *
 * @summary retrieve a list of dsc nodes.
 * x-ms-original-file: 2024-10-23/listPagedDscNodesByAutomationAccountWithNodeConfigurationCustomFilter.json
 */
async function listPagedDSCNodesByAutomationAccountWithNodeConfigurationCustomFilter() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee";
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.dscNodeOperations.listByAutomationAccount(
    "rg",
    "myAutomationAccount33",
    {
      filter:
        "contains(properties/nodeConfiguration/name,'SetupServer.localhost,SetupClient.localhost,$$Not$$Configured$$')",
      skip: 0,
      top: 4,
      inlinecount: "allpages",
    },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
