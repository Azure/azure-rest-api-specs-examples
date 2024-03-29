const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a scheduled query rule definitions in a subscription.
 *
 * @summary Retrieve a scheduled query rule definitions in a subscription.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2023-12-01/examples/listScheduledQueryRulesBySubscription.json
 */
async function listScheduledQueryRulesBySubscription() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "dd4bfc94-a096-412b-9c43-4bd13e35afbc";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.scheduledQueryRules.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
