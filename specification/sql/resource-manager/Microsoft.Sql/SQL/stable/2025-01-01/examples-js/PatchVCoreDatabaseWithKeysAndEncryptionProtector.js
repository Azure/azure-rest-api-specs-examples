const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing database.
 *
 * @summary updates an existing database.
 * x-ms-original-file: 2025-01-01/PatchVCoreDatabaseWithKeysAndEncryptionProtector.json
 */
async function patchADatabaseWithDatabaseLevelCustomerManagedKeys() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.update("Default-SQL-SouthEastAsia", "testsvr", "testdb", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi":
          {},
      },
    },
    encryptionProtector: "https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion",
    keys: { "https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion": {} },
    sku: { name: "S0", tier: "Standard" },
  });
  console.log(result);
}
