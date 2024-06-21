const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the Azure Cosmos DB accounts available under the subscription.
 *
 * @summary Lists all the Azure Cosmos DB accounts available under the subscription.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/throughputPool/CosmosDBThroughputPoolAccountsList.json
 */
async function cosmosDbThroughputPoolAccountList() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rgName";
  const throughputPoolName = "tp1";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.throughputPoolAccounts.list(
    resourceGroupName,
    throughputPoolName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
