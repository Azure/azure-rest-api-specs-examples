const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restarts a virtual machine in a VM scale set.
 *
 * @summary Restarts a virtual machine in a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Restart_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetVMRestartMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaa";
  const instanceId = "aaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.beginRestartAndWait(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
  );
  console.log(result);
}
