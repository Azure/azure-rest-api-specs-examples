const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Detaches all followers of a database owned by this cluster.
 *
 * @summary Detaches all followers of a database owned by this cluster.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClusterDetachFollowerDatabases.json
 */
async function kustoClusterDetachFollowerDatabases() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const followerDatabaseToRemove = {
    attachedDatabaseConfigurationName: "attachedDatabaseConfigurationsTest",
    clusterResourceId:
      "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/clusters/kustoCluster2",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginDetachFollowerDatabasesAndWait(
    resourceGroupName,
    clusterName,
    followerDatabaseToRemove
  );
  console.log(result);
}

kustoClusterDetachFollowerDatabases().catch(console.error);
