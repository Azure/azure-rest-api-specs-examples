const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the Virtual Routers in a subscription.
 *
 * @summary Gets all the Virtual Routers in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/VirtualRouterListBySubscription.json
 */
async function listAllVirtualRoutersForAGivenSubscription() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualRouters.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
