const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Send notification to provided channel.
 *
 * @summary Send notification to provided channel.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/NotificationChannels_Notify.json
 */
async function notificationChannelsNotify() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{notificationChannelName}";
  const notifyParameters = {
    eventName: "AutoShutdown",
    jsonPayload:
      '{"eventType":"AutoShutdown","subscriptionId":"{subscriptionId}","resourceGroupName":"resourceGroupName","labName":"{labName}"}',
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.notificationChannels.notify(
    resourceGroupName,
    labName,
    name,
    notifyParameters
  );
  console.log(result);
}

notificationChannelsNotify().catch(console.error);
