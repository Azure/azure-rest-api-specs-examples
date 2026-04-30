const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a gallery image version.
 *
 * @summary create or update a gallery image version.
 * x-ms-original-file: 2025-03-03/galleryExamples/GalleryImageVersion_Create_WithStorageAccountStrategy.json
 */
async function createOrUpdateASimpleGalleryImageVersionWithStorageAccountStrategyAndRegionalStorageAccountTypeOverride() {
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
          { name: "West US" },
          { name: "East US" },
          { name: "East US 2", storageAccountType: "Premium_LRS" },
        ],
        storageAccountStrategy: "PreferStandard_ZRS",
      },
      storageProfile: {
        source: {
          virtualMachineId:
            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/virtualMachines/{vmName}",
        },
      },
    },
  );
  console.log(result);
}
