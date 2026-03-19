const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete an existing event subscription.
 *
 * @summary delete an existing event subscription.
 * x-ms-original-file: 2025-07-15-preview/EventSubscriptions_DeleteForSubscription.json
 */
async function eventSubscriptionsDeleteForSubscription() {
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  await client.eventSubscriptions.delete(
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40",
    "examplesubscription3",
  );
}
