const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Migrate an Azure Cosmos DB MongoDB database from autoscale to manual throughput
 *
 * @summary Migrate an Azure Cosmos DB MongoDB database from autoscale to manual throughput
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBMongoDBDatabaseMigrateToManualThroughput.json
 */
async function cosmosDbMongoDbdatabaseMigrateToManualThroughput() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoDBResources.beginMigrateMongoDBDatabaseToManualThroughputAndWait(
    resourceGroupName,
    accountName,
    databaseName
  );
  console.log(result);
}

cosmosDbMongoDbdatabaseMigrateToManualThroughput().catch(console.error);
