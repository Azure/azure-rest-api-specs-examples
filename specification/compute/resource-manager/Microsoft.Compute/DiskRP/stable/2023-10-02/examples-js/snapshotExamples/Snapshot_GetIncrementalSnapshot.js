const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about a snapshot.
 *
 * @summary Gets information about a snapshot.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/snapshotExamples/Snapshot_GetIncrementalSnapshot.json
 */
async function getInformationAboutAnIncrementalSnapshot() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const snapshotName = "myIncrementalSnapshot";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.get(resourceGroupName, snapshotName);
  console.log(result);
}
