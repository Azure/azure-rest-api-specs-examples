const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update an Azure Cosmos DB Cassandra Table
 *
 * @summary Create or update an Azure Cosmos DB Cassandra Table
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBCassandraTableCreateUpdate.json
 */
async function cosmosDbCassandraTableCreateUpdate() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const keyspaceName = "keyspaceName";
  const tableName = "tableName";
  const createUpdateCassandraTableParameters = {
    location: "West US",
    options: {},
    resource: {
      schema: {
        clusterKeys: [{ name: "columnA", orderBy: "Asc" }],
        columns: [{ name: "columnA", type: "Ascii" }],
        partitionKeys: [{ name: "columnA" }],
      },
      analyticalStorageTtl: 500,
      defaultTtl: 100,
      id: "tableName",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.cassandraResources.beginCreateUpdateCassandraTableAndWait(
    resourceGroupName,
    accountName,
    keyspaceName,
    tableName,
    createUpdateCassandraTableParameters,
  );
  console.log(result);
}
