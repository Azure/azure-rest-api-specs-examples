const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of databases.
 *
 * @summary gets a list of databases.
 * x-ms-original-file: 2025-01-01/ListVCoreDatabasesByServer.json
 */
async function getsAListOfDatabases() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.databases.listByServer("Default-SQL-SouthEastAsia", "testsvr")) {
    resArray.push(item);
  }

  console.log(resArray);
}
