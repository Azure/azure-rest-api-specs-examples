const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update RUs per second of an Azure Cosmos DB Cassandra view
 *
 * @summary Update RUs per second of an Azure Cosmos DB Cassandra view
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBCassandraViewThroughputUpdate.json
 */
async function cosmosDbCassandraViewThroughputUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const keyspaceName = "keyspacename";
  const viewName = "viewname";
  const updateThroughputParameters = {
    resource: { throughput: 400 },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.cassandraResources.beginUpdateCassandraViewThroughputAndWait(
    resourceGroupName,
    accountName,
    keyspaceName,
    viewName,
    updateThroughputParameters
  );
  console.log(result);
}

cosmosDbCassandraViewThroughputUpdate().catch(console.error);
