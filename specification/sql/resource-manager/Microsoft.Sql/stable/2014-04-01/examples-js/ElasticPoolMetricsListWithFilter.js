const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns elastic pool  metrics.
 *
 * @summary Returns elastic pool  metrics.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ElasticPoolMetricsListWithFilter.json
 */
async function listDatabaseUsageMetrics() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-6730";
  const serverName = "sqlcrudtest-9007";
  const elasticPoolName = "3481";
  const filter =
    "name/value eq 'cpu_percent' and timeGrain eq '00:10:00' and startTime eq '2017-06-02T18:35:00Z' and endTime eq '2017-06-02T18:55:00Z'";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.elasticPools.listMetrics(
    resourceGroupName,
    serverName,
    elasticPoolName,
    filter
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDatabaseUsageMetrics().catch(console.error);
