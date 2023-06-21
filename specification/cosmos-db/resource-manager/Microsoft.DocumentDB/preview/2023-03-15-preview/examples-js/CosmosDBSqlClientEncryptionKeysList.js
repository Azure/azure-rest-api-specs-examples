const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the ClientEncryptionKeys under an existing Azure Cosmos DB SQL database.
 *
 * @summary Lists the ClientEncryptionKeys under an existing Azure Cosmos DB SQL database.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBSqlClientEncryptionKeysList.json
 */
async function cosmosDbClientEncryptionKeysList() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subId";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rgName";
  const accountName = "accountName";
  const databaseName = "databaseName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sqlResources.listClientEncryptionKeys(
    resourceGroupName,
    accountName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
