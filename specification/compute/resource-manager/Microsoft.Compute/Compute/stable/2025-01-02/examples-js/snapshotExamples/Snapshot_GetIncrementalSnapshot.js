const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets information about a snapshot.
 *
 * @summary gets information about a snapshot.
 * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_GetIncrementalSnapshot.json
 */
async function getInformationAboutAnIncrementalSnapshot() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.get("myResourceGroup", "myIncrementalSnapshot");
  console.log(result);
}
