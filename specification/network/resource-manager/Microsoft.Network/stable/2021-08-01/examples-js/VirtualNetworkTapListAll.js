const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the VirtualNetworkTaps in a subscription.
 *
 * @summary Gets all the VirtualNetworkTaps in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkTapListAll.json
 */
async function listAllVirtualNetworkTaps() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworkTaps.listAll()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllVirtualNetworkTaps().catch(console.error);
