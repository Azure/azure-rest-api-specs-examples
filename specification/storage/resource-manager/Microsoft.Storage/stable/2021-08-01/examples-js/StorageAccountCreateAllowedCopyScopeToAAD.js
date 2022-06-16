const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountCreateAllowedCopyScopeToAad() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9101";
  const accountName = "sto4445";
  const parameters = {
    allowBlobPublicAccess: false,
    allowSharedKeyAccess: true,
    allowedCopyScope: "AAD",
    encryption: {
      keySource: "Microsoft.Storage",
      requireInfrastructureEncryption: false,
      services: {
        blob: { enabled: true, keyType: "Account" },
        file: { enabled: true, keyType: "Account" },
      },
    },
    isHnsEnabled: true,
    keyPolicy: { keyExpirationPeriodInDays: 20 },
    kind: "Storage",
    location: "eastus",
    minimumTlsVersion: "TLS1_2",
    routingPreference: {
      publishInternetEndpoints: true,
      publishMicrosoftEndpoints: true,
      routingChoice: "MicrosoftRouting",
    },
    sasPolicy: { expirationAction: "Log", sasExpirationPeriod: "1.15:59:59" },
    sku: { name: "Standard_GRS" },
    tags: { key1: "value1", key2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.beginCreateAndWait(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

storageAccountCreateAllowedCopyScopeToAad().catch(console.error);
