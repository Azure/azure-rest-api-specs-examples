const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update RUs per second of an Azure Cosmos DB Cassandra Keyspace
 *
 * @summary Update RUs per second of an Azure Cosmos DB Cassandra Keyspace
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBCassandraKeyspaceThroughputUpdate.json
 */
async function cosmosDbCassandraKeyspaceThroughputUpdate() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const keyspaceName = "keyspaceName";
  const updateThroughputParameters = {
    location: "West US",
    resource: { throughput: 400 },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.cassandraResources.beginUpdateCassandraKeyspaceThroughputAndWait(
    resourceGroupName,
    accountName,
    keyspaceName,
    updateThroughputParameters,
  );
  console.log(result);
}
