Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicebus_6.0.0/sdk/servicebus/arm-servicebus/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a subscription description for the specified topic.
 *
 * @summary Returns a subscription description for the specified topic.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Subscriptions/SBSubscriptionGet.json
 */
async function subscriptionGet() {
  const subscriptionId = "Subscriptionid";
  const resourceGroupName = "ResourceGroup";
  const namespaceName = "sdk-Namespace-1349";
  const topicName = "sdk-Topics-8740";
  const subscriptionName = "sdk-Subscriptions-2178";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.subscriptions.get(
    resourceGroupName,
    namespaceName,
    topicName,
    subscriptionName
  );
  console.log(result);
}

subscriptionGet().catch(console.error);
```
