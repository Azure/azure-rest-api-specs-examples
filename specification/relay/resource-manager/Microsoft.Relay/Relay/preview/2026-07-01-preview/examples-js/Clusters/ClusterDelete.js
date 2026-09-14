const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes a Relay cluster.
 *
 * @summary deletes a Relay cluster.
 * x-ms-original-file: 2026-07-01-preview/Clusters/ClusterDelete.json
 */
async function clusterDelete() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const client = new RelayAPI(credential, subscriptionId);
  await client.clusters.delete("myResourceGroup", "testCluster");
}
