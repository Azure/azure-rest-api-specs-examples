const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a gallery inVMAccessControlProfile version.
 *
 * @summary delete a gallery inVMAccessControlProfile version.
 * x-ms-original-file: 2025-03-03/galleryResourceProfileExamples/GalleryInVMAccessControlProfileVersion_Delete.json
 */
async function deleteAGalleryInVMAccessControlProfileVersion() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.galleryInVMAccessControlProfileVersions.delete(
    "myResourceGroup",
    "myGalleryName",
    "myInVMAccessControlProfileName",
    "1.0.0",
  );
}
