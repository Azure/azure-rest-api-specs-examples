const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes and autoscale setting
 *
 * @summary Deletes and autoscale setting
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/deleteAutoscaleSetting.json
 */
async function deleteAnAutoscaleSetting() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "TestingMetricsScaleSet";
  const autoscaleSettingName = "MySetting";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.autoscaleSettings.delete(resourceGroupName, autoscaleSettingName);
  console.log(result);
}
