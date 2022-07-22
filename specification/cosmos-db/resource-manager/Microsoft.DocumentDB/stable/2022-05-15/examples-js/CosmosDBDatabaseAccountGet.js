const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the properties of an existing Azure Cosmos DB database account.
 *
 * @summary Retrieves the properties of an existing Azure Cosmos DB database account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBDatabaseAccountGet.json
 */
async function cosmosDbDatabaseAccountGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.databaseAccounts.get(resourceGroupName, accountName);
  console.log(result);
}

cosmosDbDatabaseAccountGet().catch(console.error);
