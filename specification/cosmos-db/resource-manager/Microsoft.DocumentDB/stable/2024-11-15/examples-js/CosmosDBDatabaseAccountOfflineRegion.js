const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Offline the specified region for the specified Azure Cosmos DB database account.
 *
 * @summary Offline the specified region for the specified Azure Cosmos DB database account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBDatabaseAccountOfflineRegion.json
 */
async function cosmosDbDatabaseAccountOfflineRegion() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const regionParameterForOffline = {
    region: "North Europe",
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.databaseAccounts.beginOfflineRegionAndWait(
    resourceGroupName,
    accountName,
    regionParameterForOffline,
  );
  console.log(result);
}
