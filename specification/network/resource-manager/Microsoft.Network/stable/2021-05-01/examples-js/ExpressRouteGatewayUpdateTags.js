const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates express route gateway tags.
 *
 * @summary Updates express route gateway tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteGatewayUpdateTags.json
 */
async function expressRouteGatewayUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "resourceGroupName";
  const expressRouteGatewayName = "expressRouteGatewayName";
  const expressRouteGatewayParameters = {
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteGateways.beginUpdateTagsAndWait(
    resourceGroupName,
    expressRouteGatewayName,
    expressRouteGatewayParameters
  );
  console.log(result);
}

expressRouteGatewayUpdate().catch(console.error);
