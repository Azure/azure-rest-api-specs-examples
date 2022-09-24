const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes virtual machines in a VM scale set.
 *
 * @summary Deletes virtual machines in a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_DeleteInstances_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetsDeleteInstancesMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaaaaaa";
  const forceDeletion = true;
  const vmInstanceIDs = {
    instanceIds: ["aaaaaaaaaaaaaaaaaaaaaaaaa"],
  };
  const options = {
    forceDeletion,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.beginDeleteInstancesAndWait(
    resourceGroupName,
    vmScaleSetName,
    vmInstanceIDs,
    options
  );
  console.log(result);
}

virtualMachineScaleSetsDeleteInstancesMaximumSetGen().catch(console.error);
