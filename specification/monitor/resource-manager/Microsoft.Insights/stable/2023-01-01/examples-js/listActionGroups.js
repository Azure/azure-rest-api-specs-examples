const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of all action groups in a subscription.
 *
 * @summary Get a list of all action groups in a subscription.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2023-01-01/examples/listActionGroups.json
 */
async function listActionGroupsAtSubscriptionLevel() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "187f412d-1758-44d9-b052-169e2564721d";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.actionGroups.listBySubscriptionId()) {
    resArray.push(item);
  }
  console.log(resArray);
}
