const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get query execution statistics by query id.
 *
 * @summary Get query execution statistics by query id.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceQueryStatisticsList.json
 */
async function obtainQueryExecutionStatistics() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const databaseName = "database_1";
  const queryId = "42";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedDatabaseQueries.listByQuery(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    queryId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

obtainQueryExecutionStatistics().catch(console.error);
