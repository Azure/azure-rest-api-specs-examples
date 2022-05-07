Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicebus_6.0.0/sdk/servicebus/arm-servicebus/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new rule and updates an existing rule
 *
 * @summary Creates a new rule and updates an existing rule
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Rules/RuleCreate.json
 */
async function rulesCreateOrUpdate() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroupName";
  const namespaceName = "sdk-Namespace-1319";
  const topicName = "sdk-Topics-2081";
  const subscriptionName = "sdk-Subscriptions-8691";
  const ruleName = "sdk-Rules-6571";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.rules.createOrUpdate(
    resourceGroupName,
    namespaceName,
    topicName,
    subscriptionName,
    ruleName,
    parameters
  );
  console.log(result);
}

rulesCreateOrUpdate().catch(console.error);
```
