const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an express route cross connection tags.
 *
 * @summary Updates an express route cross connection tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCrossConnectionUpdateTags.json
 */
async function updateExpressRouteCrossConnectionTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "CrossConnection-SiliconValley";
  const crossConnectionName = "<circuitServiceKey>";
  const crossConnectionParameters = {
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCrossConnections.updateTags(
    resourceGroupName,
    crossConnectionName,
    crossConnectionParameters
  );
  console.log(result);
}

updateExpressRouteCrossConnectionTags().catch(console.error);
