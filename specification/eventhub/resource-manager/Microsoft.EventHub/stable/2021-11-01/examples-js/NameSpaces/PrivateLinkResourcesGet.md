Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-eventhub_5.0.1/sdk/eventhub/arm-eventhub/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets lists of resources that supports Privatelinks.
 *
 * @summary Gets lists of resources that supports Privatelinks.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/PrivateLinkResourcesGet.json
 */
async function nameSpacePrivateLinkResourcesGet() {
  const subscriptionId = "subID";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-2924";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, namespaceName);
  console.log(result);
}

nameSpacePrivateLinkResourcesGet().catch(console.error);
```
