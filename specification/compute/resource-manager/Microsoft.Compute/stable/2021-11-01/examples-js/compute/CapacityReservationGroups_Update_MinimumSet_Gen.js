const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update a capacity reservation group. When updating a capacity reservation group, only tags may be modified.
 *
 * @summary The operation to update a capacity reservation group. When updating a capacity reservation group, only tags may be modified.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CapacityReservationGroups_Update_MinimumSet_Gen.json
 */
async function capacityReservationGroupsUpdateMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const capacityReservationGroupName = "aaaaaaaaaaaaaaaaaaaaaa";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservationGroups.update(
    resourceGroupName,
    capacityReservationGroupName,
    parameters
  );
  console.log(result);
}

capacityReservationGroupsUpdateMinimumSetGen().catch(console.error);
