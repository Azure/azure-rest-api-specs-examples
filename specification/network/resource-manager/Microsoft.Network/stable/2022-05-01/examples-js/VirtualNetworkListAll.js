const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all virtual networks in a subscription.
 *
 * @summary Gets all virtual networks in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkListAll.json
 */
async function listAllVirtualNetworks() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworks.listAll()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllVirtualNetworks().catch(console.error);
