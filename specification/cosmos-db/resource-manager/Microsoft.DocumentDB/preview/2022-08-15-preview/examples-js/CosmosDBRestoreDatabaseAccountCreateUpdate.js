const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Azure Cosmos DB database account. The "Update" method is preferred when performing updates on an account.
 *
 * @summary Creates or updates an Azure Cosmos DB database account. The "Update" method is preferred when performing updates on an account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBRestoreDatabaseAccountCreateUpdate.json
 */
async function cosmosDbRestoreDatabaseAccountCreateUpdateJson() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const createUpdateParameters = {
    apiProperties: { serverVersion: "3.2" },
    backupPolicy: {
      type: "Continuous",
      continuousModeProperties: { tier: "Continuous30Days" },
    },
    consistencyPolicy: {
      defaultConsistencyLevel: "BoundedStaleness",
      maxIntervalInSeconds: 10,
      maxStalenessPrefix: 200,
    },
    createMode: "Restore",
    databaseAccountOfferType: "Standard",
    enableAnalyticalStorage: true,
    enableFreeTier: false,
    enableMaterializedViews: false,
    keyVaultKeyUri: "https://myKeyVault.vault.azure.net",
    kind: "GlobalDocumentDB",
    location: "westus",
    locations: [
      {
        failoverPriority: 0,
        isZoneRedundant: false,
        locationName: "southcentralus",
      },
    ],
    restoreParameters: {
      databasesToRestore: [
        {
          collectionNames: ["collection1", "collection2"],
          databaseName: "db1",
        },
        { collectionNames: ["collection3", "collection4"], databaseName: "db2" },
      ],
      restoreMode: "PointInTime",
      restoreSource:
        "/subscriptions/subid/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/1a97b4bb-f6a0-430e-ade1-638d781830cc",
      restoreTimestampInUtc: new Date("2021-03-11T22:05:09Z"),
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.databaseAccounts.beginCreateOrUpdateAndWait(
    resourceGroupName,
    accountName,
    createUpdateParameters
  );
  console.log(result);
}

cosmosDbRestoreDatabaseAccountCreateUpdateJson().catch(console.error);
