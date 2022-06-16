const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates virtual wan vpn gateway tags.
 *
 * @summary Updates virtual wan vpn gateway tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VpnGatewayUpdateTags.json
 */
async function vpnGatewayUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "gateway1";
  const vpnGatewayParameters = {
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnGateways.beginUpdateTagsAndWait(
    resourceGroupName,
    gatewayName,
    vpnGatewayParameters
  );
  console.log(result);
}

vpnGatewayUpdate().catch(console.error);
