const createComputeManagementClient = require("@azure-rest/arm-compute").default;
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Gets a list of all virtual machine image versions for the specified location, edge zone, publisher, offer, and SKU.
 *
 * @summary Gets a list of all virtual machine image versions for the specified location, edge zone, publisher, offer, and SKU.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_List_MaximumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneListMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const location = "aaaaaaaaaaaaaaa";
  const edgeZone = "aaaaaaaaaaaaaaaaaaaaaaaaa";
  const publisherName = "aaaa";
  const offer = "aaaaaaaaaaaaaaaaaaaaaaaaaa";
  const skus = "aaaaaaaaaaaaaaaaaaaaaaa";
  const options = {
    queryParameters: {
      $expand: "aaaaaaaaaaaaaaaaaaaaaaaa",
      $top: 12,
      $orderby: "aaaaaaaaaaaaaaaaaaaaaaaaaaaa",
      "api-version": "2022-08-01",
    },
  };
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions",
      subscriptionId,
      location,
      edgeZone,
      publisherName,
      offer,
      skus,
    )
    .get(options);
  console.log(result);
}

virtualMachineImagesEdgeZoneListMaximumSetGen().catch(console.error);
