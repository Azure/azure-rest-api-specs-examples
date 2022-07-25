const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete an action group.
 *
 * @summary Delete an action group.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/deleteActionGroup.json
 */
async function deleteAnActionGroup() {
  const subscriptionId = "187f412d-1758-44d9-b052-169e2564721d";
  const resourceGroupName = "Default-NotificationRules";
  const actionGroupName = "SampleActionGroup";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.actionGroups.delete(resourceGroupName, actionGroupName);
  console.log(result);
}

deleteAnActionGroup().catch(console.error);
