Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-eventhub_5.0.1/sdk/eventhub/arm-eventhub/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Event Hubs consumer group as a nested resource within a Namespace.
 *
 * @summary Creates or updates an Event Hubs consumer group as a nested resource within a Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/ConsumerGroup/EHConsumerGroupCreate.json
 */
async function consumerGroupCreate() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-2661";
  const eventHubName = "sdk-EventHub-6681";
  const consumerGroupName = "sdk-ConsumerGroup-5563";
  const parameters = { userMetadata: "New consumergroup" };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.consumerGroups.createOrUpdate(
    resourceGroupName,
    namespaceName,
    eventHubName,
    consumerGroupName,
    parameters
  );
  console.log(result);
}

consumerGroupCreate().catch(console.error);
```
