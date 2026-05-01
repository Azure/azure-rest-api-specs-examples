const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a capacity reservation. Please note some properties can be set only during capacity reservation creation. Please refer to https://aka.ms/CapacityReservation for more details.
 *
 * @summary the operation to create or update a capacity reservation. Please note some properties can be set only during capacity reservation creation. Please refer to https://aka.ms/CapacityReservation for more details.
 * x-ms-original-file: 2025-11-01/capacityReservationExamples/BlockCapacityReservation_CreateOrUpdate.json
 */
async function createOrUpdateABlockCapacityReservation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservations.createOrUpdate(
    "myResourceGroup",
    "blockCapacityReservationGroup",
    "blockCapacityReservation",
    {
      location: "westus",
      tags: { department: "HR" },
      sku: { name: "Standard_ND96isr_H100_v5", capacity: 1 },
      scheduleProfile: { start: "2025-08-01", end: "2025-08-02" },
      zones: ["1"],
    },
  );
  console.log(result);
}
