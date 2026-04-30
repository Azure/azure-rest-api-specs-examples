const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieves information about a Shared Image Gallery.
 *
 * @summary retrieves information about a Shared Image Gallery.
 * x-ms-original-file: 2025-03-03/galleryExamples/Gallery_Get_WithExpandSharingProfileGroups.json
 */
async function getAGalleryWithExpandSharingProfileGroups() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleries.get("myResourceGroup", "myGalleryName", {
    expand: "SharingProfile/Groups",
  });
  console.log(result);
}
