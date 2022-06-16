const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new database or updates an existing database.
 *
 * @summary Creates a new database or updates an existing database.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/DatabaseCreate.json
 */
async function databaseCreate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TestGroup";
  const serverName = "testserver";
  const databaseName = "db1";
  const parameters = {
    charset: "UTF8",
    collation: "English_United States.1252",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const result = await client.databases.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    parameters
  );
  console.log(result);
}

databaseCreate().catch(console.error);
