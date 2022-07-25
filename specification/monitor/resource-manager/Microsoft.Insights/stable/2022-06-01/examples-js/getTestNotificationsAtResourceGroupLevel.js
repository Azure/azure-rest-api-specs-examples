const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the test notifications by the notification id
 *
 * @summary Get the test notifications by the notification id
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/getTestNotificationsAtResourceGroupLevel.json
 */
async function getNotificationDetailsAtResourceGroupLevel() {
  const subscriptionId = "187f412d-1758-44d9-b052-169e2564721d";
  const resourceGroupName = "Default-TestNotifications";
  const notificationId = "11000222191287";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.actionGroups.getTestNotificationsAtResourceGroupLevel(
    resourceGroupName,
    notificationId
  );
  console.log(result);
}

getNotificationDetailsAtResourceGroupLevel().catch(console.error);
