const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts packet capture on virtual network gateway connection in the specified resource group.
 *
 * @summary Starts packet capture on virtual network gateway connection in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkGatewayConnectionStartPacketCapture.json
 */
async function startPacketCaptureOnVirtualNetworkGatewayConnectionWithoutFilter() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayConnectionName = "vpngwcn1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGatewayConnections.beginStartPacketCaptureAndWait(
    resourceGroupName,
    virtualNetworkGatewayConnectionName
  );
  console.log(result);
}

startPacketCaptureOnVirtualNetworkGatewayConnectionWithoutFilter().catch(console.error);
