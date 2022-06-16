const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to get all extensions of an instance in Virtual Machine Scaleset.
 *
 * @summary The operation to get all extensions of an instance in Virtual Machine Scaleset.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtensions_List.json
 */
async function listExtensionsInVmssInstance() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "myvmScaleSet";
  const instanceId = "0";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMExtensions.list(
    resourceGroupName,
    vmScaleSetName,
    instanceId
  );
  console.log(result);
}

listExtensionsInVmssInstance().catch(console.error);
