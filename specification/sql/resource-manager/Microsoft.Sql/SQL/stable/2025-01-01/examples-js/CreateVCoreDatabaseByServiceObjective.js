const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new database or updates an existing database.
 *
 * @summary creates a new database or updates an existing database.
 * x-ms-original-file: 2025-01-01/CreateVCoreDatabaseByServiceObjective.json
 */
async function createsAVCoreDatabaseBySpecifyingServiceObjectiveName() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.createOrUpdate(
    "Default-SQL-SouthEastAsia",
    "testsvr",
    "testdb",
    { location: "southeastasia", sku: { name: "BC", capacity: 2, family: "Gen4" } },
  );
  console.log(result);
}
