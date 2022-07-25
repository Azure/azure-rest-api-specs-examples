const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the Log Search rules within a subscription group.
 *
 * @summary List the Log Search rules within a subscription group.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/listScheduledQueryRules.json
 */
async function listRules() {
  const subscriptionId = "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.scheduledQueryRules.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRules().catch(console.error);
