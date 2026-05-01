const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes virtual machines in a VM scale set.
 *
 * @summary deletes virtual machines in a VM scale set.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_DeleteInstances_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetDeleteInstancesMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.virtualMachineScaleSets.deleteInstances("rgcompute", "aaaaaaaaaaaaaaa", {
    instanceIds: ["aaaaaaaaaaaaaaaaaaaaaaaaa"],
  });
}
