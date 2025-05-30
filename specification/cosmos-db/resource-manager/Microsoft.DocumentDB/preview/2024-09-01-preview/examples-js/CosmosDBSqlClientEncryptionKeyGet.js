const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the ClientEncryptionKey under an existing Azure Cosmos DB SQL database.
 *
 * @summary Gets the ClientEncryptionKey under an existing Azure Cosmos DB SQL database.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBSqlClientEncryptionKeyGet.json
 */
async function cosmosDbClientEncryptionKeyGet() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subId";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rgName";
  const accountName = "accountName";
  const databaseName = "databaseName";
  const clientEncryptionKeyName = "cekName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.getClientEncryptionKey(
    resourceGroupName,
    accountName,
    databaseName,
    clientEncryptionKeyName,
  );
  console.log(result);
}
