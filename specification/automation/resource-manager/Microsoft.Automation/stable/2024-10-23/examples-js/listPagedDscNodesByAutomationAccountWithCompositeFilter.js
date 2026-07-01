const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieve a list of dsc nodes.
 *
 * @summary retrieve a list of dsc nodes.
 * x-ms-original-file: 2024-10-23/listPagedDscNodesByAutomationAccountWithCompositeFilter.json
 */
async function listPagedDSCNodesWithFiltersSeparatedByAnd() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee";
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.dscNodeOperations.listByAutomationAccount(
    "rg",
    "myAutomationAccount33",
    {
      filter:
        "properties/extensionHandler/any(eh: eh/version gt '2.70') and contains(name,'sql') and contains(properties/nodeConfiguration/name,'$$Not$$Configured$$')",
      skip: 0,
      top: 10,
      inlinecount: "allpages",
    },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
