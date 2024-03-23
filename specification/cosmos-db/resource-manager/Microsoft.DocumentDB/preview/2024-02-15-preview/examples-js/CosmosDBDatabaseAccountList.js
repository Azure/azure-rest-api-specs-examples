const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the Azure Cosmos DB database accounts available under the subscription.
 *
 * @summary Lists all the Azure Cosmos DB database accounts available under the subscription.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-02-15-preview/examples/CosmosDBDatabaseAccountList.json
 */
async function cosmosDbDatabaseAccountList() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.databaseAccounts.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
