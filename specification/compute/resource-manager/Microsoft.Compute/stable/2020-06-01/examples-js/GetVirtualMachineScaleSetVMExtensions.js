const { ComputeManagementClient } = require("@azure/arm-compute-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to get the VMSS VM extension.
 *
 * @summary The operation to get the VMSS VM extension.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-06-01/examples/GetVirtualMachineScaleSetVMExtensions.json
 */
async function getVirtualMachineScaleSetVMExtension() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmScaleSetName = "myvmScaleSet";
  const instanceId = "0";
  const vmExtensionName = "myVMExtension";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMExtensions.get(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
    vmExtensionName
  );
  console.log(result);
}
