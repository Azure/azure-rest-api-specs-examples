const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all delivery attributes for an event subscription.
 *
 * @summary Get all delivery attributes for an event subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/EventSubscriptions_GetDeliveryAttributes.json
 */
async function eventSubscriptionsGetDeliveryAttributes() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "aaaaaaaaaaaaaaaaaaaaaaaaa";
  const eventSubscriptionName = "aaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.eventSubscriptions.getDeliveryAttributes(
    scope,
    eventSubscriptionName
  );
  console.log(result);
}

eventSubscriptionsGetDeliveryAttributes().catch(console.error);
