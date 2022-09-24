const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update a capacity reservation.
 *
 * @summary The operation to update a capacity reservation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/capacityReservationExamples/CapacityReservation_Update_MinimumSet_Gen.json
 */
async function capacityReservationsUpdateMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const capacityReservationGroupName = "aaaaaaaaaaaaaaaaaaaaaaaaaa";
  const capacityReservationName = "aaa";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservations.beginUpdateAndWait(
    resourceGroupName,
    capacityReservationGroupName,
    capacityReservationName,
    parameters
  );
  console.log(result);
}

capacityReservationsUpdateMinimumSetGen().catch(console.error);
