const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the full endpoint URL for an event subscription of a partner topic.
 *
 * @summary Get the full endpoint URL for an event subscription of a partner topic.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerTopicEventSubscriptions_GetFullUrl.json
 */
async function partnerTopicEventSubscriptionsGetFullUrl() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const partnerTopicName = "examplePartnerTopic1";
  const eventSubscriptionName = "examplesubscription1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.partnerTopicEventSubscriptions.getFullUrl(
    resourceGroupName,
    partnerTopicName,
    eventSubscriptionName
  );
  console.log(result);
}
