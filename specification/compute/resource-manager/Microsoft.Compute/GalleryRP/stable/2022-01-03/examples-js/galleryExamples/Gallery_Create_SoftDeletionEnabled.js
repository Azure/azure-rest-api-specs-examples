const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Shared Image Gallery.
 *
 * @summary Create or update a Shared Image Gallery.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/Gallery_Create_SoftDeletionEnabled.json
 */
async function createOrUpdateASimpleGalleryWithSoftDeletionEnabled() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const gallery = {
    description: "This is the gallery description.",
    location: "West US",
    softDeletePolicy: { isSoftDeleteEnabled: true },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleries.beginCreateOrUpdateAndWait(
    resourceGroupName,
    galleryName,
    gallery
  );
  console.log(result);
}

createOrUpdateASimpleGalleryWithSoftDeletionEnabled().catch(console.error);
