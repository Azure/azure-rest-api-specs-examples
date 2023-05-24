const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all managed Cassandra clusters in this subscription.
 *
 * @summary List all managed Cassandra clusters in this subscription.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBManagedCassandraClusterListBySubscription.json
 */
async function cosmosDbManagedCassandraClusterListBySubscription() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.cassandraClusters.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
