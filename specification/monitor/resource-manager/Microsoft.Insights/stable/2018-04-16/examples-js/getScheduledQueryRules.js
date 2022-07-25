const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an Log Search rule
 *
 * @summary Gets an Log Search rule
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/getScheduledQueryRules.json
 */
async function getRule() {
  const subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const resourceGroupName = "Rac46PostSwapRG";
  const ruleName = "logalertfoo";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.scheduledQueryRules.get(resourceGroupName, ruleName);
  console.log(result);
}

getRule().catch(console.error);
