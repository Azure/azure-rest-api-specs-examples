const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the properties of a managed Cassandra data center.
 *
 * @summary Get the properties of a managed Cassandra data center.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBManagedCassandraDataCenterGet.json
 */
async function cosmosDbManagedCassandraDataCenterGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "cassandra-prod-rg";
  const clusterName = "cassandra-prod";
  const dataCenterName = "dc1";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.cassandraDataCenters.get(
    resourceGroupName,
    clusterName,
    dataCenterName
  );
  console.log(result);
}

cosmosDbManagedCassandraDataCenterGet().catch(console.error);
