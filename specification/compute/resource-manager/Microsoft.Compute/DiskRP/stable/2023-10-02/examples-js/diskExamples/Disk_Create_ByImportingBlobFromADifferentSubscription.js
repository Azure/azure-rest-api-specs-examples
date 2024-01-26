const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a disk.
 *
 * @summary Creates or updates a disk.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/diskExamples/Disk_Create_ByImportingBlobFromADifferentSubscription.json
 */
async function createAManagedDiskByImportingAnUnmanagedBlobFromADifferentSubscription() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const diskName = "myDisk";
  const disk = {
    creationData: {
      createOption: "Import",
      sourceUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
      storageAccountId:
        "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount",
    },
    location: "West US",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}
