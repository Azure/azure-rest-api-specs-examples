const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Context Cache.
 *
 * @summary create or update a Context Cache.
 * x-ms-original-file: 2026-06-01/StorageContextCacheCRUD/ContextCaches_CreateOrUpdate.json
 */
async function createAContextCache() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.contextCaches.createOrUpdate("testrg", "testaccount", {
    location: "eastus",
    tags: { environment: "test" },
    properties: { accountKind: "Regional", description: "Test Azure Context Cache account" },
  });
  console.log(result);
}
