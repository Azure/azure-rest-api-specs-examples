const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of recoverable managed databases.
 *
 * @summary Gets a list of recoverable managed databases.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListRecoverableManagedDatabasesByServer.json
 */
async function listRecoverableDatabasesByManagedInstances() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Test1";
  const managedInstanceName = "managedInstance";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recoverableManagedDatabases.listByInstance(
    resourceGroupName,
    managedInstanceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRecoverableDatabasesByManagedInstances().catch(console.error);
