const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to restart a virtual machine.
 *
 * @summary The operation to restart a virtual machine.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachines_Restart_MinimumSet_Gen.json
 */
async function virtualMachinesRestartMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmName = "aaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginRestartAndWait(resourceGroupName, vmName);
  console.log(result);
}

virtualMachinesRestartMinimumSetGen().catch(console.error);
