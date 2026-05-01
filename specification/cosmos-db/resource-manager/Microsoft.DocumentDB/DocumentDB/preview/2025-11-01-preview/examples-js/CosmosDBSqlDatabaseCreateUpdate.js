const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an Azure Cosmos DB SQL database
 *
 * @summary create or update an Azure Cosmos DB SQL database
 * x-ms-original-file: 2025-11-01-preview/CosmosDBSqlDatabaseCreateUpdate.json
 */
async function cosmosDBSqlDatabaseCreateUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.createUpdateSqlDatabase("rg1", "ddb1", "databaseName", {
    location: "West US",
    options: {},
    resource: { id: "databaseName" },
    tags: {},
  });
  console.log(result);
}
