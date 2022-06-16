const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a snapshot.
 *
 * @summary Creates or updates a snapshot.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/CreateASnapshotByImportingAnUnmanagedBlobFromTheSameSubscription.json
 */
async function createASnapshotByImportingAnUnmanagedBlobFromTheSameSubscription() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const snapshotName = "mySnapshot1";
  const snapshot = {
    creationData: {
      createOption: "Import",
      sourceUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
    },
    location: "West US",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.beginCreateOrUpdateAndWait(
    resourceGroupName,
    snapshotName,
    snapshot
  );
  console.log(result);
}

createASnapshotByImportingAnUnmanagedBlobFromTheSameSubscription().catch(console.error);
