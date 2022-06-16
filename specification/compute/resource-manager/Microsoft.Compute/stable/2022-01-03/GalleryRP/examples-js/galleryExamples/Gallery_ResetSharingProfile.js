const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update sharing profile of a gallery.
 *
 * @summary Update sharing profile of a gallery.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/Gallery_ResetSharingProfile.json
 */
async function resetSharingProfileOfAGallery() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const sharingUpdate = { operationType: "Reset" };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.gallerySharingProfile.beginUpdateAndWait(
    resourceGroupName,
    galleryName,
    sharingUpdate
  );
  console.log(result);
}

resetSharingProfileOfAGallery().catch(console.error);
