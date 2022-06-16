const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List managed database tables
 *
 * @summary List managed database tables
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseTableListBySchema.json
 */
async function listManagedDatabaseTables() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const managedInstanceName = "myManagedInstanceName";
  const databaseName = "myDatabase";
  const schemaName = "dbo";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedDatabaseTables.listBySchema(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    schemaName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedDatabaseTables().catch(console.error);
