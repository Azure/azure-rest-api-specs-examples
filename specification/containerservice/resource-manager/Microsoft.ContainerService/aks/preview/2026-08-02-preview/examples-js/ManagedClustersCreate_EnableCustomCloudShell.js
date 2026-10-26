const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a managed cluster.
 *
 * @summary creates or updates a managed cluster.
 * x-ms-original-file: 2026-08-02-preview/ManagedClustersCreate_EnableCustomCloudShell.json
 */
async function createManagedPrivateClusterWithCustomCloudShellResources() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.createOrUpdate("rg1", "clustername1", {
    location: "location1",
    apiServerAccessProfile: { enablePrivateCluster: true },
    networkProfile: {
      clusterAPIConnectivityProfile: {
        cloudShell: {
          mode: "Custom",
          network: {
            relayNamespaceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cloudshell-rg/providers/Microsoft.Relay/namespaces/cloudshell-relay",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cloudshell-rg/providers/Microsoft.Network/virtualNetworks/cloudshell-vnet/subnets/cloudshell-subnet",
          },
          storage: {
            fileShareName: "cloudshell",
            mode: "Custom",
            storageAccountId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cloudshell-rg/providers/Microsoft.Storage/storageAccounts/cloudshellstorage",
          },
        },
      },
    },
  });
  console.log(result);
}
