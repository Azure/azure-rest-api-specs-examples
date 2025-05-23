const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Migrate an Azure Cosmos DB MongoDB collection from autoscale to manual throughput
 *
 * @summary Migrate an Azure Cosmos DB MongoDB collection from autoscale to manual throughput
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBMongoDBCollectionMigrateToManualThroughput.json
 */
async function cosmosDbMongoDbcollectionMigrateToManualThroughput() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const collectionName = "collectionName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result =
    await client.mongoDBResources.beginMigrateMongoDBCollectionToManualThroughputAndWait(
      resourceGroupName,
      accountName,
      databaseName,
      collectionName,
    );
  console.log(result);
}
