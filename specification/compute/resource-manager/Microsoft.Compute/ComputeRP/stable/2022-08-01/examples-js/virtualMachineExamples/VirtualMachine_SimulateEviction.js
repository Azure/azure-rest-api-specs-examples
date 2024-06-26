const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to simulate the eviction of spot virtual machine.
 *
 * @summary The operation to simulate the eviction of spot virtual machine.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachine_SimulateEviction.json
 */
async function simulateEvictionAVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ResourceGroup";
  const vmName = "VMName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.simulateEviction(resourceGroupName, vmName);
  console.log(result);
}

simulateEvictionAVirtualMachine().catch(console.error);
