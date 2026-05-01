const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Azure Cosmos DB database account. The "Update" method is preferred when performing updates on an account.
 *
 * @summary creates or updates an Azure Cosmos DB database account. The "Update" method is preferred when performing updates on an account.
 * x-ms-original-file: 2025-11-01-preview/CosmosDBDatabaseAccountCreateMin.json
 */
async function cosmosDBDatabaseAccountCreateMin() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.databaseAccounts.createOrUpdate("rg1", "ddb1", {
    location: "westus",
    databaseAccountOfferType: "Standard",
    locations: [{ failoverPriority: 0, locationName: "southcentralus", isZoneRedundant: false }],
    createMode: "Default",
  });
  console.log(result);
}
