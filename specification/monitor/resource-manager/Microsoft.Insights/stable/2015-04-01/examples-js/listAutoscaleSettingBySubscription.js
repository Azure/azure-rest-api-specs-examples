const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the autoscale settings for a subscription
 *
 * @summary Lists the autoscale settings for a subscription
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/listAutoscaleSettingBySubscription.json
 */
async function listAutoscaleSettings() {
  const subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.autoscaleSettings.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAutoscaleSettings().catch(console.error);
