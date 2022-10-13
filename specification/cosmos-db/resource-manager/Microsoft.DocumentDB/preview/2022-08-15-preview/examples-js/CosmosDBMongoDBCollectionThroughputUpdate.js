const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the RUs per second of an Azure Cosmos DB MongoDB collection
 *
 * @summary Update the RUs per second of an Azure Cosmos DB MongoDB collection
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBMongoDBCollectionThroughputUpdate.json
 */
async function cosmosDbMongoDbcollectionThroughputUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const collectionName = "collectionName";
  const updateThroughputParameters = {
    location: "West US",
    resource: { throughput: 400 },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoDBResources.beginUpdateMongoDBCollectionThroughputAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    collectionName,
    updateThroughputParameters
  );
  console.log(result);
}

cosmosDbMongoDbcollectionThroughputUpdate().catch(console.error);
