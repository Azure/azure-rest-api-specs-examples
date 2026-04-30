const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to update a capacity reservation group. When updating a capacity reservation group, only tags and sharing profile may be modified.
 *
 * @summary the operation to update a capacity reservation group. When updating a capacity reservation group, only tags and sharing profile may be modified.
 * x-ms-original-file: 2025-11-01/capacityReservationExamples/CapacityReservationGroup_Update_MaximumSet_Gen.json
 */
async function capacityReservationGroupUpdateMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservationGroups.update(
    "rgcompute",
    "aaaaaaaaaaaaaaaaaaaaaa",
    { tags: { key5355: "aaa" } },
  );
  console.log(result);
}
