const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts packet capture on virtual network gateway in the specified resource group.
 *
 * @summary Starts packet capture on virtual network gateway in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/VirtualNetworkGatewayStartPacketCapture.json
 */
async function startPacketCaptureOnVirtualNetworkGatewayWithoutFilter() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.beginStartPacketCaptureAndWait(
    resourceGroupName,
    virtualNetworkGatewayName,
  );
  console.log(result);
}
