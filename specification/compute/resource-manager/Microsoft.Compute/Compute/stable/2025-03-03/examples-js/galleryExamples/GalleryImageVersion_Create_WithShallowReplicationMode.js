const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a gallery image version.
 *
 * @summary create or update a gallery image version.
 * x-ms-original-file: 2025-03-03/galleryExamples/GalleryImageVersion_Create_WithShallowReplicationMode.json
 */
async function createOrUpdateASimpleGalleryImageVersionUsingShallowReplicationMode() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryImageVersions.createOrUpdate(
    "myResourceGroup",
    "myGalleryName",
    "myGalleryImageName",
    "1.0.0",
    {
      location: "West US",
      publishingProfile: {
        targetRegions: [{ name: "West US", regionalReplicaCount: 1, excludeFromLatest: false }],
        replicationMode: "Shallow",
      },
      storageProfile: {
        source: {
          id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}",
        },
      },
      safetyProfile: {
        allowDeletionOfReplicatedLocations: false,
        blockDeletionBeforeEndOfLife: false,
      },
    },
  );
  console.log(result);
}
