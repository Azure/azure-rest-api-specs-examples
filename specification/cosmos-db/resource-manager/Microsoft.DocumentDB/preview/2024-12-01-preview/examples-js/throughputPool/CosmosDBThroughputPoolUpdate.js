const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates the properties of an existing Azure Cosmos DB Throughput Pool.
 *
 * @summary Updates the properties of an existing Azure Cosmos DB Throughput Pool.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/throughputPool/CosmosDBThroughputPoolUpdate.json
 */
async function cosmosDbThroughputPoolUpdate() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const throughputPoolName = "tp1";
  const body = { maxThroughput: 10000 };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.throughputPool.beginUpdateAndWait(
    resourceGroupName,
    throughputPoolName,
    options,
  );
  console.log(result);
}
