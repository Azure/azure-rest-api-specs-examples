const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get query execution statistics by query id.
 *
 * @summary Get query execution statistics by query id.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceQueryStatisticsListMin.json
 */
async function obtainQueryExecutionStatisticsMinimalExampleWithOnlyMandatoryRequestParameters() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const databaseName = "database_1";
  const queryId = "42";
  const interval = "PT1H";
  const options = { interval };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedDatabaseQueries.listByQuery(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    queryId,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
