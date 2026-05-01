const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a gallery image version.
 *
 * @summary create or update a gallery image version.
 * x-ms-original-file: 2025-03-03/galleryExamples/GalleryImageVersion_Create_WithVHD.json
 */
async function createOrUpdateASimpleGalleryImageVersionUsingVhdAsASource() {
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
        targetRegions: [
          {
            name: "West US",
            regionalReplicaCount: 1,
            encryption: {
              osDiskImage: {
                diskEncryptionSetId:
                  "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet",
              },
              dataDiskImages: [
                {
                  diskEncryptionSetId:
                    "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet",
                  lun: 1,
                },
              ],
            },
            excludeFromLatest: false,
          },
          {
            name: "East US",
            regionalReplicaCount: 2,
            storageAccountType: "Standard_ZRS",
            excludeFromLatest: false,
          },
        ],
      },
      storageProfile: {
        osDiskImage: {
          source: {
            storageAccountId:
              "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}",
            uri: "https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd",
          },
          hostCaching: "ReadOnly",
        },
        dataDiskImages: [
          {
            source: {
              storageAccountId:
                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}",
              uri: "https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd",
            },
            lun: 1,
            hostCaching: "None",
          },
        ],
      },
      safetyProfile: {
        allowDeletionOfReplicatedLocations: false,
        blockDeletionBeforeEndOfLife: false,
      },
    },
  );
  console.log(result);
}
