const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

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
