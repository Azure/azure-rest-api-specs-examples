const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get predictive autoscale metric future data
 *
 * @summary get predictive autoscale metric future data
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/GetPredictiveMetric.json
 */
async function getMetricForData() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myRG";
  const autoscaleSettingName = "vmss1-Autoscale-775";
  const timespan = "2021-10-14T22:00:00.000Z/2021-10-16T22:00:00.000Z";
  const interval = "PT1H";
  const metricNamespace = "Microsoft.Compute/virtualMachineScaleSets";
  const metricName = "PercentageCPU";
  const aggregation = "Total";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.predictiveMetric.get(
    resourceGroupName,
    autoscaleSettingName,
    timespan,
    interval,
    metricNamespace,
    metricName,
    aggregation
  );
  console.log(result);
}

getMetricForData().catch(console.error);
