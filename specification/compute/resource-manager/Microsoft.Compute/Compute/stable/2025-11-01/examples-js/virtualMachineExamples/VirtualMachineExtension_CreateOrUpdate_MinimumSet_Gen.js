const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update the extension.
 *
 * @summary the operation to create or update the extension.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachineExtension_CreateOrUpdate_MinimumSet_Gen.json
 */
async function virtualMachineExtensionCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensions.createOrUpdate(
    "rgcompute",
    "myVM",
    "myVMExtension",
    { location: "westus" },
  );
  console.log(result);
}
