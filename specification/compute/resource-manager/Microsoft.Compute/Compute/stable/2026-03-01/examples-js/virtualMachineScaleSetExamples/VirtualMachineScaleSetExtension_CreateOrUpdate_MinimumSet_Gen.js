const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update an extension.
 *
 * @summary the operation to create or update an extension.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_CreateOrUpdate_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetExtensionCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetExtensions.createOrUpdate(
    "rgcompute",
    "aaaaaaaaaaa",
    "aaaaaaaaaaa",
    {},
  );
  console.log(result);
}
