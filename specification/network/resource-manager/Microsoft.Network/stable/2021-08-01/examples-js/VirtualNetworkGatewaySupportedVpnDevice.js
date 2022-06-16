const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a xml format representation for supported vpn devices.
 *
 * @summary Gets a xml format representation for supported vpn devices.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkGatewaySupportedVpnDevice.json
 */
async function listVirtualNetworkGatewaySupportedVpnDevices() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.supportedVpnDevices(
    resourceGroupName,
    virtualNetworkGatewayName
  );
  console.log(result);
}

listVirtualNetworkGatewaySupportedVpnDevices().catch(console.error);
