const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously updates an existing event subscription.
 *
 * @summary Asynchronously updates an existing event subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/EventSubscriptions_UpdateForCustomTopic_EventHubDestination.json
 */
async function eventSubscriptionsUpdateForCustomTopicEventHubDestination() {
  const scope =
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2";
  const eventSubscriptionName = "examplesubscription1";
  const eventSubscriptionUpdateParameters = {
    destination: {
      endpointType: "EventHub",
      resourceId:
        "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.EventHub/namespaces/ContosoNamespace/eventhubs/EH1",
    },
    filter: {
      isSubjectCaseSensitive: true,
      subjectBeginsWith: "existingPrefix",
      subjectEndsWith: "newSuffix",
    },
    labels: ["label1", "label2"],
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.beginUpdateAndWait(
    scope,
    eventSubscriptionName,
    eventSubscriptionUpdateParameters
  );
  console.log(result);
}
