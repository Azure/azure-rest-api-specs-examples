const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates virtual wan p2s vpn gateway tags.
 *
 * @summary Updates virtual wan p2s vpn gateway tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/P2SVpnGatewayUpdateTags.json
 */
async function p2SVpnGatewayUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "p2sVpnGateway1";
  const p2SVpnGatewayParameters = {
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.p2SVpnGateways.beginUpdateTagsAndWait(
    resourceGroupName,
    gatewayName,
    p2SVpnGatewayParameters
  );
  console.log(result);
}

p2SVpnGatewayUpdate().catch(console.error);
