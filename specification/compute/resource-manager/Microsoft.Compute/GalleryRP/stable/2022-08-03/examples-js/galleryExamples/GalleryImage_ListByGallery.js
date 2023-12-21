const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List gallery image definitions in a gallery.
 *
 * @summary List gallery image definitions in a gallery.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-08-03/examples/galleryExamples/GalleryImage_ListByGallery.json
 */
async function listGalleryImagesInAGallery() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const galleryName = "myGalleryName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.galleryImages.listByGallery(resourceGroupName, galleryName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
