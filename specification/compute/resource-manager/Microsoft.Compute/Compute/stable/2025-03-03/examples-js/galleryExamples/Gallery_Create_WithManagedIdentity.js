const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Shared Image Gallery.
 *
 * @summary create or update a Shared Image Gallery.
 * x-ms-original-file: 2025-03-03/galleryExamples/Gallery_Create_WithManagedIdentity.json
 */
async function createOrUpdateASimpleGalleryWithSystemAssignedAndUserAssignedManagedIdentities() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleries.createOrUpdate("myResourceGroup", "myGalleryName", {
    location: "West US",
    identity: {
      type: "SystemAssigned, UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity":
          {},
      },
    },
    description: "This is the gallery description.",
  });
  console.log(result);
}
