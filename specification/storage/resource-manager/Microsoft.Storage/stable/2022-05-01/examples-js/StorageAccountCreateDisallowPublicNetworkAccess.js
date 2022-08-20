const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 *
 * @summary Asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountCreateDisallowPublicNetworkAccess.json
 */
async function storageAccountCreateDisallowPublicNetworkAccess() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9101";
  const accountName = "sto4445";
  const parameters = {
    allowBlobPublicAccess: false,
    allowSharedKeyAccess: true,
    encryption: {
      keySource: "Microsoft.Storage",
      requireInfrastructureEncryption: false,
      services: {
        blob: { enabled: true, keyType: "Account" },
        file: { enabled: true, keyType: "Account" },
      },
    },
    extendedLocation: { name: "losangeles001", type: "EdgeZone" },
    isHnsEnabled: true,
    keyPolicy: { keyExpirationPeriodInDays: 20 },
    kind: "Storage",
    location: "eastus",
    minimumTlsVersion: "TLS1_2",
    publicNetworkAccess: "Disabled",
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

storageAccountCreateDisallowPublicNetworkAccess().catch(console.error);
