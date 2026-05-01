const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a snapshot.
 *
 * @summary creates or updates a snapshot.
 * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_Create_InstantAccessSnapshot.json
 */
async function createASnapshotWhichCanBeInstantlyAccessable() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.createOrUpdate("myResourceGroup", "mySnapshot2", {
    location: "West US",
    creationData: {
      createOption: "Copy",
      sourceResourceId:
        "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk1",
      instantAccessDurationMinutes: 120,
    },
  });
  console.log(result);
}
