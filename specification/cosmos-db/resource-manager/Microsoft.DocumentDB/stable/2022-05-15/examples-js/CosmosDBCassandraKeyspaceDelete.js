const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing Azure Cosmos DB Cassandra keyspace.
 *
 * @summary Deletes an existing Azure Cosmos DB Cassandra keyspace.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBCassandraKeyspaceDelete.json
 */
async function cosmosDbCassandraKeyspaceDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const keyspaceName = "keyspaceName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.cassandraResources.beginDeleteCassandraKeyspaceAndWait(
    resourceGroupName,
    accountName,
    keyspaceName
  );
  console.log(result);
}

cosmosDbCassandraKeyspaceDelete().catch(console.error);
