const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to imports a bacpac into a new database.
 *
 * @summary imports a bacpac into a new database.
 * x-ms-original-file: 2025-01-01/ImportDatabase.json
 */
async function importsToAnExistingEmptyDatabase() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.import("Default-SQL-SouthEastAsia", "testsvr", "testdb", {
    administratorLogin: "login",
    administratorLoginPassword: "password",
    authenticationType: "Sql",
    storageKey:
      "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx==",
    storageKeyType: "StorageAccessKey",
    storageUri: "https://test.blob.core.windows.net/test.bacpac",
  });
  console.log(result);
}
