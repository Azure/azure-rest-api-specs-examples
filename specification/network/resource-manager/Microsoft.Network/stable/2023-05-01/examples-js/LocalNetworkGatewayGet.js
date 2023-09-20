const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified local network gateway in a resource group.
 *
 * @summary Gets the specified local network gateway in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/LocalNetworkGatewayGet.json
 */
async function getLocalNetworkGateway() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const localNetworkGatewayName = "localgw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.localNetworkGateways.get(resourceGroupName, localNetworkGatewayName);
  console.log(result);
}
