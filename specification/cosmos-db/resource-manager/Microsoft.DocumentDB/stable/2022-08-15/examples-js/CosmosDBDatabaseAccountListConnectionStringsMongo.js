const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the connection strings for the specified Azure Cosmos DB database account.
 *
 * @summary Lists the connection strings for the specified Azure Cosmos DB database account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBDatabaseAccountListConnectionStringsMongo.json
 */
async function cosmosDbDatabaseAccountListConnectionStringsMongo() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "mongo-ddb1";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.databaseAccounts.listConnectionStrings(
    resourceGroupName,
    accountName
  );
  console.log(result);
}

cosmosDbDatabaseAccountListConnectionStringsMongo().catch(console.error);
