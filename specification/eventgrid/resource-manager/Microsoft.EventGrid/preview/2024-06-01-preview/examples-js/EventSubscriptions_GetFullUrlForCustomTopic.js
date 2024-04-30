const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the full endpoint URL for an event subscription.
 *
 * @summary Get the full endpoint URL for an event subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/EventSubscriptions_GetFullUrlForCustomTopic.json
 */
async function eventSubscriptionsGetFullUrlForCustomTopic() {
  const scope =
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2";
  const eventSubscriptionName = "examplesubscription1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.getFullUrl(scope, eventSubscriptionName);
  console.log(result);
}
