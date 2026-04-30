const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieves information about a gallery Application Version.
 *
 * @summary retrieves information about a gallery Application Version.
 * x-ms-original-file: 2025-03-03/galleryExamples/GalleryApplicationVersion_Get.json
 */
async function getAGalleryApplicationVersion() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryApplicationVersions.get(
    "myResourceGroup",
    "myGalleryName",
    "myGalleryApplicationName",
    "1.0.0",
  );
  console.log(result);
}
