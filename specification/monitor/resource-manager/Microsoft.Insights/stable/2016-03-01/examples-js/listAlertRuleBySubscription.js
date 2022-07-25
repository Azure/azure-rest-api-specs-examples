const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the classic metric alert rules within a subscription.
 *
 * @summary List the classic metric alert rules within a subscription.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/listAlertRuleBySubscription.json
 */
async function listAlertRules() {
  const subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.alertRules.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAlertRules().catch(console.error);
