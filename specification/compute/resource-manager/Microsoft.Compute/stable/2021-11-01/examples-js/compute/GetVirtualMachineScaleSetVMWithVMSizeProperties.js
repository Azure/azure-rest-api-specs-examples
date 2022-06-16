const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a virtual machine from a VM scale set.
 *
 * @summary Gets a virtual machine from a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetVirtualMachineScaleSetVMWithVMSizeProperties.json
 */
async function getVMScaleSetVMWithVMSizeProperties() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "{vmss-name}";
  const instanceId = "0";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.get(
    resourceGroupName,
    vmScaleSetName,
    instanceId
  );
  console.log(result);
}

getVMScaleSetVMWithVMSizeProperties().catch(console.error);
