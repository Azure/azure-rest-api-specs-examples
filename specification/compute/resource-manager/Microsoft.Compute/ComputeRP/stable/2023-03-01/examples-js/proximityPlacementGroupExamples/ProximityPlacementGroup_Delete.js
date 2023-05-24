const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a proximity placement group.
 *
 * @summary Delete a proximity placement group.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/proximityPlacementGroupExamples/ProximityPlacementGroup_Delete.json
 */
async function deleteAProximityPlacementGroup() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const proximityPlacementGroupName = "myProximityPlacementGroup";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.proximityPlacementGroups.delete(
    resourceGroupName,
    proximityPlacementGroupName
  );
  console.log(result);
}
