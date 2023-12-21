const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete the extension.
 *
 * @summary The operation to delete the extension.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_Delete_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetExtensionDeleteMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const vmScaleSetName = "aaaa";
  const vmssExtensionName = "aaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetExtensions.beginDeleteAndWait(
    resourceGroupName,
    vmScaleSetName,
    vmssExtensionName,
  );
  console.log(result);
}
