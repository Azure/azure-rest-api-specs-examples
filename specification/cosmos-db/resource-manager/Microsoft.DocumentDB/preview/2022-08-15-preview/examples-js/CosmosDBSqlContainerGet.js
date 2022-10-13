const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the SQL container under an existing Azure Cosmos DB database account.
 *
 * @summary Gets the SQL container under an existing Azure Cosmos DB database account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBSqlContainerGet.json
 */
async function cosmosDbSqlContainerGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgName";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const containerName = "containerName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.getSqlContainer(
    resourceGroupName,
    accountName,
    databaseName,
    containerName
  );
  console.log(result);
}

cosmosDbSqlContainerGet().catch(console.error);
