const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update log search Rule.
 *
 * @summary Update log search Rule.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/patchScheduledQueryRules.json
 */
async function patchLogSearchRule() {
  const subscriptionId = process.env["MONITOR_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "my-resource-group";
  const ruleName = "logalertfoo";
  const parameters = { enabled: "true" };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.scheduledQueryRules.update(resourceGroupName, ruleName, parameters);
  console.log(result);
}
