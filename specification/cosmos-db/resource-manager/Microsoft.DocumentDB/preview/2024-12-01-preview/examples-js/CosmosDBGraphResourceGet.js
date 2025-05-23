const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the Graph resource under an existing Azure Cosmos DB database account with the provided name.
 *
 * @summary Gets the Graph resource under an existing Azure Cosmos DB database account with the provided name.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBGraphResourceGet.json
 */
async function cosmosDbSqlDatabaseGet() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const graphName = "graphName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.graphResources.getGraph(resourceGroupName, accountName, graphName);
  console.log(result);
}
