const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Bastion Hosts in a subscription.
 *
 * @summary Lists all Bastion Hosts in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/BastionHostListBySubscription.json
 */
async function listAllBastionHostsForAGivenSubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.bastionHosts.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllBastionHostsForAGivenSubscription().catch(console.error);
