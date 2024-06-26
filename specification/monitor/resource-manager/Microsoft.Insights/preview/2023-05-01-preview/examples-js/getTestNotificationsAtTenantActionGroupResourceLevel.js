const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the test notifications by the notification id
 *
 * @summary Get the test notifications by the notification id
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2023-05-01-preview/examples/getTestNotificationsAtTenantActionGroupResourceLevel.json
 */
async function getNotificationDetailsAtTenantActionGroupLevel() {
  const managementGroupId = "11111111-1111-1111-1111-111111111111";
  const tenantActionGroupName = "testTenantActionGroup";
  const xMsClientTenantId = "72f988bf-86f1-41af-91ab-2d7cd011db47";
  const notificationId = "11000222191287";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.getTestNotificationsAtTenantActionGroupResourceLevel(
    managementGroupId,
    tenantActionGroupName,
    xMsClientTenantId,
    notificationId,
  );
  console.log(result);
}
