const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a virtual wan vpn gateway.
 *
 * @summary Deletes a virtual wan vpn gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/VpnGatewayDelete.json
 */
async function vpnGatewayDelete() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const gatewayName = "gateway1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnGateways.beginDeleteAndWait(resourceGroupName, gatewayName);
  console.log(result);
}
