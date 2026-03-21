const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes a managed cluster.
 *
 * @summary deletes a managed cluster.
 * x-ms-original-file: 2026-01-01/ManagedClustersDelete.json
 */
async function deleteManagedCluster() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  await client.managedClusters.delete("rg1", "clustername1");
}
