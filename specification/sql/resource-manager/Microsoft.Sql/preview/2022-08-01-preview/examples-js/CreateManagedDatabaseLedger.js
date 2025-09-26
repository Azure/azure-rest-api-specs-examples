const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new database or updates an existing database.
 *
 * @summary Creates a new database or updates an existing database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/CreateManagedDatabaseLedger.json
 */
async function createsANewManagedDatabaseWithLedgerOn() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Default-SQL-SouthEastAsia";
  const managedInstanceName = "managedInstance";
  const databaseName = "managedDatabase";
  const parameters = {
    isLedgerOn: true,
    location: "southeastasia",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabases.beginCreateOrUpdateAndWait(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    parameters,
  );
  console.log(result);
}
