const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a gallery image version.
 *
 * @summary delete a gallery image version.
 * x-ms-original-file: 2026-03-03/galleryExamples/GalleryImageVersion_Delete_BypassSoftDelete.json
 */
async function permanentlyDeleteAGalleryImageVersionByBypassingSoftDelete() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.galleryImageVersions.delete(
    "myResourceGroup",
    "myGalleryName",
    "myGalleryImageName",
    "1.0.0",
    { bypassSoftDelete: true },
  );
}
