const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get the full endpoint URL for an event subscription.
 *
 * @summary get the full endpoint URL for an event subscription.
 * x-ms-original-file: 2025-07-15-preview/EventSubscriptions_GetFullUrlForSubscription.json
 */
async function eventSubscriptionsGetFullUrlForSubscription() {
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.getFullUrl(
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40",
    "examplesubscription3",
  );
  console.log(result);
}
