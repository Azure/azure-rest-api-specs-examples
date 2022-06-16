const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generates VPN profile for P2S client of the virtual network gateway in the specified resource group. Used for IKEV2 and radius based authentication.
 *
 * @summary Generates VPN profile for P2S client of the virtual network gateway in the specified resource group. Used for IKEV2 and radius based authentication.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayGenerateVpnProfile.json
 */
async function generateVirtualNetworkGatewayVpnProfile() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.beginGenerateVpnProfileAndWait(
    resourceGroupName,
    virtualNetworkGatewayName,
    parameters
  );
  console.log(result);
}

generateVirtualNetworkGatewayVpnProfile().catch(console.error);
