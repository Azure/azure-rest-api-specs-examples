const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **Lists the metric data for a subscription**. Parameters can be specified on either query params or the body.
 *
 * @summary **Lists the metric data for a subscription**. Parameters can be specified on either query params or the body.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2024-02-01/examples/PostMultiResourceMetricBody.json
 */
async function postRequestForSubscriptionLevelMetricDataUsingBodyParams() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "92d2a2d8-b514-432d-8cc9-a5f9272630d5";
  const region = "westus2";
  const body = {
    aggregation: "count",
    autoAdjustTimegrain: true,
    filter: "LUN eq '0' and Microsoft.ResourceId eq '*'",
    interval: "PT6H",
    metricNames: "Data Disk Max Burst IOPS",
    metricNamespace: "microsoft.compute/virtualmachines",
    orderBy: "count desc",
    rollUpBy: "LUN",
    timespan: "2021-06-08T19:00:00Z/2021-06-12T01:00:00Z",
    top: 10,
    validateDimensions: false,
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricsOperations.listAtSubscriptionScopePost(region, options);
  console.log(result);
}
