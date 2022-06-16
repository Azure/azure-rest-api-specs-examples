const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to get the extension.
 *
 * @summary The operation to get the extension.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtensions_Get_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetExtensionsGetMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "a";
  const vmssExtensionName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetExtensions.get(
    resourceGroupName,
    vmScaleSetName,
    vmssExtensionName
  );
  console.log(result);
}

virtualMachineScaleSetExtensionsGetMinimumSetGen().catch(console.error);
