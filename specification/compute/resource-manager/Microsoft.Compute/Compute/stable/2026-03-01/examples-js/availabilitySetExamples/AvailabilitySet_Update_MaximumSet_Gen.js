const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update an availability set.
 *
 * @summary update an availability set.
 * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_Update_MaximumSet_Gen.json
 */
async function availabilitySetUpdateMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.availabilitySets.update("rgcompute", "aaaaaaaaaaaaaaaaaaa", {
    platformFaultDomainCount: 2,
    platformUpdateDomainCount: 20,
    virtualMachines: [
      {
        id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
      },
    ],
    proximityPlacementGroup: {
      id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
    },
    sku: { name: "DSv3-Type1", tier: "aaa", capacity: 7 },
    tags: { key2574: "aaaaaaaa" },
  });
  console.log(result);
}
