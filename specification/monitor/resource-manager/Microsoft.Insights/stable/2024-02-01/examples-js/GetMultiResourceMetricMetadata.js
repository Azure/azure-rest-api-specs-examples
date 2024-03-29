const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **Lists the metric data for a subscription**. Parameters can be specified on either query params or the body.
 *
 * @summary **Lists the metric data for a subscription**. Parameters can be specified on either query params or the body.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2024-02-01/examples/GetMultiResourceMetricMetadata.json
 */
async function postRequestForSubscriptionLevelMetricMetadata() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "92d2a2d8-b514-432d-8cc9-a5f9272630d5";
  const region = "westus2";
  const timespan = "2021-06-10T02:23:16.129Z/2021-06-12T02:23:16.129Z";
  const metricnames = "Data Disk Max Burst IOPS";
  const filter = "LUN eq '0'";
  const metricnamespace = "microsoft.compute/virtualmachines";
  const options = {
    timespan,
    metricnames,
    filter,
    metricnamespace,
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricsOperations.listAtSubscriptionScopePost(region, options);
  console.log(result);
}
