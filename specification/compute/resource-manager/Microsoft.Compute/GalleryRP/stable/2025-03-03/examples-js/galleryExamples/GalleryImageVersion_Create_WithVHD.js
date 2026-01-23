const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a gallery image version.
 *
 * @summary Create or update a gallery image version.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/GalleryImageVersion_Create_WithVHD.json
 */
async function createOrUpdateASimpleGalleryImageVersionUsingVhdAsASource() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryImageName = "myGalleryImageName";
  const galleryImageVersionName = "1.0.0";
  const galleryImageVersion = {
    location: "West US",
    publishingProfile: {
      targetRegions: [
        {
          name: "West US",
          encryption: {
            dataDiskImages: [
              {
                diskEncryptionSetId:
                  "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet",
                lun: 1,
              },
            ],
            osDiskImage: {
              diskEncryptionSetId:
                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet",
            },
          },
          excludeFromLatest: false,
          regionalReplicaCount: 1,
        },
        {
          name: "East US",
          excludeFromLatest: false,
          regionalReplicaCount: 2,
          storageAccountType: "Standard_ZRS",
        },
      ],
    },
    safetyProfile: {
      allowDeletionOfReplicatedLocations: false,
      blockDeletionBeforeEndOfLife: false,
    },
    storageProfile: {
      dataDiskImages: [
        {
          hostCaching: "None",
          lun: 1,
          source: {
            storageAccountId:
              "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}",
            uri: "https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd",
          },
        },
      ],
      osDiskImage: {
        hostCaching: "ReadOnly",
        source: {
          storageAccountId:
            "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}",
          uri: "https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd",
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryImageVersions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    galleryName,
    galleryImageName,
    galleryImageVersionName,
    galleryImageVersion,
  );
  console.log(result);
}
