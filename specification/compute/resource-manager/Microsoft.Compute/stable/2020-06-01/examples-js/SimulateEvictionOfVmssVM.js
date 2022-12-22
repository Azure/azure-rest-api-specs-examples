const { ComputeManagementClient } = require("@azure/arm-compute-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to simulate the eviction of spot virtual machine in a VM scale set.
 *
 * @summary The operation to simulate the eviction of spot virtual machine in a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-06-01/examples/SimulateEvictionOfVmssVM.json
 */
async function simulateEvictionAVirtualMachine() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "ResourceGroup";
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
