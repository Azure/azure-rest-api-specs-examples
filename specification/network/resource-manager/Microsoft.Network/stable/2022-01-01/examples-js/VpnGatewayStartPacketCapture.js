const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts packet capture on vpn gateway in the specified resource group.
 *
 * @summary Starts packet capture on vpn gateway in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnGatewayStartPacketCapture.json
 */
async function startPacketCaptureOnVpnGatewayWithoutFilter() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "vpngw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnGateways.beginStartPacketCaptureAndWait(
    resourceGroupName,
    gatewayName
  );
  console.log(result);
}

startPacketCaptureOnVpnGatewayWithoutFilter().catch(console.error);
