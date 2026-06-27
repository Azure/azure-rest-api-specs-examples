const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing database.
 *
 * @summary updates an existing database.
 * x-ms-original-file: 2025-01-01/PatchDatabaseVBSEnclave.json
 */
async function updatesADatabaseWithVBSEnclaveType() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databases.update("Default-SQL-SouthEastAsia", "testsvr", "testdb", {
    preferredEnclaveType: "VBS",
  });
  console.log(result);
}
