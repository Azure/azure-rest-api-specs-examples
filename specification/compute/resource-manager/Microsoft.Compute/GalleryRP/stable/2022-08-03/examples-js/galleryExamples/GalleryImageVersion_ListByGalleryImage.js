const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List gallery image versions in a gallery image definition.
 *
 * @summary List gallery image versions in a gallery image definition.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-08-03/examples/galleryExamples/GalleryImageVersion_ListByGalleryImage.json
 */
async function listGalleryImageVersionsInAGalleryImageDefinition() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryImageName = "myGalleryImageName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.galleryImageVersions.listByGalleryImage(
    resourceGroupName,
    galleryName,
    galleryImageName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
