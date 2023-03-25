const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing database.
 *
 * @summary Updates an existing database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/PatchVCoreDatabaseWithKeysAndEncryptionProtector.json
 */
async function patchADatabaseWithDatabaseLevelCustomerManagedKeys() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const databaseName = "testdb";
  const parameters = {
    encryptionProtector: "https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion",
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000111122223333444444444444/resourcegroups/DefaultSqlSouthEastAsia/providers/MicrosoftManagedIdentity/userAssignedIdentities/umi":
          {},
        "/subscriptions/00000000111122223333444444444444/resourcegroups/DefaultSqlSouthEastAsia/providers/MicrosoftManagedIdentity/userAssignedIdentities/umiToDelete":
          {},
      },
    },
    keys: {
      "https://yourKeyVaultNameVaultAzureNet/yourKey/yourKeyVersion": {},
      "https://yourKeyVaultNameVaultAzureNet/yourKey2/yourKey2VersionToDelete": {},
    },
    sku: { name: "S0", tier: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.beginUpdateAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    parameters
  );
  console.log(result);
}
