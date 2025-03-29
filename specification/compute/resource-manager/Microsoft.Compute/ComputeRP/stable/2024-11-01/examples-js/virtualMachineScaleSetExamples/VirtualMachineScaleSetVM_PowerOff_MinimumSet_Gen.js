const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Power off (stop) a virtual machine in a VM scale set. Note that resources are still attached and you are getting charged for the resources. Instead, use deallocate to release resources and avoid charges. Additionally, this operation is not allowed on a virtual machine that is being deallocated or already has been deallocated.
 *
 * @summary Power off (stop) a virtual machine in a VM scale set. Note that resources are still attached and you are getting charged for the resources. Instead, use deallocate to release resources and avoid charges. Additionally, this operation is not allowed on a virtual machine that is being deallocated or already has been deallocated.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_PowerOff_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetVMPowerOffMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const instanceId = "aaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.beginPowerOffAndWait(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
  );
  console.log(result);
}
