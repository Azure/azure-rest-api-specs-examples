const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to asynchronously updates an existing event subscription.
 *
 * @summary asynchronously updates an existing event subscription.
 * x-ms-original-file: 2025-07-15-preview/EventSubscriptions_UpdateForCustomTopic_HybridConnectionDestination.json
 */
async function eventSubscriptionsUpdateForCustomTopicHybridConnectionDestination() {
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.update(
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2",
    "examplesubscription1",
    {
      destination: {
        endpointType: "HybridConnection",
        resourceId:
          "/subscriptions/d33c5f7a-02ea-40f4-bf52-07f17e84d6a8/resourceGroups/TestRG/providers/Microsoft.Relay/namespaces/ContosoNamespace/hybridConnections/HC1",
      },
      filter: {
        isSubjectCaseSensitive: true,
        subjectBeginsWith: "existingPrefix",
        subjectEndsWith: "newSuffix",
      },
      labels: ["label1", "label2"],
    },
  );
  console.log(result);
}
