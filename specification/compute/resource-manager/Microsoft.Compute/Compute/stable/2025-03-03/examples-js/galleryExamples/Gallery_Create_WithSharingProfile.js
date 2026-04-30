const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Shared Image Gallery.
 *
 * @summary create or update a Shared Image Gallery.
 * x-ms-original-file: 2025-03-03/galleryExamples/Gallery_Create_WithSharingProfile.json
 */
async function createOrUpdateASimpleGalleryWithSharingProfile() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleries.createOrUpdate("myResourceGroup", "myGalleryName", {
    location: "West US",
    description: "This is the gallery description.",
    sharingProfile: { permissions: "Groups" },
  });
  console.log(result);
}
