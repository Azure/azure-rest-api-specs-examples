const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified virtual network gateway by resource group.
 *
 * @summary Gets the specified virtual network gateway by resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkGatewayGet.json
 */
async function getVirtualNetworkGateway() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.get(
    resourceGroupName,
    virtualNetworkGatewayName
  );
  console.log(result);
}

getVirtualNetworkGateway().catch(console.error);
