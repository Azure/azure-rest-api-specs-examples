const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the test notifications by the notification id
 *
 * @summary Get the test notifications by the notification id
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/getTestNotificationsAtActionGroupResourceLevel.json
 */
async function getNotificationDetailsAtResourceGroupLevel() {
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const resourceGroupName = "TestRgName";
  const actionGroupName = "TestAgName";
  const notificationId = "11000222191287";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.actionGroups.getTestNotificationsAtActionGroupResourceLevel(
    resourceGroupName,
    actionGroupName,
    notificationId
  );
  console.log(result);
}

getNotificationDetailsAtResourceGroupLevel().catch(console.error);
