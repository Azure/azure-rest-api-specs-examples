const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generates VPN client package for P2S client of the virtual network gateway in the specified resource group.
 *
 * @summary Generates VPN client package for P2S client of the virtual network gateway in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/VirtualNetworkGatewayGenerateVpnClientPackage.json
 */
async function generateVpnClientPackage() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.beginGeneratevpnclientpackageAndWait(
    resourceGroupName,
    virtualNetworkGatewayName,
    parameters
  );
  console.log(result);
}
