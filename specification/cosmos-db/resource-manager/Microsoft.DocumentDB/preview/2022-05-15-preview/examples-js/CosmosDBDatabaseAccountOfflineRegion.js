const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Offline the specified region for the specified Azure Cosmos DB database account.
 *
 * @summary Offline the specified region for the specified Azure Cosmos DB database account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-05-15-preview/examples/CosmosDBDatabaseAccountOfflineRegion.json
 */
async function cosmosDbDatabaseAccountOfflineRegion() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const regionParameterForOffline = {
    region: "",
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.databaseAccounts.beginOfflineRegionAndWait(
    resourceGroupName,
    accountName,
    regionParameterForOffline
  );
  console.log(result);
}

cosmosDbDatabaseAccountOfflineRegion().catch(console.error);
