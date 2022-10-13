const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Cassandra view under an existing Azure Cosmos DB database account.
 *
 * @summary Gets the Cassandra view under an existing Azure Cosmos DB database account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBCassandraViewGet.json
 */
async function cosmosDbCassandraViewGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const keyspaceName = "keyspacename";
  const viewName = "viewname";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.cassandraResources.getCassandraView(
    resourceGroupName,
    accountName,
    keyspaceName,
    viewName
  );
  console.log(result);
}

cosmosDbCassandraViewGet().catch(console.error);
