Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-eventhub_5.0.1/sdk/eventhub/arm-eventhub/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the Event Hubs in a Namespace.
 *
 * @summary Gets all the Event Hubs in a Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubListByNameSpace.json
 */
async function eventHubsListAll() {
  const subscriptionId = "e2f361f0-3b27-4503-a9cc-21cfba380093";
  const resourceGroupName = "Default-NotificationHubs-AustraliaEast";
  const namespaceName = "sdk-Namespace-5357";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.eventHubs.listByNamespace(resourceGroupName, namespaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

eventHubsListAll().catch(console.error);
```
