const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 *
 * @summary Asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/EventSubscriptions_CreateOrUpdateForCustomTopic_EventHubDestination.json
 */
async function eventSubscriptionsCreateOrUpdateForCustomTopicEventHubDestination() {
  const scope =
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1";
  const eventSubscriptionName = "examplesubscription1";
  const eventSubscriptionInfo = {
    deadLetterDestination: {
      blobContainerName: "contosocontainer",
      endpointType: "StorageBlob",
      resourceId:
        "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg",
    },
    destination: {
      endpointType: "EventHub",
      resourceId:
        "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.EventHub/namespaces/ContosoNamespace/eventhubs/EH1",
    },
    filter: {
      isSubjectCaseSensitive: false,
      subjectBeginsWith: "ExamplePrefix",
      subjectEndsWith: "ExampleSuffix",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.beginCreateOrUpdateAndWait(
    scope,
    eventSubscriptionName,
    eventSubscriptionInfo
  );
  console.log(result);
}
