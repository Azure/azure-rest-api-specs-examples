const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all available managed rule sets.
 *
 * @summary Lists all available managed rule sets.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-11-01/examples/WafListManagedRuleSets.json
 */
async function listPoliciesInAResourceGroup() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedRuleSets.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPoliciesInAResourceGroup().catch(console.error);
