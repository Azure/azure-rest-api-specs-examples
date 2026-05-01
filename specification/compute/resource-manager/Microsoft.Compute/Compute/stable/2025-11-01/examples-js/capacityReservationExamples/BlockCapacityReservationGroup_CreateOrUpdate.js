const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a capacity reservation group. When updating a capacity reservation group, only tags and sharing profile may be modified. Please refer to https://aka.ms/CapacityReservation for more details.
 *
 * @summary the operation to create or update a capacity reservation group. When updating a capacity reservation group, only tags and sharing profile may be modified. Please refer to https://aka.ms/CapacityReservation for more details.
 * x-ms-original-file: 2025-11-01/capacityReservationExamples/BlockCapacityReservationGroup_CreateOrUpdate.json
 */
async function createOrUpdateABlockCapacityReservationGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservationGroups.createOrUpdate(
    "myResourceGroup",
    "blockCapacityReservationGroup",
    {
      location: "westus",
      tags: { department: "finance" },
      zones: ["1", "2"],
      reservationType: "Block",
    },
  );
  console.log(result);
}
