const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update RUs per second of an Azure Cosmos DB Cassandra table
 *
 * @summary Update RUs per second of an Azure Cosmos DB Cassandra table
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBCassandraTableThroughputUpdate.json
 */
async function cosmosDbCassandraTableThroughputUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const keyspaceName = "keyspaceName";
  const tableName = "tableName";
  const updateThroughputParameters = {
    location: "West US",
    resource: { throughput: 400 },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.cassandraResources.beginUpdateCassandraTableThroughputAndWait(
    resourceGroupName,
    accountName,
    keyspaceName,
    tableName,
    updateThroughputParameters
  );
  console.log(result);
}

cosmosDbCassandraTableThroughputUpdate().catch(console.error);
