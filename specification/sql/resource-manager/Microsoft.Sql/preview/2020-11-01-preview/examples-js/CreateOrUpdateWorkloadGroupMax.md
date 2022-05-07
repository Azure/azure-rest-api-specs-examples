Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a workload group.
 *
 * @summary Creates or updates a workload group.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateWorkloadGroupMax.json
 */
async function createAWorkloadGroupWithAllPropertiesSpecified() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const databaseName = "testdb";
  const workloadGroupName = "smallrc";
  const parameters = {
    importance: "normal",
    maxResourcePercent: 100,
    maxResourcePercentPerRequest: 3,
    minResourcePercent: 0,
    minResourcePercentPerRequest: 3,
    queryExecutionTimeout: 0,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.workloadGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    workloadGroupName,
    parameters
  );
  console.log(result);
}

createAWorkloadGroupWithAllPropertiesSpecified().catch(console.error);
```
