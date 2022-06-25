const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get properties of an event subscription.
 *
 * @summary Get properties of an event subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/EventSubscriptions_GetForCustomTopic_EventHubDestination.json
 */
async function eventSubscriptionsGetForCustomTopicEventHubDestination() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope =
    "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2";
  const eventSubscriptionName = "examplesubscription1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.eventSubscriptions.get(scope, eventSubscriptionName);
  console.log(result);
}

eventSubscriptionsGetForCustomTopicEventHubDestination().catch(console.error);
