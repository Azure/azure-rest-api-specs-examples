const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get the full endpoint URL for an event subscription.
 *
 * @summary Get the full endpoint URL for an event subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/EventSubscriptions_GetFullUrlForResourceGroup.json
 */
async function eventSubscriptionsGetFullUrlForResourceGroup() {
  const scope = "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg";
  const eventSubscriptionName = "examplesubscription2";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.getFullUrl(scope, eventSubscriptionName);
  console.log(result);
}
