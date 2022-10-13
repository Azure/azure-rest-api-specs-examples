const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Migrate an Azure Cosmos DB Gremlin graph from manual throughput to autoscale
 *
 * @summary Migrate an Azure Cosmos DB Gremlin graph from manual throughput to autoscale
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBGremlinGraphMigrateToAutoscale.json
 */
async function cosmosDbGremlinGraphMigrateToAutoscale() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const graphName = "graphName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.gremlinResources.beginMigrateGremlinGraphToAutoscaleAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    graphName
  );
  console.log(result);
}

cosmosDbGremlinGraphMigrateToAutoscale().catch(console.error);
