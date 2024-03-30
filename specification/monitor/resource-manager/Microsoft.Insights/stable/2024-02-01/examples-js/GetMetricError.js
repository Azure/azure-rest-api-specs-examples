const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **Lists the metric values for a resource**.
 *
 * @summary **Lists the metric values for a resource**.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2024-02-01/examples/GetMetricError.json
 */
async function getMetricWithError() {
  const resourceUri =
    "subscriptions/ac41e21f-afd6-4a79-8070-f01eba278f97/resourceGroups/todking/providers/Microsoft.DocumentDb/databaseAccounts/tk-cosmos-mongo";
  const timespan = "2021-06-07T21:51:00Z/2021-06-08T01:51:00Z";
  const interval = "FULL";
  const metricnames = "MongoRequestsCount,MongoRequests";
  const aggregation = "average";
  const metricnamespace = "microsoft.documentdb/databaseaccounts";
  const autoAdjustTimegrain = true;
  const validateDimensions = false;
  const options = {
    timespan,
    interval,
    metricnames,
    aggregation,
    metricnamespace,
    autoAdjustTimegrain,
    validateDimensions,
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.metricsOperations.list(resourceUri, options);
  console.log(result);
}
