const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a gallery image version.
 *
 * @summary create or update a gallery image version.
 * x-ms-original-file: 2025-03-03/galleryExamples/GalleryImageVersion_Create_WithSnapshotsAsSource.json
 */
async function createOrUpdateASimpleGalleryImageVersionUsingSnapshotsAsASource() {
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
                  "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet",
              },
              dataDiskImages: [
                {
                  lun: 1,
                  diskEncryptionSetId:
                    "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet",
                },
              ],
            },
            excludeFromLatest: false,
          },
          {
            name: "East US",
            regionalReplicaCount: 2,
            storageAccountType: "Standard_ZRS",
            encryption: {
              osDiskImage: {
                diskEncryptionSetId:
                  "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet",
              },
              dataDiskImages: [
                {
                  lun: 1,
                  diskEncryptionSetId:
                    "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet",
                },
              ],
            },
            excludeFromLatest: false,
          },
        ],
      },
      storageProfile: {
        osDiskImage: {
          source: {
            id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/snapshots/{osSnapshotName}",
          },
          hostCaching: "ReadOnly",
        },
        dataDiskImages: [
          {
            source: {
              id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/disks/{dataDiskName}",
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
