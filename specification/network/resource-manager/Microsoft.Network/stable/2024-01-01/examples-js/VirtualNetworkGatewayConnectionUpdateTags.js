const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a virtual network gateway connection tags.
 *
 * @summary Updates a virtual network gateway connection tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/VirtualNetworkGatewayConnectionUpdateTags.json
 */
async function updateVirtualNetworkGatewayConnectionTags() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkGatewayConnectionName = "test";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGatewayConnections.beginUpdateTagsAndWait(
    resourceGroupName,
    virtualNetworkGatewayConnectionName,
    parameters,
  );
  console.log(result);
}
