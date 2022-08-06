const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resets the VPN client shared key of the virtual network gateway in the specified resource group.
 *
 * @summary Resets the VPN client shared key of the virtual network gateway in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkGatewayResetVpnClientSharedKey.json
 */
async function resetVpnClientSharedKey() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.beginResetVpnClientSharedKeyAndWait(
    resourceGroupName,
    virtualNetworkGatewayName
  );
  console.log(result);
}

resetVpnClientSharedKey().catch(console.error);
