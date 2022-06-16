const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of virtual machine extension image versions.
 *
 * @summary Gets a list of virtual machine extension image versions.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImages_ListVersions_MaximumSet_Gen.json
 */
async function virtualMachineExtensionImagesListVersionsMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaaaaaaaaaaaaaaaaaaa";
  const publisherName = "aaaaaaaaaaaaaaaaaaaa";
  const typeParam = "aaaaaaaaaaaaaaaaaa";
  const filter = "aaaaaaaaaaaaaaaaaaaaaaaaa";
  const top = 22;
  const orderby = "a";
  const options = {
    filter,
    top,
    orderby,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensionImages.listVersions(
    location,
    publisherName,
    typeParam,
    options
  );
  console.log(result);
}

virtualMachineExtensionImagesListVersionsMaximumSetGen().catch(console.error);
