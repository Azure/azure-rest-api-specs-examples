const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List database tables
 *
 * @summary List database tables
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseTableListBySchema.json
 */
async function listDatabaseTables() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const serverName = "serverName";
  const databaseName = "myDatabase";
  const schemaName = "dbo";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.databaseTables.listBySchema(
    resourceGroupName,
    serverName,
    databaseName,
    schemaName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDatabaseTables().catch(console.error);
