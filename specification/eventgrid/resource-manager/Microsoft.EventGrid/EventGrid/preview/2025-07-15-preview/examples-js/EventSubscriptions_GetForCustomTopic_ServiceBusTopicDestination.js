const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get properties of an event subscription.
 *
 * @summary get properties of an event subscription.
 * x-ms-original-file: 2025-07-15-preview/EventSubscriptions_GetForCustomTopic_ServiceBusTopicDestination.json
 */
async function eventSubscriptionsGetForCustomTopicServiceBusTopicDestination() {
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.get(
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2",
    "examplesubscription1",
  );
  console.log(result);
}
