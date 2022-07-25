const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of all Activity Log Alert rules in a subscription.
 *
 * @summary Get a list of all Activity Log Alert rules in a subscription.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_ListBySubscriptionId.json
 */
async function getListOfAllActivityLogAlertRulesUnderASubscription() {
  const subscriptionId = "187f412d-1758-44d9-b052-169e2564721d";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.activityLogAlerts.listBySubscriptionId()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getListOfAllActivityLogAlertRulesUnderASubscription().catch(console.error);
