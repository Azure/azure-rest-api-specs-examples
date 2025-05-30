const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or updates Azure Cosmos DB MongoDB database
 *
 * @summary Create or updates Azure Cosmos DB MongoDB database
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBMongoDBDatabaseCreateUpdate.json
 */
async function cosmosDbMongoDbdatabaseCreateUpdate() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const createUpdateMongoDBDatabaseParameters = {
    location: "West US",
    options: {},
    resource: { id: "databaseName" },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoDBResources.beginCreateUpdateMongoDBDatabaseAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    createUpdateMongoDBDatabaseParameters,
  );
  console.log(result);
}
