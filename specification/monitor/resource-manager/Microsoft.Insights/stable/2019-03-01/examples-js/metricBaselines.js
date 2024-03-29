const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **Lists the metric baseline values for a resource**.
 *
 * @summary **Lists the metric baseline values for a resource**.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2019-03-01/examples/metricBaselines.json
 */
async function getMetricBaselines() {
  const resourceUri =
    "subscriptions/b368ca2f-e298-46b7-b0ab-012281956afa/resourceGroups/vms/providers/Microsoft.Compute/virtualMachines/vm1";
  const timespan = "2019-03-12T11:00:00.000Z/2019-03-12T12:00:00.000Z";
  const interval = "PT1H";
  const aggregation = "average";
  const sensitivities = "Low,Medium";
  const options = {
    timespan,
    interval,
    aggregation,
    sensitivities,
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const resArray = new Array();
  for await (let item of client.baselines.list(resourceUri, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
