const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the RUs per second of the MongoDB database under an existing Azure Cosmos DB database account with the provided name.
 *
 * @summary Gets the RUs per second of the MongoDB database under an existing Azure Cosmos DB database account with the provided name.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBMongoDBDatabaseThroughputGet.json
 */
async function cosmosDbMongoDbdatabaseThroughputGet() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoDBResources.getMongoDBDatabaseThroughput(
    resourceGroupName,
    accountName,
    databaseName,
  );
  console.log(result);
}
