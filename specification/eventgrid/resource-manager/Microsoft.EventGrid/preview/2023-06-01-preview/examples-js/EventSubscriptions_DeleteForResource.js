const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete an existing event subscription.
 *
 * @summary Delete an existing event subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_DeleteForResource.json
 */
async function eventSubscriptionsDeleteForResource() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope =
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventHub/namespaces/examplenamespace1";
  const eventSubscriptionName = "examplesubscription10";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.eventSubscriptions.beginDeleteAndWait(scope, eventSubscriptionName);
  console.log(result);
}
