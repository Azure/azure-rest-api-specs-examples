const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the CPU, memory, and disk usage statistics for each Cassandra node in a cluster.
 *
 * @summary Gets the CPU, memory, and disk usage statistics for each Cassandra node in a cluster.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBManagedCassandraStatus.json
 */
async function cosmosDbManagedCassandraStatus() {
  const subscriptionId = "subid";
  const resourceGroupName = "cassandra-prod-rg";
  const clusterName = "cassandra-prod";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.cassandraClusters.status(resourceGroupName, clusterName);
  console.log(result);
}

cosmosDbManagedCassandraStatus().catch(console.error);
