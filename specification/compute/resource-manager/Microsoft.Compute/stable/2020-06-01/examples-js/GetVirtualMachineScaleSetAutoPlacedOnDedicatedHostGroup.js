const { ComputeManagementClient } = require("@azure/arm-compute-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Display information about a virtual machine scale set.
 *
 * @summary Display information about a virtual machine scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-06-01/examples/GetVirtualMachineScaleSetAutoPlacedOnDedicatedHostGroup.json
 */
async function getAVirtualMachineScaleSetPlacedOnADedicatedHostGroupThroughAutomaticPlacement() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmScaleSetName = "myVirtualMachineScaleSet";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.get(resourceGroupName, vmScaleSetName);
  console.log(result);
}
