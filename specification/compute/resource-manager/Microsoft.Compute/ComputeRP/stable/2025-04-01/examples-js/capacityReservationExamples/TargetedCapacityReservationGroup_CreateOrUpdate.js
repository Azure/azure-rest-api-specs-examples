const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to The operation to create or update a capacity reservation group. When updating a capacity reservation group, only tags and sharing profile may be modified. Please refer to https://aka.ms/CapacityReservation for more details.
 *
 * @summary The operation to create or update a capacity reservation group. When updating a capacity reservation group, only tags and sharing profile may be modified. Please refer to https://aka.ms/CapacityReservation for more details.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/capacityReservationExamples/TargetedCapacityReservationGroup_CreateOrUpdate.json
 */
async function createOrUpdateATargetedCapacityReservationGroup() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const capacityReservationGroupName = "targetedCapacityReservationGroup";
  const parameters = {
    location: "westus",
    reservationType: "Targeted",
    sharingProfile: {
      subscriptionIds: [
        { id: "/subscriptions/{subscription-id1}" },
        { id: "/subscriptions/{subscription-id2}" },
      ],
    },
    tags: { department: "finance" },
    zones: ["1", "2"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.capacityReservationGroups.createOrUpdate(
    resourceGroupName,
    capacityReservationGroupName,
    parameters,
  );
  console.log(result);
}
