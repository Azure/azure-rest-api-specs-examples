const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists snapshots under a subscription.
 *
 * @summary Lists snapshots under a subscription.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-01-02/examples/snapshotExamples/Snapshot_ListBySubscription.json
 */
async function listAllSnapshotsInASubscription() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.snapshots.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
