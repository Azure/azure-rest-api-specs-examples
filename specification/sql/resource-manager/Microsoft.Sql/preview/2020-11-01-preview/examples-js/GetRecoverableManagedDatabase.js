const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a recoverable managed database.
 *
 * @summary Gets a recoverable managed database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetRecoverableManagedDatabase.json
 */
async function getsARecoverableDatabasesByManagedInstances() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Test1";
  const managedInstanceName = "managedInstance";
  const recoverableDatabaseName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.recoverableManagedDatabases.get(
    resourceGroupName,
    managedInstanceName,
    recoverableDatabaseName
  );
  console.log(result);
}

getsARecoverableDatabasesByManagedInstances().catch(console.error);
