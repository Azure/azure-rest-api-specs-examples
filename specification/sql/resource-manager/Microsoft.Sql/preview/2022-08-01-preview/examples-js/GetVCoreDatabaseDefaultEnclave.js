const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a database.
 *
 * @summary Gets a database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/GetVCoreDatabaseDefaultEnclave.json
 */
async function getsADatabaseConfiguredWithDefaultEnclaveType() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const databaseName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.get(resourceGroupName, serverName, databaseName);
  console.log(result);
}
