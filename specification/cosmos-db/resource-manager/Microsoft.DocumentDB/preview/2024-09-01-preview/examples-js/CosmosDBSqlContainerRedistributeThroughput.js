const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Redistribute throughput for an Azure Cosmos DB SQL container
 *
 * @summary Redistribute throughput for an Azure Cosmos DB SQL container
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBSqlContainerRedistributeThroughput.json
 */
async function cosmosDbSqlContainerRedistributeThroughput() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const containerName = "containerName";
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
  const result = await client.sqlResources.beginSqlContainerRedistributeThroughputAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    containerName,
    redistributeThroughputParameters,
  );
  console.log(result);
}
