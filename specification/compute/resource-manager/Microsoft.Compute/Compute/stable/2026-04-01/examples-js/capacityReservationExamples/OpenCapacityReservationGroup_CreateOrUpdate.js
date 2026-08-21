const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a capacity reservation group. When updating a capacity reservation group, only tags and sharing profile may be modified. Please refer to https://aka.ms/CapacityReservation for more details.
 *
 * @summary the operation to create or update a capacity reservation group. When updating a capacity reservation group, only tags and sharing profile may be modified. Please refer to https://aka.ms/CapacityReservation for more details.
 * x-ms-original-file: 2026-04-01/capacityReservationExamples/OpenCapacityReservationGroup_CreateOrUpdate.json
 */
async function createOrUpdateAnOpenCapacityReservationGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservationGroups.createOrUpdate(
    "myResourceGroup",
    "openCapacityReservationGroup",
    {
      location: "westus",
      tags: { department: "finance" },
      zones: ["1", "2"],
      sharingProfile: {
        subscriptionIds: [
          { id: "/subscriptions/{subscription-id1}" },
          { id: "/subscriptions/{subscription-id2}" },
        ],
      },
      reservationType: "Open",
    },
  );
  console.log(result);
}
