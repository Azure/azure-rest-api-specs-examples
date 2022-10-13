const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates an access key for the specified Azure Cosmos DB database account.
 *
 * @summary Regenerates an access key for the specified Azure Cosmos DB database account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBDatabaseAccountRegenerateKey.json
 */
async function cosmosDbDatabaseAccountRegenerateKey() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const keyToRegenerate = {
    keyKind: "primary",
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.databaseAccounts.beginRegenerateKeyAndWait(
    resourceGroupName,
    accountName,
    keyToRegenerate
  );
  console.log(result);
}

cosmosDbDatabaseAccountRegenerateKey().catch(console.error);
