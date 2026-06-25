const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to update an extension.
 *
 * @summary the operation to update an extension.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_Update_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetExtensionUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetExtensions.update(
    "rgcompute",
    "aaaaaaaaaaaaaaaaaaaaaaaaaa",
    "aa",
    {},
  );
  console.log(result);
}
