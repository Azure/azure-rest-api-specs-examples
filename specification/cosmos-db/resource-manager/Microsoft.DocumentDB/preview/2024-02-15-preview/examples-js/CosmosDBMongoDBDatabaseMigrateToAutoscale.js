const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Migrate an Azure Cosmos DB MongoDB database from manual throughput to autoscale
 *
 * @summary Migrate an Azure Cosmos DB MongoDB database from manual throughput to autoscale
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-02-15-preview/examples/CosmosDBMongoDBDatabaseMigrateToAutoscale.json
 */
async function cosmosDbMongoDbdatabaseMigrateToAutoscale() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoDBResources.beginMigrateMongoDBDatabaseToAutoscaleAndWait(
    resourceGroupName,
    accountName,
    databaseName,
  );
  console.log(result);
}
