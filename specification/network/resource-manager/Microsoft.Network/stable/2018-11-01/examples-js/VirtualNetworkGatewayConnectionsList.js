const { NetworkManagementClient } = require("@azure/arm-network-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections created.
 *
 * @summary The List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections created.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2018-11-01/examples/VirtualNetworkGatewayConnectionsList.json
 */
async function listVirtualNetworkGatewayConnectionsinResourceGroup() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworkGatewayConnections.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
