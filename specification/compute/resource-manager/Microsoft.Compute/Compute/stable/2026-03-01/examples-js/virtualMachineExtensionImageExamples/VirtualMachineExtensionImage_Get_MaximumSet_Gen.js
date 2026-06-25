const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a virtual machine extension image.
 *
 * @summary gets a virtual machine extension image.
 * x-ms-original-file: 2026-03-01/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_Get_MaximumSet_Gen.json
 */
async function virtualMachineExtensionImageGetMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensionImages.get(
    "aaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaa",
  );
  console.log(result);
}
