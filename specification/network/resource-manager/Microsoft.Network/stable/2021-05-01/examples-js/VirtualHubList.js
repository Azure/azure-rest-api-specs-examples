const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the VirtualHubs in a subscription.
 *
 * @summary Lists all the VirtualHubs in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualHubList.json
 */
async function virtualHubList() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualHubs.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualHubList().catch(console.error);
