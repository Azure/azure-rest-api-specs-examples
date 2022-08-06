const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections created.
 *
 * @summary The List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections created.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkGatewayConnectionsList.json
 */
async function listVirtualNetworkGatewayConnectionsinResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworkGatewayConnections.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVirtualNetworkGatewayConnectionsinResourceGroup().catch(console.error);
