const { NetworkManagementClient } = require("@azure/arm-network-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified virtual network gateway connection by resource group.
 *
 * @summary Gets the specified virtual network gateway connection by resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2018-11-01/examples/VirtualNetworkGatewayConnectionGet.json
 */
async function getVirtualNetworkGatewayConnection() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkGatewayConnectionName = "connS2S";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGatewayConnections.get(
    resourceGroupName,
    virtualNetworkGatewayConnectionName
  );
  console.log(result);
}
