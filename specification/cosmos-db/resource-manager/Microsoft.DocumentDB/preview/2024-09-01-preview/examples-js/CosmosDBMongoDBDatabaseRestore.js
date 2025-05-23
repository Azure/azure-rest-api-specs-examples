const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or updates Azure Cosmos DB MongoDB database
 *
 * @summary Create or updates Azure Cosmos DB MongoDB database
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBMongoDBDatabaseRestore.json
 */
async function cosmosDbMongoDbdatabaseRestore() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const createUpdateMongoDBDatabaseParameters = {
    location: "West US",
    options: {},
    resource: {
      createMode: "Restore",
      id: "databaseName",
      restoreParameters: {
        restoreSource:
          "/subscriptions/subid/providers/Microsoft.DocumentDB/locations/WestUS/restorableDatabaseAccounts/restorableDatabaseAccountId",
        restoreTimestampInUtc: new Date("2022-07-20T18:28:00Z"),
        restoreWithTtlDisabled: false,
      },
    },
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
