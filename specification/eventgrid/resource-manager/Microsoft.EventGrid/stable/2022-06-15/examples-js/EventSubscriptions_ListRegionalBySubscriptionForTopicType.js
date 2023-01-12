const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all event subscriptions from the given location under a specific Azure subscription and topic type.
 *
 * @summary List all event subscriptions from the given location under a specific Azure subscription and topic type.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/EventSubscriptions_ListRegionalBySubscriptionForTopicType.json
 */
async function eventSubscriptionsListRegionalBySubscriptionForTopicType() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const location = "westus2";
  const topicTypeName = "Microsoft.EventHub.namespaces";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.eventSubscriptions.listRegionalBySubscriptionForTopicType(
    location,
    topicTypeName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
