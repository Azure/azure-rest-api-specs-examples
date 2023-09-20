const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Reimages (upgrade the operating system) a specific virtual machine in a VM scale set.
 *
 * @summary Reimages (upgrade the operating system) a specific virtual machine in a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Reimage_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetVMReimageMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaa";
  const instanceId = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const vmScaleSetVMReimageInput = {
    tempDisk: true,
  };
  const options = {
    vmScaleSetVMReimageInput,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.beginReimageAndWait(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
    options
  );
  console.log(result);
}
