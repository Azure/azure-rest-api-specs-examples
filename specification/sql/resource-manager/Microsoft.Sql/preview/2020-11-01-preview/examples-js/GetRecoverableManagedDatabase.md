Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
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
```
