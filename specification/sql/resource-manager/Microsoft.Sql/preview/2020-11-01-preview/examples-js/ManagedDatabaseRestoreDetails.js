const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets managed database restore details.
 *
 * @summary Gets managed database restore details.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseRestoreDetails.json
 */
async function managedDatabaseRestoreDetails() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const managedInstanceName = "managedInstance";
  const databaseName = "testdb";
  const restoreDetailsName = "Default";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabaseRestoreDetails.get(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    restoreDetailsName
  );
  console.log(result);
}

managedDatabaseRestoreDetails().catch(console.error);
