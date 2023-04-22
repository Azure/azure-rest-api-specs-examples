const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Migrate an Azure Cosmos DB Cassandra Keyspace from manual throughput to autoscale
 *
 * @summary Migrate an Azure Cosmos DB Cassandra Keyspace from manual throughput to autoscale
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/CosmosDBCassandraKeyspaceMigrateToAutoscale.json
 */
async function cosmosDbCassandraKeyspaceMigrateToAutoscale() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const keyspaceName = "keyspaceName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.cassandraResources.beginMigrateCassandraKeyspaceToAutoscaleAndWait(
    resourceGroupName,
    accountName,
    keyspaceName
  );
  console.log(result);
}
