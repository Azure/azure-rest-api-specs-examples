const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the protection policies within a resource group.
 *
 * @summary Lists all of the protection policies within a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/WafListPolicies.json
 */
async function listsAllWafPoliciesInAResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webApplicationFirewallPolicies.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsAllWafPoliciesInAResourceGroup().catch(console.error);
