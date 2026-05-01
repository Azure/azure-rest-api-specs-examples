const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a gallery Script Version.
 *
 * @summary delete a gallery Script Version.
 * x-ms-original-file: 2025-03-03/galleryScriptExamples/GalleryScriptVersion_Delete.json
 */
async function deleteAGalleryScriptVersion() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.galleryScriptVersions.delete(
    "myResourceGroupName",
    "myGalleryName",
    "myGalleryScriptName",
    "1.0.0",
  );
}
