Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of workload classifiers for a workload group
 *
 * @summary Gets the list of workload classifiers for a workload group
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetWorkloadClassifierList.json
 */
async function getTheListOfWorkloadClassifiersForAWorkloadGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const databaseName = "testdb";
  const workloadGroupName = "wlm_workloadgroup";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workloadClassifiers.listByWorkloadGroup(
    resourceGroupName,
    serverName,
    databaseName,
    workloadGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getTheListOfWorkloadClassifiersForAWorkloadGroup().catch(console.error);
```
