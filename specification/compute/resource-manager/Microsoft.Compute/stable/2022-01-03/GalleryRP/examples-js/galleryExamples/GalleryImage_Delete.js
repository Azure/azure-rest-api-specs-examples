const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a gallery image.
 *
 * @summary Delete a gallery image.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/GalleryImage_Delete.json
 */
async function deleteAGalleryImage() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryImageName = "myGalleryImageName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryImages.beginDeleteAndWait(
    resourceGroupName,
    galleryName,
    galleryImageName
  );
  console.log(result);
}

deleteAGalleryImage().catch(console.error);
