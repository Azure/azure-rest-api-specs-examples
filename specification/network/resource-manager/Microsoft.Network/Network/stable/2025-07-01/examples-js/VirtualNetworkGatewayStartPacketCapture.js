const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to starts packet capture on virtual network gateway in the specified resource group.
 *
 * @summary starts packet capture on virtual network gateway in the specified resource group.
 * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayStartPacketCapture.json
 */
async function startPacketCaptureOnVirtualNetworkGatewayWithoutFilter() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.startPacketCapture("rg1", "vpngw");
  console.log(result);
}
