const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 *
 * @summary asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 * x-ms-original-file: 2025-07-15-preview/EventSubscriptions_CreateOrUpdateForCustomTopic_ServiceBusTopicDestination.json
 */
async function eventSubscriptionsCreateOrUpdateForCustomTopicServiceBusTopicDestination() {
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.createOrUpdate(
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1",
    "examplesubscription1",
    {
      deadLetterDestination: {
        endpointType: "StorageBlob",
        blobContainerName: "contosocontainer",
        resourceId:
          "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg",
      },
      destination: {
        endpointType: "ServiceBusTopic",
        resourceId:
          "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.ServiceBus/namespaces/ContosoNamespace/topics/SBT",
      },
      filter: {
        isSubjectCaseSensitive: false,
        subjectBeginsWith: "ExamplePrefix",
        subjectEndsWith: "ExampleSuffix",
      },
    },
  );
  console.log(result);
}
