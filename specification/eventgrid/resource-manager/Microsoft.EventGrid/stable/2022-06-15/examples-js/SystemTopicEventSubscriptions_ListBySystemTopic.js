const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List event subscriptions that belong to a specific system topic.
 *
 * @summary List event subscriptions that belong to a specific system topic.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/SystemTopicEventSubscriptions_ListBySystemTopic.json
 */
async function systemTopicEventSubscriptionsListBySystemTopic() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const systemTopicName = "exampleSystemTopic1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.systemTopicEventSubscriptions.listBySystemTopic(
    resourceGroupName,
    systemTopicName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

systemTopicEventSubscriptionsListBySystemTopic().catch(console.error);
