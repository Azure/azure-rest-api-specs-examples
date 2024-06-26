const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all IpGroups in a subscription.
 *
 * @summary Gets all IpGroups in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/IpGroupsListBySubscription.json
 */
async function listIPGroups() {
  const subscriptionId = "subId";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.ipGroups.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listIPGroups().catch(console.error);
