const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a disk.
 *
 * @summary creates or updates a disk.
 * x-ms-original-file: 2025-01-02/diskExamples/Disk_Create_FromImportSecure_WithVMMetadata.json
 */
async function createAManagedDiskFromImportSecureCreateOptionWithMetadataURIForConfidentialVM() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.createOrUpdate("myResourceGroup", "myDisk", {
    location: "West US",
    osType: "Windows",
    securityProfile: { securityType: "ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey" },
    creationData: {
      createOption: "ImportSecure",
      storageAccountId:
        "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount",
      sourceUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
      securityDataUri: "https://mystorageaccount.blob.core.windows.net/osimages/vmgs.vhd",
      securityMetadataUri: "https://mystorageaccount.blob.core.windows.net/osimages/vmmd.vhd",
    },
  });
  console.log(result);
}
