const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing Azure Cosmos DB Gremlin graph.
 *
 * @summary Deletes an existing Azure Cosmos DB Gremlin graph.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBGremlinGraphDelete.json
 */
async function cosmosDbGremlinGraphDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const graphName = "graphName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.gremlinResources.beginDeleteGremlinGraphAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    graphName
  );
  console.log(result);
}

cosmosDbGremlinGraphDelete().catch(console.error);
