const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update a gallery image version.
 *
 * @summary Update a gallery image version.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/GalleryImageVersion_Update.json
 */
async function updateASimpleGalleryImageVersionManagedImageAsSource() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryImageName = "myGalleryImageName";
  const galleryImageVersionName = "1.0.0";
  const galleryImageVersion = {
    publishingProfile: {
      targetRegions: [
        { name: "West US", regionalReplicaCount: 1 },
        {
          name: "East US",
          regionalReplicaCount: 2,
          storageAccountType: "Standard_ZRS",
        },
      ],
    },
    storageProfile: {
      source: {
        id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryImageVersions.beginUpdateAndWait(
    resourceGroupName,
    galleryName,
    galleryImageName,
    galleryImageVersionName,
    galleryImageVersion,
  );
  console.log(result);
}
