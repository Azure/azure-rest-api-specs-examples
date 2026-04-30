const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Azure Cosmos DB database account. The "Update" method is preferred when performing updates on an account.
 *
 * @summary creates or updates an Azure Cosmos DB database account. The "Update" method is preferred when performing updates on an account.
 * x-ms-original-file: 2025-11-01-preview/CosmosDBRestoreDatabaseAccountCreateUpdate.json
 */
async function cosmosDBRestoreDatabaseAccountCreateUpdateJson() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.databaseAccounts.createOrUpdate("rg1", "ddb1", {
    kind: "GlobalDocumentDB",
    location: "westus",
    apiProperties: { serverVersion: "3.2" },
    backupPolicy: { type: "Continuous", continuousModeProperties: { tier: "Continuous30Days" } },
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
    locations: [{ failoverPriority: 0, isZoneRedundant: false, locationName: "southcentralus" }],
    minimalTlsVersion: "Tls",
    restoreParameters: {
      databasesToRestore: [
        { collectionNames: ["collection1", "collection2"], databaseName: "db1" },
        { collectionNames: ["collection3", "collection4"], databaseName: "db2" },
      ],
      restoreMode: "PointInTime",
      restoreSource:
        "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/1a97b4bb-f6a0-430e-ade1-638d781830cc",
      restoreTimestampInUtc: new Date("2021-03-11T22:05:09Z"),
      restoreWithTtlDisabled: false,
      sourceBackupLocation: "westus",
    },
    tags: {},
  });
  console.log(result);
}
