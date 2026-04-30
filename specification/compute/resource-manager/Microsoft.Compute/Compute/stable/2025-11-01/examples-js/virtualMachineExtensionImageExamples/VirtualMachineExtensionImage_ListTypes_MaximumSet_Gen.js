const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of virtual machine extension image types.
 *
 * @summary gets a list of virtual machine extension image types.
 * x-ms-original-file: 2025-11-01/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_ListTypes_MaximumSet_Gen.json
 */
async function virtualMachineExtensionImageListTypesMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensionImages.listTypes(
    "aaaaaaaaaaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
  );
  console.log(result);
}
