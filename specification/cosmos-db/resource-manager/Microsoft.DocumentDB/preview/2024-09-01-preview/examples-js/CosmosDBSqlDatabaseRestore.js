const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an Azure Cosmos DB SQL database
 *
 * @summary Create or update an Azure Cosmos DB SQL database
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBSqlDatabaseRestore.json
 */
async function cosmosDbSqlDatabaseRestore() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const createUpdateSqlDatabaseParameters = {
    location: "West US",
    options: {},
    resource: {
      createMode: "Restore",
      id: "databaseName",
      restoreParameters: {
        restoreSource:
          "/subscriptions/subid/providers/Microsoft.DocumentDB/locations/WestUS/restorableDatabaseAccounts/restorableDatabaseAccountId",
        restoreTimestampInUtc: new Date("2022-07-20T18:28:00Z"),
        restoreWithTtlDisabled: true,
      },
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.beginCreateUpdateSqlDatabaseAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    createUpdateSqlDatabaseParameters,
  );
  console.log(result);
}
