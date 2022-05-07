Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get query execution statistics by query id.
 *
 * @summary Get query execution statistics by query id.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceQueryStatisticsListMax.json
 */
async function obtainQueryExecutionStatisticsExampleWithAllRequestParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const databaseName = "database_1";
  const queryId = "42";
  const startTime = "03/01/2020 16:23:09";
  const endTime = "03/11/2020 14:00:00";
  const interval = "P1D";
  const options = {
    startTime,
    endTime,
    interval,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedDatabaseQueries.listByQuery(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    queryId,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

obtainQueryExecutionStatisticsExampleWithAllRequestParameters().catch(console.error);
```
