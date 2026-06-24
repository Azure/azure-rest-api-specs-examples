const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation that retrieves information about the capacity reservation.
 *
 * @summary the operation that retrieves information about the capacity reservation.
 * x-ms-original-file: 2026-03-01/capacityReservationExamples/BlockCapacityReservation_Get.json
 */
async function getABlockCapacityReservation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscriptionId}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservations.get(
    "myResourceGroup",
    "blockCapacityReservationGroup",
    "blockCapacityReservation",
    { expand: "instanceView" },
  );
  console.log(result);
}
