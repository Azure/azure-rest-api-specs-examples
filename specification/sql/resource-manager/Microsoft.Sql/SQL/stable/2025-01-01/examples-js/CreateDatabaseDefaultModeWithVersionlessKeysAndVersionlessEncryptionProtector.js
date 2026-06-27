const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new database or updates an existing database.
 *
 * @summary creates a new database or updates an existing database.
 * x-ms-original-file: 2025-01-01/CreateDatabaseDefaultModeWithVersionlessKeysAndVersionlessEncryptionProtector.json
 */
async function createsADatabaseWithDatabaseLevelVersionlessCustomerManagedKeys() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.createOrUpdate(
    "Default-SQL-SouthEastAsia",
    "testsvr",
    "testdb",
    {
      identity: {
        type: "UserAssigned",
        userAssignedIdentities: {
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi":
            {},
        },
      },
      location: "southeastasia",
      collation: "SQL_Latin1_General_CP1_CI_AS",
      createMode: "Default",
      encryptionProtector: "https://your-key-vault-name.vault.azure.net/yourKey",
      keys: {
        "https://your-key-vault-name.vault.azure.net/yourKey": {},
        "https://your-key-vault-name.vault.azure.net/yourKey2": {},
      },
      maxSizeBytes: 1073741824,
      sku: { name: "S0", tier: "Standard" },
    },
  );
  console.log(result);
}
