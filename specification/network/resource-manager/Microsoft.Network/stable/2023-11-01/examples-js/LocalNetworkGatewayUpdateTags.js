const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a local network gateway tags.
 *
 * @summary Updates a local network gateway tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/LocalNetworkGatewayUpdateTags.json
 */
async function updateLocalNetworkGatewayTags() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const localNetworkGatewayName = "lgw";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.localNetworkGateways.updateTags(
    resourceGroupName,
    localNetworkGatewayName,
    parameters,
  );
  console.log(result);
}
