Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicebus_6.0.0/sdk/servicebus/arm-servicebus/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a subscription from the specified topic.
 *
 * @summary Deletes a subscription from the specified topic.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Subscriptions/SBSubscriptionDelete.json
 */
async function subscriptionDelete() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "ResourceGroup";
  const namespaceName = "sdk-Namespace-5882";
  const topicName = "sdk-Topics-1804";
  const subscriptionName = "sdk-Subscriptions-3670";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.subscriptions.delete(
    resourceGroupName,
    namespaceName,
    topicName,
    subscriptionName
  );
  console.log(result);
}

subscriptionDelete().catch(console.error);
```
