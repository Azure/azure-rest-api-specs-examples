const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update ExpressRoutePort tags.
 *
 * @summary Update ExpressRoutePort tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRoutePortUpdateTags.json
 */
async function expressRoutePortUpdateTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const expressRoutePortName = "portName";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRoutePorts.updateTags(
    resourceGroupName,
    expressRoutePortName,
    parameters
  );
  console.log(result);
}

expressRoutePortUpdateTags().catch(console.error);
