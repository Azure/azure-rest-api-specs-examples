const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a Shared Image Gallery.
 *
 * @summary Delete a Shared Image Gallery.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/DeleteAGallery.json
 */
async function deleteAGallery() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleries.beginDeleteAndWait(resourceGroupName, galleryName);
  console.log(result);
}

deleteAGallery().catch(console.error);
