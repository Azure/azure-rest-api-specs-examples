const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all event subscriptions that have been created for a specific domain topic.
 *
 * @summary List all event subscriptions that have been created for a specific domain topic.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/EventSubscriptions_ListByDomainTopic.json
 */
async function eventSubscriptionsListByDomainTopic() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const domainName = "domain1";
  const topicName = "topic1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.eventSubscriptions.listByDomainTopic(
    resourceGroupName,
    domainName,
    topicName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

eventSubscriptionsListByDomainTopic().catch(console.error);
