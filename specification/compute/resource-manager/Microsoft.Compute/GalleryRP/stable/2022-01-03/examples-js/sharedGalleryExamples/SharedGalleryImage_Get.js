const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a shared gallery image by subscription id or tenant id.
 *
 * @summary Get a shared gallery image by subscription id or tenant id.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/sharedGalleryExamples/SharedGalleryImage_Get.json
 */
async function getASharedGalleryImage() {
  const subscriptionId = "{subscription-id}";
  const location = "myLocation";
  const galleryUniqueName = "galleryUniqueName";
  const galleryImageName = "myGalleryImageName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sharedGalleryImages.get(
    location,
    galleryUniqueName,
    galleryImageName
  );
  console.log(result);
}

getASharedGalleryImage().catch(console.error);
