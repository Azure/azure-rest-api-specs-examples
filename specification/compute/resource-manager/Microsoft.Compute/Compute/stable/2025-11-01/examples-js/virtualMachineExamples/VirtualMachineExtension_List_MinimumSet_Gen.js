const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to get all extensions of a Virtual Machine.
 *
 * @summary the operation to get all extensions of a Virtual Machine.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachineExtension_List_MinimumSet_Gen.json
 */
async function virtualMachineExtensionListMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensions.list(
    "rgcompute",
    "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
  );
  console.log(result);
}
