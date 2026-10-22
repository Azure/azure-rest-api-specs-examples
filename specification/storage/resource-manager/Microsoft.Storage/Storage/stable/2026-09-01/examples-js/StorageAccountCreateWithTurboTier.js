const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 *
 * @summary asynchronously creates a new storage account with the specified parameters. If an account is already created and a subsequent create request is issued with different properties, the account properties will be updated. If an account is already created and a subsequent create or update request is issued with the exact same set of properties, the request will succeed.
 * x-ms-original-file: 2026-09-01/StorageAccountCreateWithTurboTier.json
 */
async function storageAccountCreateWithTurboTier() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.create("res9101", "staturbocontoso01", {
    kind: "StorageV2",
    location: "eastus2",
    accessTier: "Hot",
    turboTier: { status: "Enabled", targetPercent: 10 },
    sku: { name: "Standard_LRS" },
  });
  console.log(result);
}
