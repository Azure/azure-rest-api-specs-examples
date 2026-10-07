const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a gallery image version.
 *
 * @summary create or update a gallery image version.
 * x-ms-original-file: 2026-03-03/galleryExamples/GalleryImageVersion_Create_WithCVMDataDiskEncryption.json
 */
async function createOrUpdateAGalleryImageVersionWithCVMDataDiskEncryptionUsingCustomerManagedKey() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryImageVersions.createOrUpdate(
    "myResourceGroup",
    "myGalleryName",
    "myGalleryImageName",
    "1.0.0",
    {
      location: "eastus",
      publishingProfile: {
        targetRegions: [
          {
            name: "eastus",
            regionalReplicaCount: 1,
            storageAccountType: "Standard_ZRS",
            encryption: {
              osDiskImage: {
                securityProfile: { confidentialVMEncryptionType: "EncryptedWithPmk" },
              },
              dataDiskImages: [
                {
                  lun: 0,
                  securityProfile: {
                    confidentialVMEncryptionType: "DataDiskEncryptedWithCmk",
                    secureVMDiskEncryptionSetId:
                      "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet",
                  },
                },
              ],
            },
            excludeFromLatest: false,
          },
        ],
        replicaCount: 1,
        excludeFromLatest: false,
        replicationMode: "Full",
      },
      storageProfile: {
        source: {
          virtualMachineId:
            "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM",
        },
      },
    },
  );
  console.log(result);
}
