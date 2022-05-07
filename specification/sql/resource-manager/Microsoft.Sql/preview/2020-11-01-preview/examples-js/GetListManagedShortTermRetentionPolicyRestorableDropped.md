Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a dropped database's short term retention policy list.
 *
 * @summary Gets a dropped database's short term retention policy list.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetListManagedShortTermRetentionPolicyRestorableDropped.json
 */
async function getTheShortTermRetentionPolicyListForTheDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const managedInstanceName = "testsvr";
  const restorableDroppedDatabaseId = "testdb,131403269876900000";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedRestorableDroppedDatabaseBackupShortTermRetentionPolicies.listByRestorableDroppedDatabase(
    resourceGroupName,
    managedInstanceName,
    restorableDroppedDatabaseId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getTheShortTermRetentionPolicyListForTheDatabase().catch(console.error);
```
