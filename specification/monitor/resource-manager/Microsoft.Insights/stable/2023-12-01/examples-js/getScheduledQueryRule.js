const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve an scheduled query rule definition.
 *
 * @summary Retrieve an scheduled query rule definition.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2023-12-01/examples/getScheduledQueryRule.json
 */
async function getAScheduledQueryRuleForSingleResource() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "dd4bfc94-a096-412b-9c43-4bd13e35afbc";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "QueryResourceGroupName";
  const ruleName = "perf";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.scheduledQueryRules.get(resourceGroupName, ruleName);
  console.log(result);
}
