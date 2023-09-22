const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new database or updates an existing database.
 *
 * @summary Creates a new database or updates an existing database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/CreateDatabaseDefaultModeWithKeysAndEncryptionProtector.json
 */
async function createsADatabaseWithDatabaseLevelCustomerManagedKeys() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const databaseName = "testdb";
  const parameters = {
    collation: "SQL_Latin1_General_CP1_CI_AS",
    createMode: "Default",
    encryptionProtector: "https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion",
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000111122223333444444444444/resourcegroups/DefaultSqlSouthEastAsia/providers/MicrosoftManagedIdentity/userAssignedIdentities/umi":
          {},
      },
    },
    keys: {
      "https://yourKeyVaultNameVaultAzureNet/yourKey/yourKeyVersion": {},
      "https://yourKeyVaultNameVaultAzureNet/yourKey2/yourKey2Version": {},
    },
    location: "southeastasia",
    maxSizeBytes: 1073741824,
    sku: { name: "S0", tier: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    parameters
  );
  console.log(result);
}
