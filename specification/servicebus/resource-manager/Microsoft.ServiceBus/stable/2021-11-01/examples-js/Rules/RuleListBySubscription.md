Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicebus_6.0.0/sdk/servicebus/arm-servicebus/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the rules within given topic-subscription
 *
 * @summary List all the rules within given topic-subscription
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Rules/RuleListBySubscription.json
 */
async function rulesListBySubscriptions() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-1319";
  const topicName = "sdk-Topics-2081";
  const subscriptionName = "sdk-Subscriptions-8691";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.rules.listBySubscriptions(
    resourceGroupName,
    namespaceName,
    topicName,
    subscriptionName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

rulesListBySubscriptions().catch(console.error);
```
