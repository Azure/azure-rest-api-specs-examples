const createComputeManagementClient = require("@azure-rest/arm-compute").default;
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
 *
 * @summary Gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineImageExamples/VirtualMachineImages_List_MaximumSet_Gen.json
 */
async function virtualMachineImagesListMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const location = "aaaaaaaaaaaaaaa";
  const publisherName = "aaaaaa";
  const offer = "aaaaaaaaaaaaaaaa";
  const skus = "aaaaaaaaaaaaaaaaaaaaaaa";
  const options = {
    queryParameters: {
      $expand: "aaaaaaaaaaaaaaaaaaaaaaaa",
      $top: 18,
      $orderby: "aa",
      "api-version": "2022-08-01",
    },
  };
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions",
      subscriptionId,
      location,
      publisherName,
      offer,
      skus
    )
    .get(options);
  console.log(result);
}

virtualMachineImagesListMaximumSetGen().catch(console.error);
