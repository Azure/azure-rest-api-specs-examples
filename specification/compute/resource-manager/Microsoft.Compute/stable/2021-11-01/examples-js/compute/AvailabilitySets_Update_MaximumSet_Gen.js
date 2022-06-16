const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function availabilitySetsUpdateMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const availabilitySetName = "aaaaaaaaaaaaaaaaaaa";
  const parameters = {
    platformFaultDomainCount: 2,
    platformUpdateDomainCount: 20,
    proximityPlacementGroup: {
      id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
    },
    sku: { name: "DSv3-Type1", capacity: 7, tier: "aaa" },
    tags: { key2574: "aaaaaaaa" },
    virtualMachines: [
      {
        id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.availabilitySets.update(
    resourceGroupName,
    availabilitySetName,
    parameters
  );
  console.log(result);
}

availabilitySetsUpdateMaximumSetGen().catch(console.error);
