const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to simulate the eviction of spot virtual machine in a VM scale set.
 *
 * @summary The operation to simulate the eviction of spot virtual machine in a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/SimulateEvictionOfVmssVM.json
 */
async function simulateEvictionAVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ResourceGroup";
  const vmScaleSetName = "VmScaleSetName";
  const instanceId = "InstanceId";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.simulateEviction(
    resourceGroupName,
    vmScaleSetName,
    instanceId
  );
  console.log(result);
}

simulateEvictionAVirtualMachine().catch(console.error);
