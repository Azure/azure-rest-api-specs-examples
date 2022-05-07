Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-notificationhubs_2.0.1/sdk/notificationhubs/arm-notificationhubs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NotificationHubsManagementClient } = require("@azure/arm-notificationhubs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available namespaces within the subscription irrespective of the resourceGroups.
 *
 * @summary Lists all the available namespaces within the subscription irrespective of the resourceGroups.
 * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceList.json
 */
async function nameSpaceList() {
  const subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
  const credential = new DefaultAzureCredential();
  const client = new NotificationHubsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.namespaces.listAll()) {
    resArray.push(item);
  }
  console.log(resArray);
}

nameSpaceList().catch(console.error);
```
