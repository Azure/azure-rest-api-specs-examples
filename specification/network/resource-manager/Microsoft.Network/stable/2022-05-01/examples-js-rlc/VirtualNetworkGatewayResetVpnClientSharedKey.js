const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Resets the VPN client shared key of the virtual network gateway in the specified resource group.
 *
 * @summary Resets the VPN client shared key of the virtual network gateway in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkGatewayResetVpnClientSharedKey.json
 */
async function resetVpnClientSharedKey() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const options = {
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/resetvpnclientsharedkey",
      subscriptionId,
      resourceGroupName,
      virtualNetworkGatewayName,
    )
    .post(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

resetVpnClientSharedKey().catch(console.error);
