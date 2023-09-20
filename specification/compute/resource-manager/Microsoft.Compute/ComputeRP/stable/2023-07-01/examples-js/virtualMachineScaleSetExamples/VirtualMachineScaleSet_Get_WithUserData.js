const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Display information about a virtual machine scale set.
 *
 * @summary Display information about a virtual machine scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Get_WithUserData.json
 */
async function getAVirtualMachineScaleSetWithUserData() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmScaleSetName = "myVirtualMachineScaleSet";
  const expand = "userData";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.get(
    resourceGroupName,
    vmScaleSetName,
    options
  );
  console.log(result);
}
