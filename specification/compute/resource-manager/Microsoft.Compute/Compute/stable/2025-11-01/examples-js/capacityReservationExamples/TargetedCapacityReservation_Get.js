const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation that retrieves information about the capacity reservation.
 *
 * @summary the operation that retrieves information about the capacity reservation.
 * x-ms-original-file: 2025-11-01/capacityReservationExamples/TargetedCapacityReservation_Get.json
 */
async function getATargetedCapacityReservation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscriptionId}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservations.get(
    "myResourceGroup",
    "targetedCapacityReservationGroup",
    "targetedCapacityReservation",
    { expand: "instanceView" },
  );
  console.log(result);
}
