const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a managed database's short term retention policy list.
 *
 * @summary Gets a managed database's short term retention policy list.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetListManagedShortTermRetentionPolicy.json
 */
async function getTheShortTermRetentionPolicyListForTheDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const managedInstanceName = "testsvr";
  const databaseName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedBackupShortTermRetentionPolicies.listByDatabase(
    resourceGroupName,
    managedInstanceName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getTheShortTermRetentionPolicyListForTheDatabase().catch(console.error);
