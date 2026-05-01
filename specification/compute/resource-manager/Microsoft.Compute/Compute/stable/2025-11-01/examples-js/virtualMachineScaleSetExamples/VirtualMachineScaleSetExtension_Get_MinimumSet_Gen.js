const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to get the extension.
 *
 * @summary the operation to get the extension.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_Get_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetExtensionGetMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetExtensions.get(
    "rgcompute",
    "a",
    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
  );
  console.log(result);
}
