const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List event subscriptions that belong to a specific partner topic.
 *
 * @summary List event subscriptions that belong to a specific partner topic.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PartnerTopicEventSubscriptions_ListByPartnerTopic.json
 */
async function partnerTopicEventSubscriptionsListByPartnerTopic() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const partnerTopicName = "examplePartnerTopic1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.partnerTopicEventSubscriptions.listByPartnerTopic(
    resourceGroupName,
    partnerTopicName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
