const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing Azure Cosmos DB Gremlin database.
 *
 * @summary Deletes an existing Azure Cosmos DB Gremlin database.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBGremlinDatabaseDelete.json
 */
async function cosmosDbGremlinDatabaseDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.gremlinResources.beginDeleteGremlinDatabaseAndWait(
    resourceGroupName,
    accountName,
    databaseName
  );
  console.log(result);
}

cosmosDbGremlinDatabaseDelete().catch(console.error);
