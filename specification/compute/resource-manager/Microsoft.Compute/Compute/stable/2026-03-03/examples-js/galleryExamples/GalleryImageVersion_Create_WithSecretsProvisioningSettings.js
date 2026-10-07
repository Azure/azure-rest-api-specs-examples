const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a gallery image version.
 *
 * @summary create or update a gallery image version.
 * x-ms-original-file: 2026-03-03/galleryExamples/GalleryImageVersion_Create_WithSecretsProvisioningSettings.json
 */
async function createOrUpdateASimpleGalleryImageVersionWithSecretsProvisioningSettings() {
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
      },
      storageProfile: {
        osDiskImage: {
          source: {
            storageAccountId:
              "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}",
            uri: "https://gallerysourcencus.blob.core.windows.net/myvhds/Linux-VM-2024.vhd",
          },
          hostCaching: "ReadOnly",
        },
      },
      securityProfile: {
        secretsProvisioningSettings: {
          isSupported: true,
          osName: "mariner",
          components: [
            { name: "AzureGuestAgent", version: "2.7.0" },
            { name: "SecretsProvisioningLibrary", version: "1.0.0" },
          ],
        },
      },
    },
  );
  console.log(result);
}
