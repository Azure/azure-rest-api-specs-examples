const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists snapshots under a subscription.
 *
 * @summary Lists snapshots under a subscription.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/ListSnapshotsInASubscription.json
 */
async function listAllSnapshotsInASubscription() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.snapshots.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllSnapshotsInASubscription().catch(console.error);
