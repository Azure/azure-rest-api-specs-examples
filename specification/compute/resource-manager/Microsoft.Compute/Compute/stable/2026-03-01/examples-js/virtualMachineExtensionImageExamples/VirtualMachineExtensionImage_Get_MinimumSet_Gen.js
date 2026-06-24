const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a virtual machine extension image.
 *
 * @summary gets a virtual machine extension image.
 * x-ms-original-file: 2026-03-01/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_Get_MinimumSet_Gen.json
 */
async function virtualMachineExtensionImageGetMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensionImages.get(
    "aaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaaaaa",
    "aa",
    "aaa",
  );
  console.log(result);
}
