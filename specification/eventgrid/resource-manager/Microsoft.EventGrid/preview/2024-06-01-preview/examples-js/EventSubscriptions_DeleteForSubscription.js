const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete an existing event subscription.
 *
 * @summary Delete an existing event subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/EventSubscriptions_DeleteForSubscription.json
 */
async function eventSubscriptionsDeleteForSubscription() {
  const scope = "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const eventSubscriptionName = "examplesubscription3";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.beginDeleteAndWait(scope, eventSubscriptionName);
  console.log(result);
}
