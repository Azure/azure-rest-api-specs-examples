const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of virtual machine extension image types.
 *
 * @summary Gets a list of virtual machine extension image types.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImages_ListTypes_MaximumSet_Gen.json
 */
async function virtualMachineExtensionImagesListTypesMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaaaaaaaaaaaaaaaaaaa";
  const publisherName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensionImages.listTypes(location, publisherName);
  console.log(result);
}

virtualMachineExtensionImagesListTypesMaximumSetGen().catch(console.error);
