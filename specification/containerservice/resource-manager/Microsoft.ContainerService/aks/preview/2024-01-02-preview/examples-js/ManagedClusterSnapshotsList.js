const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of managed cluster snapshots in the specified subscription.
 *
 * @summary Gets a list of managed cluster snapshots in the specified subscription.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-01-02-preview/examples/ManagedClusterSnapshotsList.json
 */
async function listManagedClusterSnapshots() {
  const subscriptionId =
    process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedClusterSnapshots.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
