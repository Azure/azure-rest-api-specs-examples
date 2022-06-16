const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves information about a proximity placement group .
 *
 * @summary Retrieves information about a proximity placement group .
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetAProximityPlacementGroup.json
 */
async function createAProximityPlacementGroup() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const proximityPlacementGroupName = "myProximityPlacementGroup";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.proximityPlacementGroups.get(
    resourceGroupName,
    proximityPlacementGroupName
  );
  console.log(result);
}

createAProximityPlacementGroup().catch(console.error);
