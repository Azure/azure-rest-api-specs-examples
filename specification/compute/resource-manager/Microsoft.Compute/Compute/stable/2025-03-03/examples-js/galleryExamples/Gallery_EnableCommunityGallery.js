const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update sharing profile of a gallery.
 *
 * @summary update sharing profile of a gallery.
 * x-ms-original-file: 2025-03-03/galleryExamples/Gallery_EnableCommunityGallery.json
 */
async function shareAGalleryToCommunity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.gallerySharingProfile.update("myResourceGroup", "myGalleryName", {
    operationType: "EnableCommunity",
  });
  console.log(result);
}
