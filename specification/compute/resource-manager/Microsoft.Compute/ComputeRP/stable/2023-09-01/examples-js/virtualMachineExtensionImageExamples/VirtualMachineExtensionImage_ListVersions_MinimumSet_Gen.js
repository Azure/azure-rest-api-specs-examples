const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of virtual machine extension image versions.
 *
 * @summary Gets a list of virtual machine extension image versions.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_ListVersions_MinimumSet_Gen.json
 */
async function virtualMachineExtensionImageListVersionsMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const location = "aaaaaaaaa";
  const publisherName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const typeParam = "aaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensionImages.listVersions(
    location,
    publisherName,
    typeParam,
  );
  console.log(result);
}
