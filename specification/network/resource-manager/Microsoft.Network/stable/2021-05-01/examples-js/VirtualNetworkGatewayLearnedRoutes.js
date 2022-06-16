const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation retrieves a list of routes the virtual network gateway has learned, including routes learned from BGP peers.
 *
 * @summary This operation retrieves a list of routes the virtual network gateway has learned, including routes learned from BGP peers.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayLearnedRoutes.json
 */
async function getVirtualNetworkGatewayLearnedRoutes() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.beginGetLearnedRoutesAndWait(
    resourceGroupName,
    virtualNetworkGatewayName
  );
  console.log(result);
}

getVirtualNetworkGatewayLearnedRoutes().catch(console.error);
