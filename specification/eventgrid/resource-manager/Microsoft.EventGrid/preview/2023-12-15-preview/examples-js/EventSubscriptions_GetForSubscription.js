const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get properties of an event subscription.
 *
 * @summary Get properties of an event subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/EventSubscriptions_GetForSubscription.json
 */
async function eventSubscriptionsGetForSubscription() {
  const scope = "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const eventSubscriptionName = "examplesubscription3";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.get(scope, eventSubscriptionName);
  console.log(result);
}
