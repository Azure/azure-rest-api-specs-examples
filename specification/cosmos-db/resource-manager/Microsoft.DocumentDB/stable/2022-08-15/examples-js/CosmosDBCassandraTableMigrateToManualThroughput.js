const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Migrate an Azure Cosmos DB Cassandra table from autoscale to manual throughput
 *
 * @summary Migrate an Azure Cosmos DB Cassandra table from autoscale to manual throughput
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBCassandraTableMigrateToManualThroughput.json
 */
async function cosmosDbCassandraTableMigrateToManualThroughput() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const keyspaceName = "keyspaceName";
  const tableName = "tableName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result =
    await client.cassandraResources.beginMigrateCassandraTableToManualThroughputAndWait(
      resourceGroupName,
      accountName,
      keyspaceName,
      tableName
    );
  console.log(result);
}

cosmosDbCassandraTableMigrateToManualThroughput().catch(console.error);
