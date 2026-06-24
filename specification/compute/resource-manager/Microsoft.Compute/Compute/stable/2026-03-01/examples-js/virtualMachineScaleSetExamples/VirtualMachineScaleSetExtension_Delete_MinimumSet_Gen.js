const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to delete the extension.
 *
 * @summary the operation to delete the extension.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_Delete_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetExtensionDeleteMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.virtualMachineScaleSetExtensions.delete(
    "rgcompute",
    "aaaa",
    "aaaaaaaaaaaaaaaaaaaaaaa",
  );
}
