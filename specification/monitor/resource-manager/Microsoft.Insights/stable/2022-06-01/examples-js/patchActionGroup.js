const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing action group's tags. To update other fields use the CreateOrUpdate method.
 *
 * @summary Updates an existing action group's tags. To update other fields use the CreateOrUpdate method.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/patchActionGroup.json
 */
async function patchAnActionGroup() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "187f412d-1758-44d9-b052-169e2564721d";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "Default-NotificationRules";
  const actionGroupName = "SampleActionGroup";
  const actionGroupPatch = {
    enabled: false,
    tags: { key1: "value1", key2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.actionGroups.update(
    resourceGroupName,
    actionGroupName,
    actionGroupPatch
  );
  console.log(result);
}
