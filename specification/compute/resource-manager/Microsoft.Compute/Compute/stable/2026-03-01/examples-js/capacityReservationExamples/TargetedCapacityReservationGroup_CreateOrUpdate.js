const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a capacity reservation group. When updating a capacity reservation group, only tags and sharing profile may be modified. Please refer to https://aka.ms/CapacityReservation for more details.
 *
 * @summary the operation to create or update a capacity reservation group. When updating a capacity reservation group, only tags and sharing profile may be modified. Please refer to https://aka.ms/CapacityReservation for more details.
 * x-ms-original-file: 2026-03-01/capacityReservationExamples/TargetedCapacityReservationGroup_CreateOrUpdate.json
 */
async function createOrUpdateATargetedCapacityReservationGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservationGroups.createOrUpdate(
    "myResourceGroup",
    "targetedCapacityReservationGroup",
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
      reservationType: "Targeted",
    },
  );
  console.log(result);
}
