const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 *
 * @summary asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 * x-ms-original-file: 2026-04-01/StorageAccountCreateWithDataCollaborationPolicy.json
 */
async function storageAccountCreateWithDataCollaborationPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.create("res9101", "sto4445", {
    kind: "Storage",
    location: "eastus",
    dataCollaborationPolicyProperties: {
      allowStorageConnectors: true,
      allowStorageDataShares: true,
      allowCrossTenantDataSharing: false,
    },
    sku: { name: "Standard_GRS" },
  });
  console.log(result);
}
