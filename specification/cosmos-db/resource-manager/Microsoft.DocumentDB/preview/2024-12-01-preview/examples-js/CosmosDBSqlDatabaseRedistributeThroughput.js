const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Redistribute throughput for an Azure Cosmos DB SQL database
 *
 * @summary Redistribute throughput for an Azure Cosmos DB SQL database
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBSqlDatabaseRedistributeThroughput.json
 */
async function cosmosDbSqlDatabaseRedistributeThroughput() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const redistributeThroughputParameters = {
    resource: {
      sourcePhysicalPartitionThroughputInfo: [{ id: "2", throughput: 5000 }, { id: "3" }],
      targetPhysicalPartitionThroughputInfo: [
        { id: "0", throughput: 5000 },
        { id: "1", throughput: 5000 },
      ],
      throughputPolicy: "custom",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.beginSqlDatabaseRedistributeThroughputAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    redistributeThroughputParameters,
  );
  console.log(result);
}
