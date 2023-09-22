const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Imports a bacpac into a new database.
 *
 * @summary Imports a bacpac into a new database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/ImportNewDatabase.json
 */
async function importsToANewDatabase() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const parameters = {
    administratorLogin: "login",
    administratorLoginPassword: "password",
    authenticationType: "Sql",
    databaseName: "testdb",
    storageKey:
      "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx==",
    storageKeyType: "StorageAccessKey",
    storageUri: "https://test.blob.core.windows.net/test.bacpac",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.servers.beginImportDatabaseAndWait(
    resourceGroupName,
    serverName,
    parameters
  );
  console.log(result);
}
