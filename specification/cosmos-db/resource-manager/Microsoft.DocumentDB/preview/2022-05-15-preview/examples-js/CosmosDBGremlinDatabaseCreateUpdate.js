const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an Azure Cosmos DB Gremlin database
 *
 * @summary Create or update an Azure Cosmos DB Gremlin database
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-05-15-preview/examples/CosmosDBGremlinDatabaseCreateUpdate.json
 */
async function cosmosDbGremlinDatabaseCreateUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const createUpdateGremlinDatabaseParameters = {
    location: "West US",
    options: {},
    resource: { id: "databaseName" },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.gremlinResources.beginCreateUpdateGremlinDatabaseAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    createUpdateGremlinDatabaseParameters
  );
  console.log(result);
}

cosmosDbGremlinDatabaseCreateUpdate().catch(console.error);
