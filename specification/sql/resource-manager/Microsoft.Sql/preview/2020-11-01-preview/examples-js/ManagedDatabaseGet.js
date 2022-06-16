const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a managed database.
 *
 * @summary Gets a managed database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseGet.json
 */
async function getsAManagedDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Test1";
  const managedInstanceName = "managedInstance";
  const databaseName = "managedDatabase";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabases.get(
    resourceGroupName,
    managedInstanceName,
    databaseName
  );
  console.log(result);
}

getsAManagedDatabase().catch(console.error);
