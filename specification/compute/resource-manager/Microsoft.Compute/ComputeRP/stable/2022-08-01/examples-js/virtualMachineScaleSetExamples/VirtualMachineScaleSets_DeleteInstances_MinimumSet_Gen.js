const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes virtual machines in a VM scale set.
 *
 * @summary Deletes virtual machines in a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_DeleteInstances_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetsDeleteInstancesMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaa";
  const vmInstanceIDs = {
    instanceIds: ["aaaaaaaaaaaaaaaaaaaaaaaaa"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.beginDeleteInstancesAndWait(
    resourceGroupName,
    vmScaleSetName,
    vmInstanceIDs
  );
  console.log(result);
}

virtualMachineScaleSetsDeleteInstancesMinimumSetGen().catch(console.error);
