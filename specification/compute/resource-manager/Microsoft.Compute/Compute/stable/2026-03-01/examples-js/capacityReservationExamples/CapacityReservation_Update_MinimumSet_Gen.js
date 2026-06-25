const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to update a capacity reservation.
 *
 * @summary the operation to update a capacity reservation.
 * x-ms-original-file: 2026-03-01/capacityReservationExamples/CapacityReservation_Update_MinimumSet_Gen.json
 */
async function capacityReservationUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservations.update(
    "rgcompute",
    "aaaaaaaaaaaaaaaaaaaaaaaaaa",
    "aaa",
    {},
  );
  console.log(result);
}
