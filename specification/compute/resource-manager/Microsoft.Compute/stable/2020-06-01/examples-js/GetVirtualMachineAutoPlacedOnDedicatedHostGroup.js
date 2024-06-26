const { ComputeManagementClient } = require("@azure/arm-compute-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves information about the model view or the instance view of a virtual machine.
 *
 * @summary Retrieves information about the model view or the instance view of a virtual machine.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-06-01/examples/GetVirtualMachineAutoPlacedOnDedicatedHostGroup.json
 */
async function getAVirtualMachinePlacedOnADedicatedHostGroupThroughAutomaticPlacement() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmName = "myVM";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.get(resourceGroupName, vmName);
  console.log(result);
}
