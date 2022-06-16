const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the consumer groups in a Namespace. An empty feed is returned if no consumer group exists in the Namespace.
 *
 * @summary Gets all the consumer groups in a Namespace. An empty feed is returned if no consumer group exists in the Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/ConsumerGroup/EHConsumerGroupListByEventHub.json
 */
async function consumerGroupsListAll() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-2661";
  const eventHubName = "sdk-EventHub-6681";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.consumerGroups.listByEventHub(
    resourceGroupName,
    namespaceName,
    eventHubName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

consumerGroupsListAll().catch(console.error);
