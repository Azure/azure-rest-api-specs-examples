const createComputeManagementClient = require("@azure-rest/arm-compute").default;
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Gets a list of virtual machine extension image types.
 *
 * @summary Gets a list of virtual machine extension image types.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImages_ListTypes_MinimumSet_Gen.json
 */
async function virtualMachineExtensionImagesListTypesMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const location = "aaaa";
  const publisherName = "aa";
  const options = {
    queryParameters: { "api-version": "2022-08-01" },
  };
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types",
      subscriptionId,
      location,
      publisherName
    )
    .get(options);
  console.log(result);
}

virtualMachineExtensionImagesListTypesMinimumSetGen().catch(console.error);
