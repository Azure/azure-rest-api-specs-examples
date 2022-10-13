const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update RUs per second of an Azure Cosmos DB SQL database
 *
 * @summary Update RUs per second of an Azure Cosmos DB SQL database
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBSqlDatabaseThroughputUpdate.json
 */
async function cosmosDbSqlDatabaseThroughputUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const updateThroughputParameters = {
    location: "West US",
    resource: { throughput: 400 },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.beginUpdateSqlDatabaseThroughputAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    updateThroughputParameters
  );
  console.log(result);
}

cosmosDbSqlDatabaseThroughputUpdate().catch(console.error);
