const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get query execution statistics by query id.
 *
 * @summary get query execution statistics by query id.
 * x-ms-original-file: 2025-01-01/ManagedInstanceQueryStatisticsList.json
 */
async function obtainQueryExecutionStatistics() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.managedDatabaseQueries.listByQuery(
    "sqlcrudtest-7398",
    "sqlcrudtest-4645",
    "database_1",
    "42",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
