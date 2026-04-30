const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates (patches) a snapshot.
 *
 * @summary updates (patches) a snapshot.
 * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_Update.json
 */
async function updateASnapshot() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.update("myResourceGroup", "mySnapshot", {
    diskSizeGB: 20,
    tags: { department: "Development", project: "UpdateSnapshots" },
  });
  console.log(result);
}
