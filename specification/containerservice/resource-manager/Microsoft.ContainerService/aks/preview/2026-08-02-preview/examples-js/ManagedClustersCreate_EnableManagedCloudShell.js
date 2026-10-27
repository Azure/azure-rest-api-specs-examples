const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a managed cluster.
 *
 * @summary creates or updates a managed cluster.
 * x-ms-original-file: 2026-08-02-preview/ManagedClustersCreate_EnableManagedCloudShell.json
 */
async function createManagedPrivateClusterWithAKSManagedCloudShellResources() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.createOrUpdate("rg1", "clustername1", {
    location: "location1",
    apiServerAccessProfile: { enablePrivateCluster: true },
    networkProfile: {
      clusterAPIConnectivityProfile: {
        cloudShell: { mode: "Managed", storage: { mode: "Managed" } },
      },
    },
  });
  console.log(result);
}
