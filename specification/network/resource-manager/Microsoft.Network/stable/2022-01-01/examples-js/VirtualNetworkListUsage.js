const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists usage stats.
 *
 * @summary Lists usage stats.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkListUsage.json
 */
async function vnetGetUsage() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkName = "vnetName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworks.listUsage(resourceGroupName, virtualNetworkName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

vnetGetUsage().catch(console.error);
