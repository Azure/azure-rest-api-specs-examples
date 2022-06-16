const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts a virtual machine in a VM scale set.
 *
 * @summary Starts a virtual machine in a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMs_Start_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetVMSStartMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaa";
  const instanceId = "aaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.beginStartAndWait(
    resourceGroupName,
    vmScaleSetName,
    instanceId
  );
  console.log(result);
}

virtualMachineScaleSetVMSStartMaximumSetGen().catch(console.error);
