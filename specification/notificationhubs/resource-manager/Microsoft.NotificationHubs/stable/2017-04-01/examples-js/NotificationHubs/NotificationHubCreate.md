Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-notificationhubs_2.0.1/sdk/notificationhubs/arm-notificationhubs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NotificationHubsManagementClient } = require("@azure/arm-notificationhubs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates/Update a NotificationHub in a namespace.
 *
 * @summary Creates/Update a NotificationHub in a namespace.
 * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubCreate.json
 */
async function notificationHubCreate() {
  const subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
  const resourceGroupName = "5ktrial";
  const namespaceName = "nh-sdk-ns";
  const notificationHubName = "nh-sdk-hub";
  const parameters = {
    location: "eastus",
  };
  const credential = new DefaultAzureCredential();
  const client = new NotificationHubsManagementClient(credential, subscriptionId);
  const result = await client.notificationHubs.createOrUpdate(
    resourceGroupName,
    namespaceName,
    notificationHubName,
    parameters
  );
  console.log(result);
}

notificationHubCreate().catch(console.error);
```
