const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a snapshot.
 *
 * @summary creates or updates a snapshot.
 * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_Create_ByImportingAnUnmanagedBlobFromADifferentSubscription.json
 */
async function createASnapshotByImportingAnUnmanagedBlobFromADifferentSubscription() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.createOrUpdate("myResourceGroup", "mySnapshot1", {
    location: "West US",
    creationData: {
      createOption: "Import",
      storageAccountId:
        "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount",
      sourceUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
    },
  });
  console.log(result);
}
