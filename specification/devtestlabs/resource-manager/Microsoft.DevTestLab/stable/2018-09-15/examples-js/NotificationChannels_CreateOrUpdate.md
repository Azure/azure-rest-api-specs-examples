Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-devtestlabs_4.0.1/sdk/devtestlabs/arm-devtestlabs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace an existing notification channel.
 *
 * @summary Create or replace an existing notification channel.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/NotificationChannels_CreateOrUpdate.json
 */
async function notificationChannelsCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{notificationChannelName}";
  const notificationChannel = {
    description: "Integration configured for auto-shutdown",
    emailRecipient: "{email}",
    events: [{ eventName: "AutoShutdown" }],
    notificationLocale: "en",
    webHookUrl: "{webhookUrl}",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.notificationChannels.createOrUpdate(
    resourceGroupName,
    labName,
    name,
    notificationChannel
  );
  console.log(result);
}

notificationChannelsCreateOrUpdate().catch(console.error);
```
