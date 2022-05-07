Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicebus_6.0.0/sdk/servicebus/arm-servicebus/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Primary and secondary connection strings to the queue.
 *
 * @summary Primary and secondary connection strings to the queue.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/SBQueueAuthorizationRuleListKey.json
 */
async function queueAuthorizationRuleListKey() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-namespace-7982";
  const queueName = "sdk-Queues-2317";
  const authorizationRuleName = "sdk-AuthRules-5800";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.queues.listKeys(
    resourceGroupName,
    namespaceName,
    queueName,
    authorizationRuleName
  );
  console.log(result);
}

queueAuthorizationRuleListKey().catch(console.error);
```
