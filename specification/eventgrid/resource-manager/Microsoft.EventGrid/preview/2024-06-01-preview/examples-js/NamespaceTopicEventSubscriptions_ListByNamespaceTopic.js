const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List event subscriptions that belong to a specific namespace topic.
 *
 * @summary List event subscriptions that belong to a specific namespace topic.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/NamespaceTopicEventSubscriptions_ListByNamespaceTopic.json
 */
async function namespaceTopicEventSubscriptionsListByNamespaceTopic() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const namespaceName = "examplenamespace2";
  const topicName = "examplenamespacetopic2";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.namespaceTopicEventSubscriptions.listByNamespaceTopic(
    resourceGroupName,
    namespaceName,
    topicName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
