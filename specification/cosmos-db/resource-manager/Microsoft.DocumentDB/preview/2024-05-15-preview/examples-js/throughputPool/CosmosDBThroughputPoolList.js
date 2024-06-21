const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the Azure Cosmos DB Throughput Pools available under the subscription.
 *
 * @summary Lists all the Azure Cosmos DB Throughput Pools available under the subscription.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/throughputPool/CosmosDBThroughputPoolList.json
 */
async function cosmosDbThroughputPoolList() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.throughputPools.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
