const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously creates a new event subscription or updates an existing event subscription.
 *
 * @summary Asynchronously creates a new event subscription or updates an existing event subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/DomainTopicEventSubscriptions_CreateOrUpdate.json
 */
async function domainTopicEventSubscriptionsCreateOrUpdate() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const domainName = "exampleDomain1";
  const topicName = "exampleDomainTopic1";
  const eventSubscriptionName = "exampleEventSubscriptionName1";
  const eventSubscriptionInfo = {
    destination: {
      endpointType: "WebHook",
      endpointUrl: "https://requestb.in/15ksip71",
    },
    filter: {
      isSubjectCaseSensitive: false,
      subjectBeginsWith: "ExamplePrefix",
      subjectEndsWith: "ExampleSuffix",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.domainTopicEventSubscriptions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    domainName,
    topicName,
    eventSubscriptionName,
    eventSubscriptionInfo
  );
  console.log(result);
}

domainTopicEventSubscriptionsCreateOrUpdate().catch(console.error);
