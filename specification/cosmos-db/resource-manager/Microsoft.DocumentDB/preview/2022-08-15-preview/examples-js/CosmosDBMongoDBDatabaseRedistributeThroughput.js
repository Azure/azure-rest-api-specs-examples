const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Redistribute throughput for an Azure Cosmos DB MongoDB database
 *
 * @summary Redistribute throughput for an Azure Cosmos DB MongoDB database
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBMongoDBDatabaseRedistributeThroughput.json
 */
async function cosmosDbMongoDbdatabaseRedistributeThroughput() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
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
  const result = await client.mongoDBResources.beginMongoDBDatabaseRedistributeThroughputAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    redistributeThroughputParameters
  );
  console.log(result);
}

cosmosDbMongoDbdatabaseRedistributeThroughput().catch(console.error);
