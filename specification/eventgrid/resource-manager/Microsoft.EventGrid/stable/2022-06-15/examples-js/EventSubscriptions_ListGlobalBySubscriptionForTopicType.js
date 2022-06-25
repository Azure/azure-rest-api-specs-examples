const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all global event subscriptions under an Azure subscription for a topic type.
 *
 * @summary List all global event subscriptions under an Azure subscription for a topic type.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/EventSubscriptions_ListGlobalBySubscriptionForTopicType.json
 */
async function eventSubscriptionsListGlobalBySubscriptionForTopicType() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const topicTypeName = "Microsoft.Resources.Subscriptions";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.eventSubscriptions.listGlobalBySubscriptionForTopicType(
    topicTypeName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

eventSubscriptionsListGlobalBySubscriptionForTopicType().catch(console.error);
