const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a virtual machine extension image.
 *
 * @summary Gets a virtual machine extension image.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImages_Get_MinimumSet_Gen.json
 */
async function virtualMachineExtensionImagesGetMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const location = "aaaaaaaaaaaaaa";
  const publisherName = "aaaaaaaaaaaaaaaaaaaaaaaaaa";
  const typeParam = "aa";
  const version = "aaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensionImages.get(
    location,
    publisherName,
    typeParam,
    version
  );
  console.log(result);
}
