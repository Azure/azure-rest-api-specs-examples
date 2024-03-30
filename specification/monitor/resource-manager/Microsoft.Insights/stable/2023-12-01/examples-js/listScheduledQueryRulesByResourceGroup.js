const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve scheduled query rule definitions in a resource group.
 *
 * @summary Retrieve scheduled query rule definitions in a resource group.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2023-12-01/examples/listScheduledQueryRulesByResourceGroup.json
 */
async function listScheduledQueryRulesByResourceGroup() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "dd4bfc94-a096-412b-9c43-4bd13e35afbc";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "QueryResourceGroupName";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.scheduledQueryRules.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
