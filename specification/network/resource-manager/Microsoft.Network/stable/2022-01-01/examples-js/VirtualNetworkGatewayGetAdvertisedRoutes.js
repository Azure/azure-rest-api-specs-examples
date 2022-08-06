const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation retrieves a list of routes the virtual network gateway is advertising to the specified peer.
 *
 * @summary This operation retrieves a list of routes the virtual network gateway is advertising to the specified peer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkGatewayGetAdvertisedRoutes.json
 */
async function getVirtualNetworkGatewayAdvertisedRoutes() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const peer = "test";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.beginGetAdvertisedRoutesAndWait(
    resourceGroupName,
    virtualNetworkGatewayName,
    peer
  );
  console.log(result);
}

getVirtualNetworkGatewayAdvertisedRoutes().catch(console.error);
