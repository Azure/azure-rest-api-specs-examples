const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a proximity placement group.
 *
 * @summary Update a proximity placement group.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/PatchAProximityPlacementGroup.json
 */
async function createAProximityPlacementGroup() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const proximityPlacementGroupName = "myProximityPlacementGroup";
  const parameters = {
    tags: { additionalProp1: "string" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.proximityPlacementGroups.update(
    resourceGroupName,
    proximityPlacementGroupName,
    parameters
  );
  console.log(result);
}

createAProximityPlacementGroup().catch(console.error);
