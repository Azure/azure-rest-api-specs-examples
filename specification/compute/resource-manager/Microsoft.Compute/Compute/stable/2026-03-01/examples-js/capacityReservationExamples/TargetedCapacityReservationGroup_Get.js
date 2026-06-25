const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation that retrieves information about a capacity reservation group.
 *
 * @summary the operation that retrieves information about a capacity reservation group.
 * x-ms-original-file: 2026-03-01/capacityReservationExamples/TargetedCapacityReservationGroup_Get.json
 */
async function getATargetedCapacityReservationGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscriptionId}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservationGroups.get(
    "myResourceGroup",
    "targetedCapacityReservationGroup",
    { expand: "instanceView" },
  );
  console.log(result);
}
