const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the autoscale settings for a resource group
 *
 * @summary Lists the autoscale settings for a resource group
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/listAutoscaleSetting.json
 */
async function listAutoscaleSettings() {
  const subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const resourceGroupName = "TestingMetricsScaleSet";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.autoscaleSettings.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAutoscaleSettings().catch(console.error);
