const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete the extension.
 *
 * @summary The operation to delete the extension.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtensions_Delete_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetExtensionsDeleteMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaa";
  const vmssExtensionName = "aaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetExtensions.beginDeleteAndWait(
    resourceGroupName,
    vmScaleSetName,
    vmssExtensionName
  );
  console.log(result);
}

virtualMachineScaleSetExtensionsDeleteMinimumSetGen().catch(console.error);
