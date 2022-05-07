Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all managed instances in an instance pool.
 *
 * @summary Gets a list of all managed instances in an instance pool.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceListByInstancePoolWithExpandEqualsAdministrators.json
 */
async function listManagedInstancesByInstancePoolWithExpandAdministratorsOrActivedirectory() {
  const subscriptionId = "20D7082A-0FC7-4468-82BD-542694D5042B";
  const resourceGroupName = "Test1";
  const instancePoolName = "pool1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedInstances.listByInstancePool(
    resourceGroupName,
    instancePoolName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedInstancesByInstancePoolWithExpandAdministratorsOrActivedirectory().catch(console.error);
```
