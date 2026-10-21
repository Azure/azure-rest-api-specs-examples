const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Context Cache.
 *
 * @summary create or update a Context Cache.
 * x-ms-original-file: 2026-09-01/StorageContextCacheCRUD/ContextCaches_CreateOrUpdate_SystemAssignedUserAssigned.json
 */
async function createAContextCacheWithSystemAssignedAndUserAssignedIdentity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.contextCaches.createOrUpdate("testrg", "testcontextcache", {
    location: "eastus",
    tags: { environment: "test" },
    identity: {
      type: "SystemAssigned,UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-uami":
          {},
      },
    },
    properties: {
      accountKind: "Regional",
      description: "Test context cache",
      encryption: {
        customerManagedKeyEncryption: {
          keyEncryptionKeyIdentity: { identityType: "systemAssignedIdentity" },
          keyEncryptionKeyUrl: "https://mykeyvault.vault.azure.net/keys/encryptionKey",
        },
      },
    },
  });
  console.log(result);
}
