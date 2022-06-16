const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an express route circuit tags.
 *
 * @summary Updates an express route circuit tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteCircuitUpdateTags.json
 */
async function updateExpressRouteCircuitTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "ertest";
  const circuitName = "er1";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuits.updateTags(
    resourceGroupName,
    circuitName,
    parameters
  );
  console.log(result);
}

updateExpressRouteCircuitTags().catch(console.error);
