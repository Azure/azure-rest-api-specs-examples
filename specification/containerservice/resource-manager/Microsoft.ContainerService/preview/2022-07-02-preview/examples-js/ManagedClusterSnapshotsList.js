const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of managed cluster snapshots in the specified subscription.
 *
 * @summary Gets a list of managed cluster snapshots in the specified subscription.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-07-02-preview/examples/ManagedClusterSnapshotsList.json
 */
async function listManagedClusterSnapshots() {
  const subscriptionId = "subid1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedClusterSnapshots.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedClusterSnapshots().catch(console.error);
