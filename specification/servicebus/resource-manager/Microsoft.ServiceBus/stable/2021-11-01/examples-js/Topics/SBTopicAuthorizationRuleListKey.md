Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicebus_6.0.0/sdk/servicebus/arm-servicebus/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the primary and secondary connection strings for the topic.
 *
 * @summary Gets the primary and secondary connection strings for the topic.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/SBTopicAuthorizationRuleListKey.json
 */
async function topicAuthorizationRuleListKey() {
  const subscriptionId = "e2f361f0-3b27-4503-a9cc-21cfba380093";
  const resourceGroupName = "Default-ServiceBus-WestUS";
  const namespaceName = "sdk-Namespace8408";
  const topicName = "sdk-Topics2075";
  const authorizationRuleName = "sdk-Authrules5067";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.topics.listKeys(
    resourceGroupName,
    namespaceName,
    topicName,
    authorizationRuleName
  );
  console.log(result);
}

topicAuthorizationRuleListKey().catch(console.error);
```
