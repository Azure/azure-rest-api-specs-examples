const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **Lists the metric values for a resource**.
 *
 * @summary **Lists the metric values for a resource**.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/GetMetricError.json
 */
async function getMetricWithError() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri =
    "subscriptions/ac41e21f-afd6-4a79-8070-f01eba278f97/resourceGroups/todking/providers/Microsoft.DocumentDb/databaseAccounts/tk-cosmos-mongo";
  const timespan = "2021-06-07T21:51:00Z/2021-06-08T01:51:00Z";
  const interval = "FULL";
  const metricnames = "MongoRequestsCount,MongoRequests";
  const aggregation = "average";
  const metricnamespace = "microsoft.documentdb/databaseaccounts";
  const options = {
    timespan,
    interval,
    metricnames,
    aggregation,
    metricnamespace,
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metrics.list(resourceUri, options);
  console.log(result);
}

getMetricWithError().catch(console.error);
