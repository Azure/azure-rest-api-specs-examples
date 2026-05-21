const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 *
 * @summary asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 * x-ms-original-file: 2025-08-01/StorageAccountCreateWithSmartAccessTier.json
 */
async function storageAccountCreateWithSmartAccessTier() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.create("res9101", "sto4445", {
    extendedLocation: { name: "losangeles001", type: "EdgeZone" },
    kind: "Storage",
    location: "eastus",
    accessTier: "Smart",
    allowBlobPublicAccess: false,
    allowSharedKeyAccess: true,
    defaultToOAuthAuthentication: false,
    encryption: {
      keySource: "Microsoft.Storage",
      requireInfrastructureEncryption: false,
      services: {
        blob: { enabled: true, keyType: "Account" },
        file: { enabled: true, keyType: "Account" },
      },
    },
    isHnsEnabled: true,
    isSftpEnabled: true,
    keyPolicy: { keyExpirationPeriodInDays: 20 },
    minimumTlsVersion: "TLS1_2",
    routingPreference: {
      publishInternetEndpoints: true,
      publishMicrosoftEndpoints: true,
      routingChoice: "MicrosoftRouting",
    },
    sasPolicy: { expirationAction: "Log", sasExpirationPeriod: "1.15:59:59" },
    geoPriorityReplicationStatus: { isBlobEnabled: true },
    sku: { name: "Standard_GRS" },
    tags: { key1: "value1", key2: "value2" },
  });
  console.log(result);
}
