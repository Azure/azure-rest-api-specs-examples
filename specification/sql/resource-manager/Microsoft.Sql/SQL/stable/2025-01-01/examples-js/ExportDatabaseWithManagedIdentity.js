const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to exports a database.
 *
 * @summary exports a database.
 * x-ms-original-file: 2025-01-01/ExportDatabaseWithManagedIdentity.json
 */
async function exportsADatabaseUsingManagedIdentityToCommunicateWithSQLServerAndStorageAccount() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.export("Default-SQL-SouthEastAsia", "testsvr", "testdb", {
    administratorLogin:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName",
    authenticationType: "ManagedIdentity",
    storageKey:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName",
    storageKeyType: "ManagedIdentity",
    storageUri: "https://test.blob.core.windows.net/test.bacpac",
  });
  console.log(result);
}
