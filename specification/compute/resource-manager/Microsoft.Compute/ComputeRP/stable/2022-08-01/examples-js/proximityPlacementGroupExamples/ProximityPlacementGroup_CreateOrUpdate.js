const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a proximity placement group.
 *
 * @summary Create or update a proximity placement group.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/proximityPlacementGroupExamples/ProximityPlacementGroup_CreateOrUpdate.json
 */
async function createOrUpdateAProximityPlacementGroup() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const proximityPlacementGroupName = "myProximityPlacementGroup";
  const parameters = {
    intent: { vmSizes: ["Basic_A0", "Basic_A2"] },
    location: "westus",
    proximityPlacementGroupType: "Standard",
    zones: ["1"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.proximityPlacementGroups.createOrUpdate(
    resourceGroupName,
    proximityPlacementGroupName,
    parameters
  );
  console.log(result);
}

createOrUpdateAProximityPlacementGroup().catch(console.error);
