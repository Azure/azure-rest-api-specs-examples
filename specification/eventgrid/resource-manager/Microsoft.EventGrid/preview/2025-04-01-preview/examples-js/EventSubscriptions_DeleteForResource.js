const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Delete an existing event subscription.
 *
 * @summary Delete an existing event subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/EventSubscriptions_DeleteForResource.json
 */
async function eventSubscriptionsDeleteForResource() {
  const scope =
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventHub/namespaces/examplenamespace1";
  const eventSubscriptionName = "examplesubscription10";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.beginDeleteAndWait(scope, eventSubscriptionName);
  console.log(result);
}
